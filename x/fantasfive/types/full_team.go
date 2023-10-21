package types

import (
	"fantasfive/x/fantasfive/rules"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedTeam StoredTeam) GetTeamOwnerAddress() (black sdk.AccAddress, err error) {
	owner, errOwner := sdk.AccAddressFromBech32(storedTeam.Owner)
	return owner, sdkerrors.Wrapf(errOwner, ErrOwner.Error(), storedTeam.Owner)
}

func (storedTeam StoredTeam) ParseTeam() (Team *rules.Team, err error) {
	players, err := rules.ParsePlayers(storedTeam.Players)
	if err != nil {
		return nil, err
	}
	mWId, err := rules.ParseMatchWeekId(storedTeam.MwId)
	if err != nil {
		return nil, fmt.Errorf("error parsing match week id: %s", storedTeam.MwId)
	}
	teamId, err := rules.ParseTeamId(storedTeam.Index)
	if err != nil {
		return nil, fmt.Errorf("error parsing match week id: %s", storedTeam.Index)
	}
	_, err = rules.NewTeam(teamId, mWId, storedTeam.Owner, *players)
	if err != nil {
		return nil, err
	}

	points, err := rules.ParsePoints(storedTeam.Points)
	if err != nil {
		return nil, err
	}
	rank, err := rules.ParseRank(storedTeam.Rank)
	if err != nil {
		return nil, err
	}

	team := &rules.Team{
		Id:      teamId,
		MWId:    mWId,
		Owner:   storedTeam.Owner,
		Players: *players,
		Points:  points,
		Rank:    rank,
	}

	return team, nil
}

func (storedTeam StoredTeam) Validate() (err error) {
	_, err = storedTeam.GetTeamOwnerAddress()
	if err != nil {
		return err
	}
	_, err = storedTeam.ParseTeam()
	return err
}
