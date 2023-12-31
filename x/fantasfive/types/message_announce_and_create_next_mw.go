package types

import (
	"fantasfive/x/fantasfive/rules"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAnnounceAndCreateNextMw = "announce_and_create_next_mw"

var _ sdk.Msg = &MsgAnnounceAndCreateNextMw{}

func NewMsgAnnounceAndCreateNextMw(creator string, mwId string, playerPerf string) *MsgAnnounceAndCreateNextMw {
	return &MsgAnnounceAndCreateNextMw{
		Creator:    creator,
		MwId:       mwId,
		PlayerPerf: playerPerf,
	}
}

func (msg *MsgAnnounceAndCreateNextMw) Route() string {
	return RouterKey
}

func (msg *MsgAnnounceAndCreateNextMw) Type() string {
	return TypeMsgAnnounceAndCreateNextMw
}

func (msg *MsgAnnounceAndCreateNextMw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAnnounceAndCreateNextMw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAnnounceAndCreateNextMw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// TODO: validate Creator is Authorized b/c announcing is important

	// Check if PlayerPerf is valid
	playerPerf, err := rules.ParsePlayersPerformance(msg.PlayerPerf)
	if err != nil {
		return err
	}

	// Check if MwId is valid
	err = playerPerf.ValidPlayersPerformance()
	if err != nil {
		return err
	}

	return nil
}
