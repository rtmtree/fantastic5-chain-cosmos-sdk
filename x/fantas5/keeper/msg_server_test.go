package keeper_test

import (
	"context"
	"testing"

	keepertest "fantas5/testutil/keeper"
	"fantas5/x/fantas5/keeper"
	"fantas5/x/fantas5/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Fantas5Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
