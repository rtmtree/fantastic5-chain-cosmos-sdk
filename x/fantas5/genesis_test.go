package fantas5_test

import (
	"testing"

	keepertest "fantas5/testutil/keeper"
	"fantas5/testutil/nullify"
	"fantas5/x/fantas5"
	"fantas5/x/fantas5/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Fantas5Keeper(t)
	fantas5.InitGenesis(ctx, *k, genesisState)
	got := fantas5.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
