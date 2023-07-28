package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		KycList:  []Kyc{},
		Kyc2List: []Kyc2{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in kyc
	kycIndexMap := make(map[string]struct{})

	for _, elem := range gs.KycList {
		index := string(KycKey(elem.Index))
		if _, ok := kycIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for kyc")
		}
		kycIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in kyc2
	kyc2IndexMap := make(map[string]struct{})

	for _, elem := range gs.Kyc2List {
		index := string(Kyc2Key(elem.Address))
		if _, ok := kyc2IndexMap[index]; ok {
			return fmt.Errorf("duplicated index for kyc2")
		}
		kyc2IndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
