package keeper

import (
	"fantasfive/x/fantasfive/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredTeam set a specific storedTeam in the store from its index
func (k Keeper) SetStoredTeam(ctx sdk.Context, storedTeam types.StoredTeam) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredTeamKeyPrefix))
	b := k.cdc.MustMarshal(&storedTeam)
	store.Set(types.StoredTeamKey(
		storedTeam.Index,
	), b)
}

// GetStoredTeam returns a storedTeam from its index
func (k Keeper) GetStoredTeam(
	ctx sdk.Context,
	index string,

) (val types.StoredTeam, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredTeamKeyPrefix))

	b := store.Get(types.StoredTeamKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredTeam removes a storedTeam from the store
func (k Keeper) RemoveStoredTeam(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredTeamKeyPrefix))
	store.Delete(types.StoredTeamKey(
		index,
	))
}

// GetAllStoredTeam returns all storedTeam
func (k Keeper) GetAllStoredTeam(ctx sdk.Context) (list []types.StoredTeam) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredTeamKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredTeam
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
