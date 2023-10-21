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

func (k Keeper) StoredTeamAll(c context.Context, req *types.QueryAllStoredTeamRequest) (*types.QueryAllStoredTeamResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedTeams []types.StoredTeam
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	storedTeamStore := prefix.NewStore(store, types.KeyPrefix(types.StoredTeamKeyPrefix))

	pageRes, err := query.Paginate(storedTeamStore, req.Pagination, func(key []byte, value []byte) error {
		var storedTeam types.StoredTeam
		if err := k.cdc.Unmarshal(value, &storedTeam); err != nil {
			return err
		}

		storedTeams = append(storedTeams, storedTeam)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredTeamResponse{StoredTeam: storedTeams, Pagination: pageRes}, nil
}

func (k Keeper) StoredTeam(c context.Context, req *types.QueryGetStoredTeamRequest) (*types.QueryGetStoredTeamResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStoredTeam(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredTeamResponse{StoredTeam: val}, nil
}
