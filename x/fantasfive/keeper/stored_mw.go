package keeper

import (
	"fantasfive/x/fantasfive/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredMW set a specific storedMW in the store from its index
func (k Keeper) SetStoredMW(ctx sdk.Context, storedMW types.StoredMW) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMWKeyPrefix))
	b := k.cdc.MustMarshal(&storedMW)
	store.Set(types.StoredMWKey(
		storedMW.Index,
	), b)
}

// GetStoredMW returns a storedMW from its index
func (k Keeper) GetStoredMW(
	ctx sdk.Context,
	index string,

) (val types.StoredMW, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMWKeyPrefix))

	b := store.Get(types.StoredMWKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredMW removes a storedMW from the store
func (k Keeper) RemoveStoredMW(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMWKeyPrefix))
	store.Delete(types.StoredMWKey(
		index,
	))
}

// GetAllStoredMW returns all storedMW
func (k Keeper) GetAllStoredMW(ctx sdk.Context) (list []types.StoredMW) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredMWKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredMW
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
