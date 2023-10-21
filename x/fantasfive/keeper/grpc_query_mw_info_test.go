package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "fantasfive/testutil/keeper"
	"fantasfive/testutil/nullify"
	"fantasfive/x/fantasfive/types"
)

func TestMwInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.FantasfiveKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestMwInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetMwInfoRequest
		response *types.QueryGetMwInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMwInfoRequest{},
			response: &types.QueryGetMwInfoResponse{MwInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MwInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
