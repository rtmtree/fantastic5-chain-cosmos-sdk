package keeper

import (
	"fantasfive/x/fantasfive/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetMwInfo set mwInfo in the store
func (k Keeper) SetMwInfo(ctx sdk.Context, mwInfo types.MwInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MwInfoKey))
	b := k.cdc.MustMarshal(&mwInfo)
	store.Set([]byte{0}, b)
}

// GetMwInfo returns mwInfo
func (k Keeper) GetMwInfo(ctx sdk.Context) (val types.MwInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MwInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMwInfo removes mwInfo from the store
func (k Keeper) RemoveMwInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MwInfoKey))
	store.Delete([]byte{0})
}
