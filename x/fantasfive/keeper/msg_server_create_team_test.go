package keeper_test

import (
	"testing"

	"fantasfive/x/fantasfive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestCreateTeam(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateTeam(t)
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
		TeamId: "0",
	}, *createResponse)
}

func TestCreate1TeamEmitted(t *testing.T) {
	msgSrvr, _, context := setupMsgServerCreateTeam(t)
	_, _ = msgSrvr.CreateTeam(context, &types.MsgCreateTeam{
		Creator: alice,
		Player0: "Beckham",
		Player1: "Ronaldo",
		Player2: "Messi",
		Player3: "Maradona",
		Player4: "Pele",
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 1)
	event := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "new-team-created",
		Attributes: []sdk.Attribute{
			{Key: ("creator"), Value: (alice)},
			{Key: ("team-id"), Value: ("0")},
			{Key: ("mw-id"), Value: ("1")},
			{Key: ("player-0"), Value: ("Beckham")},
			{Key: ("player-1"), Value: ("Ronaldo")},
			{Key: ("player-2"), Value: ("Messi")},
			{Key: ("player-3"), Value: ("Maradona")},
			{Key: ("player-4"), Value: ("Pele")},
			{Key: ("captain-index"), Value: ("0")},
		},
	}, event)
}
