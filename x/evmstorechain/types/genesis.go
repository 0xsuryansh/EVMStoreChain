package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		VoteList: []Vote{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in vote
	voteIdMap := make(map[uint64]bool)
	voteCount := gs.GetVoteCount()
	for _, elem := range gs.VoteList {
		if _, ok := voteIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for vote")
		}
		if elem.Id >= voteCount {
			return fmt.Errorf("vote id should be lower or equal than the last id")
		}
		voteIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
