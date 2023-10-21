package keeper

import (
	"context"

	"fantasfive/x/fantasfive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MwInfo(c context.Context, req *types.QueryGetMwInfoRequest) (*types.QueryGetMwInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMwInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMwInfoResponse{MwInfo: val}, nil
}
