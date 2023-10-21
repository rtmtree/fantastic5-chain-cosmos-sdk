package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "fantasfive/testutil/keeper"
	"fantasfive/testutil/nullify"
	"fantasfive/x/fantasfive/keeper"
	"fantasfive/x/fantasfive/types"
)

func createTestTeamInfo(keeper *keeper.Keeper, ctx sdk.Context) types.TeamInfo {
	item := types.TeamInfo{}
	keeper.SetTeamInfo(ctx, item)
	return item
}

func TestTeamInfoGet(t *testing.T) {
	keeper, ctx := keepertest.FantasfiveKeeper(t)
	item := createTestTeamInfo(keeper, ctx)
	rst, found := keeper.GetTeamInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTeamInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.FantasfiveKeeper(t)
	createTestTeamInfo(keeper, ctx)
	keeper.RemoveTeamInfo(ctx)
	_, found := keeper.GetTeamInfo(ctx)
	require.False(t, found)
}
