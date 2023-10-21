package keeper

import (
	"context"
	"fmt"
	"strconv"

	"fantasfive/x/fantasfive/rules"
	"fantasfive/x/fantasfive/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AnnounceAndCreateNextMw(goCtx context.Context, msg *types.MsgAnnounceAndCreateNextMw) (*types.MsgAnnounceAndCreateNextMwResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	mwInfo, found := k.Keeper.GetMwInfo(ctx)
	if !found {
		panic("mwInfo not found")
	}
	teamInfo, found := k.Keeper.GetTeamInfo(ctx)
	if !found {
		panic("teamInfo not found")
	}

	playerPerf, err := rules.ParsePlayersPerformance(msg.PlayerPerf)
	if err != nil {
		panic(err)
	}

	// cache team ids, points
	teamIds := []uint{}
	teamPoints := []uint{}

	lastTeamId := uint(teamInfo.NextId) - 1
	// panic(fmt.Sprintf("lastTeamId: %d", lastTeamId))
	// go from last team to first team
	for i := lastTeamId; true; i-- {
		teamId := strconv.FormatUint(uint64(i), 10)
		team, found := k.Keeper.GetStoredTeam(ctx, teamId)
		if !found {
			panic(fmt.Sprintf("team %d/%s not found", uint64(i), teamId))
		}
		teamMwId, err := rules.ParseMatchWeekId(team.MwId)
		if err != nil {
			panic(err)
		}
		if teamMwId == uint(mwInfo.NextId) {
			//assume team is valid b/c it was validated when it was created
			team_players, _ := rules.ParsePlayers(team.Players)
			team_cap_index, _ := rules.ParseCaptainIndex(team.CaptainIndex)

			team_points := (playerPerf.CalculatePointByTeam(rules.Team{
				Players:      *team_players,
				CaptainIndex: team_cap_index,
			}))

			teamIds = append(teamIds, uint(i))
			teamPoints = append(teamPoints, team_points)
		}

		// stop when we reach the first team of the matchweek
		if teamMwId != uint(mwInfo.NextId) || i == 0 {
			break
		}
	}

	// cache team ranks
	teamRank, err := rules.CalculateRankByPoints(teamPoints)
	if err != nil {
		panic(err)
	}

	// save team points and ranks to store by team id
	for i, teamId := range teamIds {
		team, found := k.Keeper.GetStoredTeam(ctx, strconv.FormatUint(uint64(teamId), 10))
		if !found {
			panic("team not found when saving points and ranks")
		}
		team.Points = rules.StringFromUint(teamPoints[i])
		team.Rank = rules.StringFromUint(teamRank[i])

		k.Keeper.SetStoredTeam(ctx, team)
	}

	// update mwInfo
	mwInfo.NextId++
	nextMwId := uint(mwInfo.NextId)
	k.Keeper.SetMwInfo(ctx, mwInfo)

	return &types.MsgAnnounceAndCreateNextMwResponse{
		NextMwId: rules.StringFromUint(nextMwId),
	}, nil
}
