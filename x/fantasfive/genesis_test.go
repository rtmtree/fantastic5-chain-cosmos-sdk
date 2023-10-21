package fantasfive_test

import (
	"testing"

	keepertest "fantasfive/testutil/keeper"
	"fantasfive/testutil/nullify"
	"fantasfive/x/fantasfive"
	"fantasfive/x/fantasfive/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfoList: []types.SystemInfo{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		StoredMWList: []types.StoredMW{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		StoredTeamList: []types.StoredTeam{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FantasFiveKeeper(t)
	fantasfive.InitGenesis(ctx, *k, genesisState)
	got := fantasfive.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.SystemInfoList, got.SystemInfoList)
	require.ElementsMatch(t, genesisState.StoredMWList, got.StoredMWList)
	require.ElementsMatch(t, genesisState.StoredTeamList, got.StoredTeamList)
	// this line is used by starport scaffolding # genesis/test/assert
}
