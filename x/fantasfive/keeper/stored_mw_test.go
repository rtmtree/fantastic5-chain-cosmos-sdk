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

func createNStoredMW(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredMW {
	items := make([]types.StoredMW, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredMW(ctx, items[i])
	}
	return items
}

func TestStoredMWGet(t *testing.T) {
	keeper, ctx := keepertest.FantasFiveKeeper(t)
	items := createNStoredMW(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredMW(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredMWRemove(t *testing.T) {
	keeper, ctx := keepertest.FantasFiveKeeper(t)
	items := createNStoredMW(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredMW(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredMW(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredMWGetAll(t *testing.T) {
	keeper, ctx := keepertest.FantasFiveKeeper(t)
	items := createNStoredMW(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredMW(ctx)),
	)
}
