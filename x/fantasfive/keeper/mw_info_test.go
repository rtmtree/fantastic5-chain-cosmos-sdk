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

func createTestMwInfo(keeper *keeper.Keeper, ctx sdk.Context) types.MwInfo {
	item := types.MwInfo{}
	keeper.SetMwInfo(ctx, item)
	return item
}

func TestMwInfoGet(t *testing.T) {
	keeper, ctx := keepertest.FantasfiveKeeper(t)
	item := createTestMwInfo(keeper, ctx)
	rst, found := keeper.GetMwInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestMwInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.FantasfiveKeeper(t)
	createTestMwInfo(keeper, ctx)
	keeper.RemoveMwInfo(ctx)
	_, found := keeper.GetMwInfo(ctx)
	require.False(t, found)
}
