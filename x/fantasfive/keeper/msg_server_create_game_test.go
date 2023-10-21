package keeper_test

import (
	"testing"

	"fantasfive/x/fantasfive/types"

	"github.com/stretchr/testify/require"
)

func TestCreateTeam(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateTeam(context, &types.MsgCreateTeam{
		Creator: alice,
		Player0: "Beckham",
		Player1: "Ronaldo",
		Player2: "Messi",
		Player3: "Maradona",
		Player4: "Pele",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateTeamResponse{
		TeamId: "", // TODO: update with a proper value when updated
	}, *createResponse)
}
