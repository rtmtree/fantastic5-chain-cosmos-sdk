package rules_test

import (
	"fantasfive/x/fantasfive/rules"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsePlayerPerf(t *testing.T) {
	tests := []struct {
		name       string
		playerPerf string
		err        error
		expected   *rules.PlayersPerformance
	}{
		{
			name:       "invalid player performance",
			playerPerf: "Ronaldo-0--9-Messi-1-8",
			err:        fmt.Errorf("strconv.ParseUint: parsing \"\": invalid syntax"),
			expected:   &rules.PlayersPerformance{},
		},
		{
			name:       "invalid player performance",
			playerPerf: "Ronaldo-a-9-Messi-1-8",
			err:        fmt.Errorf("strconv.ParseUint: parsing \"a\": invalid syntax"),
			expected:   &rules.PlayersPerformance{},
		},
		{
			name:       "invalid player performance",
			playerPerf: "Ronaldo-1-2-Messi-1",
			err:        fmt.Errorf("invalid players performance length"),
			expected:   &rules.PlayersPerformance{},
		},
		{
			name:       "valid player performance",
			playerPerf: "Ronaldo-0-9-Messi-1-8",
			err:        nil,
			expected: &rules.PlayersPerformance{
				Players: []string{
					"Ronaldo",
					"Messi",
				},
				Goals: []uint{
					0,
					1,
				},
				Assists: []uint{
					9,
					8,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := rules.ParsePlayersPerformance(tt.playerPerf)
			if tt.err != nil {
				require.Equal(t, tt.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculatePointByTeam(t *testing.T) {
	playerPerf := &rules.PlayersPerformance{
		Players: []string{
			"Ronaldo",
			"Messi",
			"Neymar",
			"Salah",
			"Kane",
			"Mane",
			"Son",
			"De Bruyne",
			"Kimmich",
			"Robertson",
			"Ramos",
		},
		Goals: []uint{
			0,
			1,
			0,
			3,
			3,
			9,
			2,
			1,
			0,
			3,
			0,
		},
		Assists: []uint{
			1,
			3,
			2,
			1,
			3,
			0,
			2,
			1,
			3,
			0,
			0,
		},
	}

	testTeams := struct {
		name     string
		teams    []rules.Team
		expected []uint
	}{
		name: "valid teams",
		teams: []rules.Team{
			rules.Team{
				Players: [rules.PLAYER_LEN]string{
					"Ronaldo",
					"Messi",
					"Neymar",
					"Salah",
					"Kane",
				},
				CaptainIndex: 0,
			},
			rules.Team{
				Players: [rules.PLAYER_LEN]string{
					"Ronaldo",
					"Messi",
					"Neymar",
					"Salah",
					"Ramos",
				},
				CaptainIndex: 0,
			},
			rules.Team{
				Players: [rules.PLAYER_LEN]string{
					"De Bruyne",
					"Messi",
					"Neymar",
					"Salah",
					"Ramos",
				},
				CaptainIndex: 0,
			},
		},
		expected: []uint{
			6 + 15 + 6 + 21 + 27,
			6 + 15 + 6 + 21 + 0,
			18 + 15 + 6 + 21 + 0,
		},
	}

	t.Run(testTeams.name, func(t *testing.T) {
		for i, team := range testTeams.teams {
			result := playerPerf.CalculatePointByTeam(team)
			require.Equal(t, testTeams.expected[i], result)
		}
	})
}

func TestCalculatePointByTeam2(t *testing.T) {
	playerPerf := &rules.PlayersPerformance{
		Players: []string{
			"Son",
			"Kane",
			"Ric",
			"Hal",
		},
		Goals: []uint{
			4,
			1,
			1,
			9,
		},
		Assists: []uint{
			9,
			0,
			1,
			3,
		},
	}

	testTeams := struct {
		name     string
		teams    []rules.Team
		expected []uint
	}{
		name: "valid teams",
		teams: []rules.Team{
			rules.Team{
				Players: [rules.PLAYER_LEN]string{
					"Hal",
					"Kev",
					"Kyl",
					"Rub",
					"Ede",
				},
				CaptainIndex: 0,
			},
			rules.Team{
				Players: [rules.PLAYER_LEN]string{
					"Hal",
					"Kev",
					"San",
					"Kyl",
					"Ede",
				},
				CaptainIndex: 0,
			},
			rules.Team{
				Players: [rules.PLAYER_LEN]string{
					"Hal",
					"Kev",
					"San",
					"Ric",
					"Ede",
				},
				CaptainIndex: 0,
			},
		},
		expected: []uint{
			126,
			126,
			135,
		},
	}

	t.Run(testTeams.name, func(t *testing.T) {
		for i, team := range testTeams.teams {
			result := playerPerf.CalculatePointByTeam(team)
			require.Equal(t, testTeams.expected[i], result)
		}
	})
}
