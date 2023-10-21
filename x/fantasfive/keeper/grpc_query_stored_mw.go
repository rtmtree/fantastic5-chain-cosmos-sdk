package keeper

import (
	"context"

	"fantasfive/x/fantasfive/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredMWAll(c context.Context, req *types.QueryAllStoredMWRequest) (*types.QueryAllStoredMWResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedMWs []types.StoredMW
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedMWStore := prefix.NewStore(store, types.KeyPrefix(types.StoredMWKeyPrefix))

	pageRes, err := query.Paginate(storedMWStore, req.Pagination, func(key []byte, value []byte) error {
		var storedMW types.StoredMW
		if err := k.cdc.Unmarshal(value, &storedMW); err != nil {
			return err
		}

		storedMWs = append(storedMWs, storedMW)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredMWResponse{StoredMW: storedMWs, Pagination: pageRes}, nil
}

func (k Keeper) StoredMW(c context.Context, req *types.QueryGetStoredMWRequest) (*types.QueryGetStoredMWResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredMW(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredMWResponse{StoredMW: val}, nil
}
