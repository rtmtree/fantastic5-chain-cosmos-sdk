package keeper_test

import (
	"testing"

	testkeeper "fantas5/testutil/keeper"
	"fantas5/x/fantas5/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Fantas5Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
