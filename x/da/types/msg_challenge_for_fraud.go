package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgChallengeForFraud{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgChallengeForFraud) ValidateBasic() error {
	return nil
}
