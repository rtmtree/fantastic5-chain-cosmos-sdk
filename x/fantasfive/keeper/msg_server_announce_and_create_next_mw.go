package keeper

import (
	"context"

	"fantasfive/x/fantasfive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AnnounceAndCreateNextMw(goCtx context.Context, msg *types.MsgAnnounceAndCreateNextMw) (*types.MsgAnnounceAndCreateNextMwResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAnnounceAndCreateNextMwResponse{}, nil
}
