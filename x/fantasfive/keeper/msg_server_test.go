package keeper_test

import (
	"context"
	"testing"

	keepertest "fantasfive/testutil/keeper"
	"fantasfive/x/fantasfive/keeper"
	"fantasfive/x/fantasfive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FantasFiveKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
