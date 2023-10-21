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

func TestTeamInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.FantasfiveKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTeamInfo(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetTeamInfoRequest
		response *types.QueryGetTeamInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTeamInfoRequest{},
			response: &types.QueryGetTeamInfoResponse{TeamInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TeamInfo(wctx, tc.request)
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
