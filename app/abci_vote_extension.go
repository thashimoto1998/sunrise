package app

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	blocksdkabci "github.com/skip-mev/block-sdk/v2/abci"
	"github.com/spf13/cast"
	"github.com/sunriselayer/sunrise/x/da/keeper"
	damodulekeeper "github.com/sunriselayer/sunrise/x/da/keeper"
	"github.com/sunriselayer/sunrise/x/da/types"
	"github.com/sunriselayer/sunrise/x/da/zkp"
)

const flagDAShardHashesAPI = "da.shard_hashes_api"

type DAConfig struct {
	ShardHashesAPI string `mapstructure:"shard_hashes_api"`
}

type DAShardHashesResponse struct {
	ShardHashes []string `json:"shard_hashes"`
}

// ReadOracleConfig reads the wasm specifig configuration
func ReadDAConfig(opts servertypes.AppOptions) (DAConfig, error) {
	cfg := DAConfig{}
	var err error
	if v := opts.Get(flagDAShardHashesAPI); v != nil {
		if cfg.ShardHashesAPI, err = cast.ToStringE(v); err != nil {
			return cfg, err
		}
	}

	return cfg, nil
}

func GetDataShardHashes(daConfig DAConfig, metadataUri string, n, threshold int64, valAddr sdk.ValAddress) ([]int64, [][]byte, error) {
	indices := types.ShardIndicesForValidator(valAddr, n, threshold)
	fmt.Println("GetDataShardHashes-1", n, threshold, len(indices), time.Now())

	url := daConfig.ShardHashesAPI + "?metadata_uri=" + metadataUri + "&indices=" + strings.Trim(strings.Replace(fmt.Sprint(indices), " ", ",", -1), "[]")
	fmt.Println("GetDataShardHashes-1-1", url)
	res, err := http.Get(url)
	fmt.Println("GetDataShardHashes-1-2", err)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	fmt.Println("GetDataShardHashes-2", time.Now())

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("GetDataShardHashes-3", time.Now())
	daShardHashes := DAShardHashesResponse{}
	err = json.Unmarshal(resBody, &daShardHashes)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("GetDataShardHashes-4", time.Now())
	shares := [][]byte{}
	for _, shareEncoded := range daShardHashes.ShardHashes {
		share, err := base64.StdEncoding.DecodeString(shareEncoded)
		if err != nil {
			continue
		}
		shares = append(shares, share)
	}
	fmt.Println("GetDataShardHashes-5", time.Now())
	return indices, shares, nil
}

type VoteExtHandler struct {
	Keeper        keeper.Keeper
	stakingKeeper *stakingkeeper.Keeper
}

func NewVoteExtHandler(
	keeper keeper.Keeper,
	stakingKeeper *stakingkeeper.Keeper,
) *VoteExtHandler {
	return &VoteExtHandler{
		Keeper:        keeper,
		stakingKeeper: stakingKeeper,
	}
}

func (h *VoteExtHandler) ValidatorsInfo(ctx sdk.Context, consAddr []byte) (int64, sdk.ValAddress, error) {
	iterator, err := h.stakingKeeper.ValidatorsPowerStoreIterator(ctx)
	if err != nil {
		return 0, nil, err
	}

	defer iterator.Close()

	foundValAddr := sdk.ValAddress{}
	numValidators := int64(0)
	for ; iterator.Valid(); iterator.Next() {
		valAddr := sdk.ValAddress(iterator.Value())
		validator, err := h.stakingKeeper.Validator(ctx, iterator.Value())
		if err != nil {
			return 0, nil, err
		}

		if !validator.IsBonded() {
			continue
		}

		valConsAddr, err := validator.GetConsAddr()
		if err != nil {
			return 0, nil, err
		}

		if bytes.Equal(consAddr, valConsAddr) {
			foundValAddr = valAddr
		}

		numValidators++
	}
	return numValidators, foundValAddr, nil
}

func (h *VoteExtHandler) ExtendVoteHandler(daConfig DAConfig, dec sdk.TxDecoder, handler sdk.AnteHandler, daKeeper damodulekeeper.Keeper) sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
		voteExt := types.VoteExtension{
			Data: []*types.PublishedData{},
			// Shares: []*types.DataShares{},
		}

		txs := req.Txs
		for _, tx := range txs {
			sdkTx, err := dec(tx)
			if err != nil {
				continue
			}

			ctx, err = handler(ctx, sdkTx, false)
			if err != nil {
				continue
			}

			params := daKeeper.GetParams(ctx)
			numValidators, valAddr, err := h.ValidatorsInfo(ctx, req.ProposerAddress)
			if err != nil {
				continue
			}

			msgs := sdkTx.GetMsgs()
			for _, msg := range msgs {
				switch msg := msg.(type) {
				case *types.MsgPublishData:
					fmt.Println("ExtendVoteHandler-1", time.Now())
					threshold := params.ReplicationFactor.QuoInt64(numValidators).MulInt64(int64(len(msg.ShardDoubleHashes))).RoundInt64()
					if threshold > int64(len(msg.ShardDoubleHashes)) {
						threshold = int64(len(msg.ShardDoubleHashes))
					}

					indices, shares, err := GetDataShardHashes(daConfig, msg.MetadataUri, int64(len(msg.ShardDoubleHashes)), threshold, valAddr)
					if err != nil {
						continue
					}

					fmt.Println("ExtendVoteHandler-2", time.Now())
					// filter zkp verified data
					err = zkp.VerifyData(indices, shares, msg.ShardDoubleHashes, int(threshold))
					if err != nil {
						continue
					}

					fmt.Println("ExtendVoteHandler-3", time.Now())
					voteExt.Data = append(voteExt.Data, &types.PublishedData{
						MetadataUri:       msg.MetadataUri,
						ParityShardCount:  msg.ParityShardCount,
						ShardDoubleHashes: msg.ShardDoubleHashes,
					})
					fmt.Println("ExtendVoteHandler-4")
					// voteExt.Shares = append(voteExt.Shares, &types.DataShares{
					// 	Indices: indices,
					// 	Shares:  shares,
					// })
				}
			}
		}

		bz, err := voteExt.Marshal()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal vote extension: %w", err)
		}

		return &abci.ResponseExtendVote{VoteExtension: bz}, nil
	}
}

func (h *VoteExtHandler) VerifyVoteExtensionHandler(daConfig DAConfig, daKeeper damodulekeeper.Keeper) sdk.VerifyVoteExtensionHandler {
	return func(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
		var voteExt types.VoteExtension

		if len(req.VoteExtension) == 0 {
			return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
		}

		err := voteExt.Unmarshal(req.VoteExtension)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal vote extension: %w", err)
		}

		// check vote extension data with zkp
		// params := daKeeper.GetParams(ctx)
		// numValidators, _, err := h.ValidatorsInfo(ctx, req.ValidatorAddress)
		// if err != nil {
		// 	return nil, err
		// }

		// for _, data := range voteExt.Data {
		// 	threshold := params.ReplicationFactor.QuoInt64(numValidators).MulInt64(int64(len(data.ShardDoubleHashes))).RoundInt64()
		// 	if threshold > int64(len(data.ShardDoubleHashes)) {
		// 		threshold = int64(len(data.ShardDoubleHashes))
		// 	}
		// 	err = zkp.VerifyData(voteExt.Shares[i].Indices, voteExt.Shares[i].Shares, data.ShardDoubleHashes, int(threshold))
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// }

		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
	}
}

type VoteExtensionTx struct {
	VotedData          map[string]*types.PublishedData
	FaultValidators    map[string]sdk.ValAddress
	ExtendedCommitInfo abci.ExtendedCommitInfo
}

type ProposalHandler struct {
	logger                 log.Logger
	keeper                 keeper.Keeper
	stakingKeeper          *stakingkeeper.Keeper
	DefaultProposalHandler *blocksdkabci.ProposalHandler
	ModuleManager          *module.Manager
}

func NewProposalHandler(logger log.Logger, keeper keeper.Keeper, stakingKeeper *stakingkeeper.Keeper, ModuleManager *module.Manager, proposalHandler *blocksdkabci.ProposalHandler) *ProposalHandler {
	return &ProposalHandler{
		logger:                 logger,
		keeper:                 keeper,
		stakingKeeper:          stakingKeeper,
		ModuleManager:          ModuleManager,
		DefaultProposalHandler: proposalHandler,
	}
}

func (h *ProposalHandler) PrepareProposal() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		defaultHandler := h.DefaultProposalHandler.PrepareProposalHandler()
		defaultResponse, err := defaultHandler(ctx, req)
		if err != nil {
			return nil, err
		}

		proposalTxs := defaultResponse.Txs

		fmt.Println("PrepareProposal-1")
		cp := ctx.ConsensusParams()
		voteExtEnabled := cp.Abci != nil && req.Height > cp.Abci.VoteExtensionsEnableHeight && cp.Abci.VoteExtensionsEnableHeight != 0
		if !voteExtEnabled {
			return &abci.ResponsePrepareProposal{Txs: proposalTxs}, nil
		}

		fmt.Println("PrepareProposal-2")
		err = baseapp.ValidateVoteExtensions(ctx, h.stakingKeeper, req.Height, ctx.ChainID(), req.LocalLastCommit)
		if err != nil {
			return nil, err
		}

		fmt.Println("PrepareProposal-3")
		votedData, faultValidators, err := h.GetVotedDataAndFaultValidators(ctx, req.LocalLastCommit)
		votedDataBz, _ := json.Marshal(votedData)
		faultValidatorsBz, _ := json.Marshal(faultValidators)
		fmt.Println("PrepareProposal-4", string(votedDataBz), string(faultValidatorsBz), err)
		if err != nil {
			return nil, errors.New("failed to get voted data and fault validators")
		}

		voteExtTx := VoteExtensionTx{
			VotedData:          votedData,
			FaultValidators:    faultValidators,
			ExtendedCommitInfo: req.LocalLastCommit,
		}

		fmt.Println("PrepareProposal-5")
		bz, err := json.Marshal(voteExtTx)
		if err != nil {
			h.logger.Error("failed to marshal vote extension tx", "err", err)
			return nil, errors.New("failed to marshal vote extension tx")
		}

		fmt.Println("PrepareProposal-6")
		proposalTxs = append(proposalTxs, bz)
		return &abci.ResponsePrepareProposal{Txs: proposalTxs}, nil
	}
}

func (h *ProposalHandler) ProcessProposal() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		txs := [][]byte{}
		var voteExtTx VoteExtensionTx
		fmt.Println("ProcessProposal-1")
		for _, tx := range req.Txs {
			if err := json.Unmarshal(tx, &voteExtTx); err == nil {
				fmt.Println("ProcessProposal-2")
				err := baseapp.ValidateVoteExtensions(ctx, h.stakingKeeper, req.Height, ctx.ChainID(), voteExtTx.ExtendedCommitInfo)
				if err != nil {
					return nil, err
				}

				fmt.Println("ProcessProposal-3")
				votedData, faultValidators, err := h.GetVotedDataAndFaultValidators(ctx, voteExtTx.ExtendedCommitInfo)
				if err != nil {
					return nil, errors.New("failed to get voted data and fault validators")
				}

				fmt.Println("ProcessProposal-4")
				if err := ConfirmVotedData(voteExtTx.VotedData, votedData); err != nil {
					return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
				}

				fmt.Println("ProcessProposal-5")
				if err := ConfirmFaultValidators(voteExtTx.FaultValidators, faultValidators); err != nil {
					return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, nil
				}
				fmt.Println("ProcessProposal-6")
			} else {
				txs = append(txs, tx)
			}
		}

		fmt.Println("ProcessProposal-7")
		defaultReq := *req
		defaultReq.Txs = txs
		defaultHandler := h.DefaultProposalHandler.ProcessProposalHandler()
		fmt.Println("ProcessProposal-8")
		return defaultHandler(ctx, &defaultReq)
	}
}

func (h *ProposalHandler) PreBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	fmt.Println("PreBlocker-1")
	ctx = ctx.WithEventManager(sdk.NewEventManager())
	paramsChanged := false
	for _, moduleName := range h.ModuleManager.OrderPreBlockers {
		if module, ok := h.ModuleManager.Modules[moduleName].(appmodule.HasPreBlocker); ok {
			rsp, err := module.PreBlock(ctx)
			if err != nil {
				return nil, err
			}
			if rsp.IsConsensusParamsChanged() {
				paramsChanged = true
			}
		}
	}
	fmt.Println("PreBlocker-2")

	for _, txBytes := range req.Txs {
		fmt.Println("PreBlocker-3")
		var voteExtTx VoteExtensionTx
		if err := json.Unmarshal(txBytes, &voteExtTx); err != nil {
			continue
		}

		fmt.Println("PreBlocker-4")
		for _, data := range voteExtTx.VotedData {
			published := h.keeper.GetPublishedData(ctx, data.MetadataUri)
			published.MetadataUri = data.MetadataUri
			published.ParityShardCount = data.ParityShardCount
			published.ShardDoubleHashes = data.ShardDoubleHashes
			published.Status = "vote_extension"
			if err := h.keeper.SetPublishedData(ctx, published); err != nil {
				panic(err)
			}
		}

		fmt.Println("PreBlocker-5")
		for _, valAddr := range voteExtTx.FaultValidators {
			h.keeper.SetFaultCounter(ctx, valAddr, h.keeper.GetFaultCounter(ctx, valAddr)+1)
		}

		fmt.Println("PreBlocker-6")
		params := h.keeper.GetParams(ctx)
		if ctx.BlockHeight()%int64(params.SlashEpoch) == 0 {
			h.keeper.HandleSlashEpoch(ctx)
		}
		fmt.Println("PreBlocker-7")
	}

	fmt.Println("PreBlocker-8")
	return &sdk.ResponsePreBlock{ConsensusParamsChanged: paramsChanged}, nil
}

type DataVote struct {
	Data  types.PublishedData
	Voter sdk.ValAddress
	Power int64
}

type DataVotes []DataVote

func (dv DataVotes) Len() int {
	return len(dv)
}

func (dv DataVotes) Less(i, j int) bool {
	return strings.Compare(dv[i].Data.String(), dv[j].Data.String()) < 0
}

func (dv DataVotes) Swap(i, j int) {
	dv[i], dv[j] = dv[j], dv[i]
}

type ValidatorPower struct {
	Power   int64
	ValAddr sdk.ValAddress
}

func (h *ProposalHandler) GetDataVotesMapByHash(
	commitInfo abci.ExtendedCommitInfo,
	valPowerMap map[string]ValidatorPower,
	valConsAddrMap map[string]sdk.ValAddress,
) (dataVotesMap map[string]DataVotes) {

	dataVotesMap = map[string]DataVotes{}
	for _, v := range commitInfo.Votes {
		valAddr := valConsAddrMap[sdk.ConsAddress(v.Validator.Address).String()]
		valPower, ok := valPowerMap[valAddr.String()]
		if ok {
			var voteExt types.VoteExtension
			if err := voteExt.Unmarshal(v.VoteExtension); err != nil {
				h.logger.Error("failed to decode vote extension", "validator", valAddr.String(), "error", err)
				return
			}

			for _, data := range voteExt.Data {
				dataVote := DataVote{
					Data:  *data,
					Voter: valAddr,
					Power: valPower.Power,
				}
				dataVotesMap[data.MetadataUri] = append(dataVotesMap[data.MetadataUri], dataVote)
			}
		}
	}

	// sort data votes by full data string representation
	for dataHash, dataVotes := range dataVotesMap {
		sort.Sort(dataVotes)
		dataVotesMap[dataHash] = dataVotes
	}

	return
}

func GetAboveThresholdVotedData(
	dataVotes DataVotes, thresholdPower int64,
	valPowerMap map[string]ValidatorPower, faultValidators map[string]sdk.ValAddress,
) (types.PublishedData, bool) {
	lastPublishedData := types.PublishedData{}
	lastPower := int64(0)

	for _, vote := range dataVotes {
		if lastPublishedData.String() == vote.Data.String() {
			lastPower += vote.Power
			if lastPower >= thresholdPower {
				break
			}
			continue
		}
		lastPublishedData = vote.Data
		lastPower = vote.Power
	}

	if lastPower < thresholdPower {
		return lastPublishedData, false
	}

	for _, vote := range dataVotes {
		key := vote.Voter.String()
		valPower := valPowerMap[key]
		if lastPublishedData.String() == vote.Data.String() {
			valPower := valPowerMap[key]
			valPowerMap[key] = valPower
		} else {
			faultValidators[valPower.ValAddr.String()] = valPower.ValAddr
		}
	}

	return lastPublishedData, true
}

func ApplyNotVotedValidators(
	votedData map[string]*types.PublishedData,
	dataVotesMap map[string]DataVotes,
	valPowerMap map[string]ValidatorPower,
	faultValidators map[string]sdk.ValAddress,
) {
	dataHashVoterMap := map[string]map[string]bool{}
	for dataHash := range votedData {
		dataHashVoterMap[dataHash] = map[string]bool{}
	}

	for dataHash, dataVotes := range dataVotesMap {
		for _, dataVote := range dataVotes {
			_, ok := dataHashVoterMap[dataHash]
			if !ok {
				continue
			}

			dataHashVoterMap[dataHash][dataVote.Voter.String()] = true
		}
	}

	for _, valPower := range valPowerMap {
		for dataHash := range votedData {
			_, ok := dataHashVoterMap[dataHash][valPower.ValAddr.String()]
			if !ok {
				faultValidators[valPower.ValAddr.String()] = valPower.ValAddr
				break
			}
		}
	}
}

func (h *ProposalHandler) GetVotedDataAndFaultValidators(ctx sdk.Context, commitInfo abci.ExtendedCommitInfo) (map[string]*types.PublishedData, map[string]sdk.ValAddress, error) {
	votedData := make(map[string]*types.PublishedData)
	valPowerMap := make(map[string]ValidatorPower)
	valConsAddrMap := make(map[string]sdk.ValAddress)

	powerReduction := h.stakingKeeper.PowerReduction(ctx)
	iterator, err := h.stakingKeeper.ValidatorsPowerStoreIterator(ctx)
	if err != nil {
		return nil, nil, err
	}

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		validator, err := h.stakingKeeper.Validator(ctx, iterator.Value())
		if err != nil {
			return nil, nil, err
		}

		if validator.IsBonded() {
			valAddrStr := validator.GetOperator()
			valAddr, err := sdk.ValAddressFromBech32(valAddrStr)
			if err != nil {
				return nil, nil, err
			}

			valPowerMap[valAddr.String()] = ValidatorPower{
				Power:   validator.GetConsensusPower(powerReduction),
				ValAddr: valAddr,
			}

			consAddr, err := validator.GetConsAddr()
			if err != nil {
				return nil, nil, err
			}
			valConsAddrMap[sdk.ConsAddress(consAddr).String()] = valAddr
		}
	}

	params := h.keeper.GetParams(ctx)
	faultValidators := map[string]sdk.ValAddress{}
	dataVotesMap := h.GetDataVotesMapByHash(commitInfo, valPowerMap, valConsAddrMap)
	for dataHash, dataVote := range dataVotesMap {
		bondedTokens, err := h.stakingKeeper.TotalBondedTokens(ctx)
		if err != nil {
			return nil, nil, err
		}

		totalBondedPower := sdk.TokensToConsensusPower(bondedTokens, h.stakingKeeper.PowerReduction(ctx))
		thresholdPower := params.VoteThreshold.MulInt64(totalBondedPower).RoundInt().Int64()

		publishedData, aboveThreshold := GetAboveThresholdVotedData(dataVote, thresholdPower, valPowerMap, faultValidators)
		if !aboveThreshold {
			return nil, nil, err
		}

		votedData[dataHash] = &publishedData
	}

	ApplyNotVotedValidators(votedData, dataVotesMap, valPowerMap, faultValidators)
	return votedData, faultValidators, nil
}

func ConfirmVotedData(actual, expected map[string]*types.PublishedData) error {
	for dataHash, data := range actual {
		if data.String() != expected[dataHash].String() {
			return fmt.Errorf("invalid voted data: %x", []byte(dataHash))
		}
	}

	for dataHash, data := range expected {
		if data.String() != actual[dataHash].String() {
			return fmt.Errorf("invalid voted data: %x", []byte(dataHash))
		}
	}

	return nil
}

func ConfirmFaultValidators(actual, expected map[string]sdk.ValAddress) error {
	if !reflect.DeepEqual(actual, expected) {
		return errors.New("invalid fault validators")
	}
	return nil
}
