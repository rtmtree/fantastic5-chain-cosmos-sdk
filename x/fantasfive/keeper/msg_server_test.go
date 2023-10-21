package keeper_test

import (
	"context"
	"testing"

	keepertest "fantasfive/testutil/keeper"
	"fantasfive/x/fantasfive"
	"fantasfive/x/fantasfive/keeper"
	"fantasfive/x/fantasfive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FantasfiveKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func setupMsgServerCreateTeam(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.FantasfiveKeeper(t)
	fantasfive.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}
func setupMsgServerAnnouceMw(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.FantasfiveKeeper(t)
	fantasfive.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}
