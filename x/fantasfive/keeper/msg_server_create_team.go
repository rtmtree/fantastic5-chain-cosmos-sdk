package keeper

import (
	"context"

	"fantasfive/x/fantasfive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTeam(goCtx context.Context, msg *types.MsgCreateTeam) (*types.MsgCreateTeamResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateTeamResponse{}, nil
}
