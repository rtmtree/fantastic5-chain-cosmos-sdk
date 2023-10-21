package keeper_test

import (
	"strconv"
	"testing"

	keepertest "fantasfive/testutil/keeper"
	"fantasfive/testutil/nullify"
	"fantasfive/x/fantasfive/keeper"
	"fantasfive/x/fantasfive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNSystemInfo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SystemInfo {
	items := make([]types.SystemInfo, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetSystemInfo(ctx, items[i])
	}
	return items
}

func TestSystemInfoGet(t *testing.T) {
	keeper, ctx := keepertest.FantasFiveKeeper(t)
	items := createNSystemInfo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetSystemInfo(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestSystemInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.FantasFiveKeeper(t)
	items := createNSystemInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSystemInfo(ctx,
			item.Index,
		)
		_, found := keeper.GetSystemInfo(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestSystemInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.FantasFiveKeeper(t)
	items := createNSystemInfo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSystemInfo(ctx)),
	)
}
