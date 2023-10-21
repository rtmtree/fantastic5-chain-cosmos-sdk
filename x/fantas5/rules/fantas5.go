package rules

import (
	"bytes"
	"fmt"
)

const (
	PLAYER_LEN            = 5
	POINT_PER_ASSIST      = 3
	POINT_PER_GOAL        = 6
	CAPTAIN_POINT_FACTOR  = 2 // factor to multiply the points of the captain
	DEFAULT_CAPTAIN_INDEX = 0 // index of captain in player array, for now we'll just force it's the first player
)

type Team struct {
	Id           int
	MWId         int                // matchweek id
	Players      [PLAYER_LEN]string // array of player names, for now we'll use strings for an easy visualization
	CaptainIndex int
	Points       int // default -1
	Rank         int // default -1
}

type PlayersPerformance struct {
	Players []string
	Goals   []int
	Assists []int
}

func init() {

}

func (team *Team) String() string {
	var buf bytes.Buffer

	//print PLAYERS
	for i, player := range team.Players {
		buf.WriteString(player)
		if i < (PLAYER_LEN - 1) {
			buf.WriteString(" ")
		}
	}
	buf.WriteString(fmt.Sprintf(" Captain: %s", team.Players[team.CaptainIndex]))

	//print RANK and POINTS
	if team.Rank != -1 {
		buf.WriteString(fmt.Sprintf(" Rank: %d", team.Rank))
	} else {
		buf.WriteString(" Rank: Unranked")
	}
	if team.Points != -1 {
		buf.WriteString(fmt.Sprintf(" Points: %d", team.Points))
	} else {
		buf.WriteString(" Points: Unannounced")
	}
	return buf.String()
}

func New(id int, matchweekId int, players [PLAYER_LEN]string) (*Team, error) {
	// in the future we can add captain index as a parameter
	team := &Team{id, matchweekId, players, DEFAULT_CAPTAIN_INDEX, -1, -1}
	if err := team.Valid(); err != nil {
		return nil, err
	}
	return team, nil
}

func (team *Team) Valid() error {
	//check if there are 5 players
	if len(team.Players) != PLAYER_LEN {
		return fmt.Errorf("invalid number of players")
	}

	//check if there is a player with the same name in the team
	for i, player := range team.Players {
		for j, p := range team.Players {
			if i != j && player == p {
				return fmt.Errorf("duplicated player")
			}
		}
	}

	//check if captain index is valid
	if team.CaptainIndex < 0 || team.CaptainIndex >= PLAYER_LEN {
		return fmt.Errorf("invalid captain index")
	}

	//in the future we can add more validations here

	return nil
}

func (playersperformance *PlayersPerformance) String() string {
	var buf bytes.Buffer
	for i, player := range playersperformance.Players {
		buf.WriteString(player)
		buf.WriteString(" ")
		buf.WriteString("G:")
		buf.WriteString(fmt.Sprintf("%d", playersperformance.Goals[i]))
		buf.WriteString(" ")
		buf.WriteString("A:")
		buf.WriteString(fmt.Sprintf("%d", playersperformance.Assists[i]))
		if i < (PLAYER_LEN - 1) {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}

func (playersperformance *PlayersPerformance) ValidPlayersPerformance() error {
	//check if there is a duplicared player with the same name in the players
	for i, player := range playersperformance.Players {
		for j, p := range playersperformance.Players {
			if i != j && player == p {
				return fmt.Errorf("duplicated player")
			}
		}
	}

	//check if all lists have the same length
	if len(playersperformance.Players) != len(playersperformance.Goals) || len(playersperformance.Players) != len(playersperformance.Assists) {
		return fmt.Errorf("invalid players performance length")
	}

	//check if all goals and assists are positive
	for i, goal := range playersperformance.Goals {
		if goal < 0 {
			return fmt.Errorf("invalid goal value for player %s", playersperformance.Players[i])
		}
	}
	for i, assist := range playersperformance.Assists {
		if assist < 0 {
			return fmt.Errorf("invalid assist value for player %s", playersperformance.Players[i])
		}
	}

	return nil
}

func (playersperformance *PlayersPerformance) CalculatePointByTeam(team Team) int {
	// calculate points for each player
	points := 0
	for i, player := range team.Players {
		point := playersperformance.CalculatePointByPlayer(player)
		// if the player is the captain, multiply the points by CAPTAIN_POINT_FACTOR
		if i == team.CaptainIndex {
			point = point * CAPTAIN_POINT_FACTOR
		}
		points = points + point
	}
	return points
}

func (playersperformance *PlayersPerformance) CalculatePointByPlayer(player string) int {
	// calculate points for a player
	for j, p := range playersperformance.Players {
		if player == p {
			return (playersperformance.Goals[j] * POINT_PER_GOAL) + (playersperformance.Assists[j] * POINT_PER_ASSIST)
		}
	}
	return 0
}

func CalculateRank(teams []Team) ([]int, error) {
	// check if all teams have points
	for _, team := range teams {
		if team.Points == -1 {
			return nil, fmt.Errorf("invalid points for team %s", team.String())
		}
	}

	// calculate rank
	// if points are equal, sort by index, we consider the earlier team deserves a higher rank
	rankList := make([]int, len(teams))
	for i, team := range teams {
		rank := 1
		for j, t := range teams {
			if i != j {
				if team.Points < t.Points {
					rank++
				} else if team.Points == t.Points && i > j {
					rank++
				}
			}
		}
		rankList[i] = rank
	}
	return rankList, nil
}
