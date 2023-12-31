package types

import (
	"fmt"
)

// Default index
const DefaultMwIndex uint64 = 1
const DefaultTeamIndex uint64 = 0

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SystemInfoList: []SystemInfo{},
		StoredMWList:   []StoredMW{},
		StoredTeamList: []StoredTeam{},
		MwInfo: MwInfo{
			NextId: uint64(DefaultMwIndex),
		},
		TeamInfo: TeamInfo{
			NextId: uint64(DefaultTeamIndex),
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in systemInfo
	systemInfoIndexMap := make(map[string]struct{})

	for _, elem := range gs.SystemInfoList {
		index := string(SystemInfoKey(elem.Index))
		if _, ok := systemInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for systemInfo")
		}
		systemInfoIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in storedMW
	storedMWIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredMWList {
		index := string(StoredMWKey(elem.Index))
		if _, ok := storedMWIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedMW")
		}
		storedMWIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in storedTeam
	storedTeamIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredTeamList {
		index := string(StoredTeamKey(elem.Index))
		if _, ok := storedTeamIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedTeam")
		}
		storedTeamIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
