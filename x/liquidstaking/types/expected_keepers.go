package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/core/address"
	"cosmossdk.io/math"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here

	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error

	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	SendCoinsFromModuleToModule(ctx sdk.Context, senderPool, recipientPool string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error

	IterateTotalSupply(ctx sdk.Context, cb func(sdk.Coin) bool)
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

// StakingKeeper expected staking keeper
type StakingKeeper interface {
	ValidatorAddressCodec() address.Codec
	ConsensusAddressCodec() address.Codec
	// iterate through validators by operator address, execute func for each validator
	IterateValidators(context.Context,
		func(index int64, validator stakingtypes.ValidatorI) (stop bool)) error

	Validator(context.Context, sdk.ValAddress) (stakingtypes.ValidatorI, error)            // get a particular validator by operator address
	ValidatorByConsAddr(context.Context, sdk.ConsAddress) (stakingtypes.ValidatorI, error) // get a particular validator by consensus address

	// slash the validator and delegators of the validator, specifying offense height, offense power, and slash fraction
	Slash(context.Context, sdk.ConsAddress, int64, int64, math.LegacyDec) (math.Int, error)
	SlashWithInfractionReason(context.Context, sdk.ConsAddress, int64, int64, math.LegacyDec, stakingtypes.Infraction) (math.Int, error)
	Jail(context.Context, sdk.ConsAddress) error   // jail a validator
	Unjail(context.Context, sdk.ConsAddress) error // unjail a validator

	// Delegation allows for getting a particular delegation for a given validator
	// and delegator outside the scope of the staking module.
	Delegation(context.Context, sdk.AccAddress, sdk.ValAddress) (stakingtypes.DelegationI, error)
	GetAllValidators(ctx context.Context) ([]stakingtypes.Validator, error)

	// MaxValidators returns the maximum amount of bonded validators
	MaxValidators(context.Context) (uint32, error)

	// IsValidatorJailed returns if the validator is jailed.
	IsValidatorJailed(ctx context.Context, addr sdk.ConsAddress) (bool, error)
}

type DistributionKeeper interface {
	GetDelegatorWithdrawAddr(ctx sdk.Context, delAddr sdk.AccAddress) sdk.AccAddress
	WithdrawDelegationRewards(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error)
}
