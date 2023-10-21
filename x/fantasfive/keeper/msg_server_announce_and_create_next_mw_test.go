package keeper_test

import (
	"testing"

	"fantasfive/x/fantasfive/types"

	"github.com/stretchr/testify/require"
)

func TestAnnounceAndCreateNextMw(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.AnnounceAndCreateNextMw(context, &types.MsgAnnounceAndCreateNextMw{
		Creator:    alice,
		MwId:       "0",
		PlayerPerf: "Ras 0 10 Sal 1 9",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgAnnounceAndCreateNextMwResponse{
		NextMwId: "1", // TODO: update with a proper value when updated
	}, *createResponse)
}
