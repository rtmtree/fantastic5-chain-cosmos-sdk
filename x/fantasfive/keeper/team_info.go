package keeper

import (
	"fantasfive/x/fantasfive/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTeamInfo set teamInfo in the store
func (k Keeper) SetTeamInfo(ctx sdk.Context, teamInfo types.TeamInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TeamInfoKey))
	b := k.cdc.MustMarshal(&teamInfo)
	store.Set([]byte{0}, b)
}

// GetTeamInfo returns teamInfo
func (k Keeper) GetTeamInfo(ctx sdk.Context) (val types.TeamInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TeamInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTeamInfo removes teamInfo from the store
func (k Keeper) RemoveTeamInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TeamInfoKey))
	store.Delete([]byte{0})
}
