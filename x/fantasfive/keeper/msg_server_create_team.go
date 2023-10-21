package keeper

import (
	"context"

	"fantasfive/x/fantasfive/rules"
	"fantasfive/x/fantasfive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTeam(goCtx context.Context, msg *types.MsgCreateTeam) (*types.MsgCreateTeamResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	mwInfo, found := k.Keeper.GetMwInfo(ctx)
	if !found {
		panic("mwInfo not found")
	}

	teamInfo, found := k.Keeper.GetTeamInfo(ctx)
	if !found {
		panic("teamInfo not found")
	}

	team, err := rules.NewTeam(uint(teamInfo.NextId), uint(mwInfo.NextId), msg.Creator, [rules.PLAYER_LEN]string{msg.Player0, msg.Player1, msg.Player2, msg.Player3, msg.Player4})
	if err != nil {
		return nil, err
	}

	storedTeam := types.StoredTeam{
		CaptainIndex: team.StringCaptainIndex(),
		Index:        team.StringTeamId(),
		MwId:         team.StringMatchWeekId(),
		Owner:        team.Owner,
		Players:      team.StringPlayers(),
		Points:       team.StringPoints(),
		Rank:         team.StringRank(),
	}

	err = storedTeam.Validate()
	if err != nil {
		return nil, err
	}

	k.Keeper.SetStoredTeam(ctx, storedTeam)
	teamInfo.NextId++
	k.Keeper.SetTeamInfo(ctx, teamInfo)

	return &types.MsgCreateTeamResponse{
		TeamId: team.StringTeamId(),
	}, nil
}
