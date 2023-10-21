package keeper_test

import (
	"testing"

	"fantasfive/x/fantasfive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestAnnounceAndCreateNextMw(t *testing.T) {
	msgServer, _, context := setupMsgServerAnnouceMw(t)
	createResponse, err := msgServer.AnnounceAndCreateNextMw(context, &types.MsgAnnounceAndCreateNextMw{
		Creator:    alice,
		MwId:       "1",
		PlayerPerf: "Ras-0-10-Sal-1-9",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgAnnounceAndCreateNextMwResponse{
		NextMwId: "2", // TODO: update with a proper value when updated
	}, *createResponse)
}

func TestAnnounceAndCreateNextMwEmitted(t *testing.T) {
	msgSrvr, _, context := setupMsgServerAnnouceMw(t)

	//create a team
	_, _ = msgSrvr.CreateTeam(context, &types.MsgCreateTeam{
		Creator: alice,
		Player0: "Beckham",
		Player1: "Ronaldo",
		Player2: "Messi",
		Player3: "Maradona",
		Player4: "Pele",
	})

	_, _ = msgSrvr.AnnounceAndCreateNextMw(context, &types.MsgAnnounceAndCreateNextMw{
		Creator:    alice,
		MwId:       "1",
		PlayerPerf: "Beckham-0-10-Ronaldo-1-9",
	})
	ctx := sdk.UnwrapSDKContext(context)
	require.NotNil(t, ctx)
	events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
	require.Len(t, events, 3)
	eventAnnouced := events[0]
	require.EqualValues(t, sdk.StringEvent{
		Type: "match-week-announced",
		Attributes: []sdk.Attribute{
			{Key: ("announcer"), Value: (alice)},
			{Key: ("mw-id"), Value: ("1")},
			{Key: ("team-count"), Value: ("1")},
			{Key: ("winner-team-id"), Value: ("0")},
			{Key: ("players-performance"), Value: ("Beckham-0-10-Ronaldo-1-9")},
		},
	}, eventAnnouced)
	eventCreated := events[1]
	require.EqualValues(t, sdk.StringEvent{
		Type: "match-week-created",
		Attributes: []sdk.Attribute{
			{Key: ("creator"), Value: (alice)},
			{Key: ("mw-id"), Value: ("2")},
		},
	}, eventCreated)
}
