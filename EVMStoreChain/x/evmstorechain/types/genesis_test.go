package types_test

import (
	"testing"

	"EVMStoreChain/x/evmstorechain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
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

				VoteList: []types.Vote{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				VoteCount: 2,
				BlockstoragestateList: []types.Blockstoragestate{
					{
						Blocknumber: "0",
					},
					{
						Blocknumber: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated vote",
			genState: &types.GenesisState{
				VoteList: []types.Vote{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid vote count",
			genState: &types.GenesisState{
				VoteList: []types.Vote{
					{
						Id: 1,
					},
				},
				VoteCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated blockstoragestate",
			genState: &types.GenesisState{
				BlockstoragestateList: []types.Blockstoragestate{
					{
						Blocknumber: "0",
					},
					{
						Blocknumber: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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
