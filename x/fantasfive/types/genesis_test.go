package types_test

import (
	"testing"

	"fantasfive/x/fantasfive/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				SystemInfoList: []types.SystemInfo{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				StoredMWList: []types.StoredMW{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				StoredTeamList: []types.StoredTeam{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				MwInfo: &types.MwInfo{
					NextId: 96,
				},
				TeamInfo: &types.TeamInfo{
					NextId: 4,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated systemInfo",
			genState: &types.GenesisState{
				SystemInfoList: []types.SystemInfo{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated storedMW",
			genState: &types.GenesisState{
				StoredMWList: []types.StoredMW{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated storedTeam",
			genState: &types.GenesisState{
				StoredTeamList: []types.StoredTeam{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
