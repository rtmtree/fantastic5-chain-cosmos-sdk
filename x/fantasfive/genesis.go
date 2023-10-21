package fantasfive

import (
	"fantasfive/x/fantasfive/keeper"
	"fantasfive/x/fantasfive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the systemInfo
	for _, elem := range genState.SystemInfoList {
		k.SetSystemInfo(ctx, elem)
	}
	// Set all the storedMW
	for _, elem := range genState.StoredMWList {
		k.SetStoredMW(ctx, elem)
	}
	// Set all the storedTeam
	for _, elem := range genState.StoredTeamList {
		k.SetStoredTeam(ctx, elem)
	}
	// Set mwInfo
	k.SetMwInfo(ctx, genState.MwInfo)
	// Set teamInfo
	k.SetTeamInfo(ctx, genState.TeamInfo)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.SystemInfoList = k.GetAllSystemInfo(ctx)
	genesis.StoredMWList = k.GetAllStoredMW(ctx)
	genesis.StoredTeamList = k.GetAllStoredTeam(ctx)
	// Get all mwInfo
	mwInfo, found := k.GetMwInfo(ctx)
	if found {
		genesis.MwInfo = mwInfo
	}
	// Get all teamInfo
	teamInfo, found := k.GetTeamInfo(ctx)
	if found {
		genesis.TeamInfo = teamInfo
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
