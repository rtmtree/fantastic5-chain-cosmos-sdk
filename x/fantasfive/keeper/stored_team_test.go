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

func createNStoredTeam(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredTeam {
	items := make([]types.StoredTeam, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredTeam(ctx, items[i])
	}
	return items
}

func TestStoredTeamGet(t *testing.T) {
	keeper, ctx := keepertest.FantasfiveKeeper(t)
	items := createNStoredTeam(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredTeam(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredTeamRemove(t *testing.T) {
	keeper, ctx := keepertest.FantasfiveKeeper(t)
	items := createNStoredTeam(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredTeam(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredTeam(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredTeamGetAll(t *testing.T) {
	keeper, ctx := keepertest.FantasfiveKeeper(t)
	items := createNStoredTeam(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredTeam(ctx)),
	)
}
