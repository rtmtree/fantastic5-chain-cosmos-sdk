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

func (k Keeper) SystemInfoAll(c context.Context, req *types.QueryAllSystemInfoRequest) (*types.QueryAllSystemInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var systemInfos []types.SystemInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	systemInfoStore := prefix.NewStore(store, types.KeyPrefix(types.SystemInfoKeyPrefix))

	pageRes, err := query.Paginate(systemInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var systemInfo types.SystemInfo
		if err := k.cdc.Unmarshal(value, &systemInfo); err != nil {
			return err
		}

		systemInfos = append(systemInfos, systemInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSystemInfoResponse{SystemInfo: systemInfos, Pagination: pageRes}, nil
}

func (k Keeper) SystemInfo(c context.Context, req *types.QueryGetSystemInfoRequest) (*types.QueryGetSystemInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSystemInfo(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSystemInfoResponse{SystemInfo: val}, nil
}
