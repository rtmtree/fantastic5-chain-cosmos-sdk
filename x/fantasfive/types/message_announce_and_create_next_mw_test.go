package types

import (
	"testing"

	"fantasfive/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAnnounceAndCreateNextMw_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAnnounceAndCreateNextMw
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgAnnounceAndCreateNextMw{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgAnnounceAndCreateNextMw{
				Creator: sample.AccAddress(),
			},
		}, {
			name: "invalid player performance",
			msg: MsgAnnounceAndCreateNextMw{
				Creator:    sample.AccAddress(),
				PlayerPerf: "Ronaldo 0 -9 Messi 1 8",
			},
		}, {
			name: "valid player performance",
			msg: MsgAnnounceAndCreateNextMw{
				Creator:    sample.AccAddress(),
				PlayerPerf: "Ronaldo 0 9 Messi 1 8",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
