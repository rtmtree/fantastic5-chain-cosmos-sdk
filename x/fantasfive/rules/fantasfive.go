package rules

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	PLAYER_LEN                 = 5
	POINT_PER_ASSIST      uint = 3
	POINT_PER_GOAL        uint = 6
	CAPTAIN_POINT_FACTOR       = 2 // factor to multiply the points of the captain
	DEFAULT_CAPTAIN_INDEX      = 0 // index of captain in player array, for now we'll just force it's the first player
)

type Team struct {
	Id           uint
	MWId         uint               // matchweek id
	Owner        string             // owner address/ reference to the owner
	Players      [PLAYER_LEN]string // array of player names, for now we'll use strings for an easy visualization
	CaptainIndex uint
	Points       int  // default -1
	Rank         uint // default 0
}

type PlayersPerformance struct {
	Players []string
	Goals   []uint
	Assists []uint
}

func init() {

}

func (team *Team) String() string {
	var buf bytes.Buffer

	//print ID, OWNER and MWID
	buf.WriteString(fmt.Sprintf("Id:%d Owner:%s MWId:%d", team.Id, team.Owner, team.MWId))

	//print PLAYERS
	for i, player := range team.Players {
		buf.WriteString(player)
		if i < (PLAYER_LEN - 1) {
			buf.WriteString("-")
		}
	}
	buf.WriteString(fmt.Sprintf(" Captain:%s", team.Players[team.CaptainIndex]))

	//print RANK and POINTS
	if team.Rank != 0 {
		buf.WriteString(fmt.Sprintf(" Rank: %d", team.Rank))
	} else {
		buf.WriteString(" Rank:Unranked")
	}
	if team.Points != -1 {
		buf.WriteString(fmt.Sprintf(" Points: %d", team.Points))
	} else {
		buf.WriteString(" Points:Unannounced")
	}
	return buf.String()
}

func Parse(teamAsString string) (*Team, error) {
	// reverse of String()

	// split by space
	players := [PLAYER_LEN]string{}
	captain_index := DEFAULT_CAPTAIN_INDEX
	team_id := -1
	team_owner := ""
	team_mwid := -1
	team_points := -1
	team_rank := 0
	for y, row := range strings.Split(teamAsString, " ") {
		if y < PLAYER_LEN {
			players[y] = row
			continue
		}

		// check if captain
		if strings.Contains(row, "Captain:") {
			captain := strings.Split(row, ":")
			if len(captain) != 2 {
				return nil, fmt.Errorf("invalid captain format")
			}
			captain_index = y
		}

		// check if rank
		if strings.Contains(row, "Rank:") {
			rank := strings.Split(row, ":")
			if len(rank) != 2 {
				return nil, fmt.Errorf("invalid rank format")
			}
			if rank[1] == "Unranked" {
				team_rank = 0
			} else {
				team_rank = y
			}
		}

		// check if points
		if strings.Contains(row, "Points:") {
			points := strings.Split(row, ":")
			if len(points) != 2 {
				return nil, fmt.Errorf("invalid points format")
			}
			if points[1] == "Unannounced" {
				team_points = -1
			} else {
				team_points = y
			}
		}

		// check if id
		if strings.Contains(row, "Id:") {
			id := strings.Split(row, ":")
			if len(id) != 2 {
				return nil, fmt.Errorf("invalid id format")
			}
			team_id = y
		}

		// check if owner
		if strings.Contains(row, "Owner:") {
			owner := strings.Split(row, ":")
			if len(owner) != 2 {
				return nil, fmt.Errorf("invalid owner format")
			}
			team_owner = owner[1]
		}

		// check if mwid
		if strings.Contains(row, "MWId:") {
			mwid := strings.Split(row, ":")
			if len(mwid) != 2 {
				return nil, fmt.Errorf("invalid mwid format")
			}
			team_mwid = y
		}

	}

	// return team
	return &Team{uint(team_id), uint(team_mwid), team_owner, players, uint(captain_index), team_points, uint(team_rank)}, nil
}

func (team *Team) StringPlayers() string {
	var buf bytes.Buffer

	//print PLAYERS
	for i, player := range team.Players {
		buf.WriteString(player)
		if i < (PLAYER_LEN - 1) {
			buf.WriteString("-")
		}
	}
	return buf.String()
}

func ParsePlayers(teamAsString string) (*[PLAYER_LEN]string, error) {
	// reverse of StringPlayers()

	// split by space
	players := [PLAYER_LEN]string{}
	for y, row := range strings.Split(teamAsString, "-") {
		if y < PLAYER_LEN {
			players[y] = row
			continue
		}
	}
	return &players, nil
}

func NewTeam(id uint, matchweekId uint, owner string, players [PLAYER_LEN]string) (*Team, error) {
	// in the future we can add captain index as a parameter
	team := &Team{id, matchweekId, owner, players, DEFAULT_CAPTAIN_INDEX, -1, 0}
	if err := team.Valid(); err != nil {
		return nil, err
	}
	return team, nil
}

func ParseMatchWeekId(matchweekIdAsString string) (uint, error) {
	// string to uint
	idAsU64, err := strconv.ParseUint(matchweekIdAsString, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(idAsU64), nil
}

func ParseTeamId(teamIdAsString string) (uint, error) {
	// string to uint
	idAsU64, err := strconv.ParseUint(teamIdAsString, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(idAsU64), nil
}
func ParsePoints(pointsAsString string) (int, error) {
	// string to int
	pointsAsI64, err := strconv.ParseInt(pointsAsString, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(pointsAsI64), nil
}
func ParseRank(rankAsString string) (uint, error) {
	// string to uint
	rankAsU64, err := strconv.ParseUint(rankAsString, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(rankAsU64), nil
}

func (team *Team) StringPoints() string {
	return fmt.Sprintf("%d", team.Points)
}
func (team *Team) StringRank() string {
	return fmt.Sprintf("%d", team.Rank)
}
func (team *Team) StringTeamId() string {
	return fmt.Sprintf("%d", team.Id)
}
func (team *Team) StringMatchWeekId() string {
	return fmt.Sprintf("%d", team.MWId)
}
func (team *Team) StringCaptainIndex() string {
	return fmt.Sprintf("%d", team.CaptainIndex)
}

func StringFromUint(input uint) string {
	return fmt.Sprintf("%d", input)
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
	if team.CaptainIndex >= PLAYER_LEN {
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

	return nil
}

func (playersperformance *PlayersPerformance) CalculatePointByTeam(team Team) uint {
	// calculate points for each player
	points := uint(0)
	for i, player := range team.Players {
		point := playersperformance.CalculatePointByPlayer(player)
		// if the player is the captain, multiply the points by CAPTAIN_POINT_FACTOR
		if uint(i) == team.CaptainIndex {
			point = point * CAPTAIN_POINT_FACTOR
		}
		points += point
	}
	return points
}

func (playersperformance *PlayersPerformance) CalculatePointByPlayer(player string) uint {
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

func CalculateRankByPoints(points []uint) ([]uint, error) {
	// calculate rank
	// if points are equal, sort by index, we consider the earlier team deserves a higher rank
	rankList := make([]uint, len(points))
	for i, point := range points {
		rank := uint(1)
		for j, p := range points {
			if i != j {
				if point < p {
					rank++
				} else if point == p && i > j {
					rank++
				}
			}
		}
		rankList[i] = rank
	}
	return rankList, nil
}

func NewPlayersPerformanceFromString(playersAsString string) (*PlayersPerformance, error) {
	playerPerf, err := ParsePlayersPerformance(playersAsString)
	if err != nil {
		return nil, err
	}
	return playerPerf, nil
}

// Let use pattern of "player0-goal0-assist0-player1-goal1-assist1-..."
func ParsePlayersPerformance(playersAsString string) (*PlayersPerformance, error) {
	// split by space
	players := []string{}
	goals := []uint{}
	assists := []uint{}
	for y, row := range strings.Split(playersAsString, "-") {
		if y%3 == 0 {
			players = append(players, row)
		} else if y%3 == 1 {
			goal, err := strconv.ParseUint(row, 10, 32)
			if err != nil {
				return nil, err
			}
			goals = append(goals, uint(goal))
		} else {
			assist, err := strconv.ParseUint(row, 10, 32)
			if err != nil {
				return nil, err
			}
			assists = append(assists, uint(assist))
		}
	}

	// check if 3 lists have the same length
	if len(players) != len(goals) || len(players) != len(assists) {
		return nil, fmt.Errorf("invalid players performance length")
	}

	return &PlayersPerformance{players, goals, assists}, nil
}

func ParseCaptainIndex(captainIndex string) (uint, error) {
	result, err := strconv.ParseUint(captainIndex, 10, 32)
	return uint(result), err
}
