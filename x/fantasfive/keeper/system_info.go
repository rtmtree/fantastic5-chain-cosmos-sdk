package keeper

import (
	"fantasfive/x/fantasfive/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSystemInfo set a specific systemInfo in the store from its index
func (k Keeper) SetSystemInfo(ctx sdk.Context, systemInfo types.SystemInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKeyPrefix))
	b := k.cdc.MustMarshal(&systemInfo)
	store.Set(types.SystemInfoKey(
		systemInfo.Index,
	), b)
}

// GetSystemInfo returns a systemInfo from its index
func (k Keeper) GetSystemInfo(
	ctx sdk.Context,
	index string,

) (val types.SystemInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKeyPrefix))

	b := store.Get(types.SystemInfoKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSystemInfo removes a systemInfo from the store
func (k Keeper) RemoveSystemInfo(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKeyPrefix))
	store.Delete(types.SystemInfoKey(
		index,
	))
}

// GetAllSystemInfo returns all systemInfo
func (k Keeper) GetAllSystemInfo(ctx sdk.Context) (list []types.SystemInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SystemInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
