package main

import (
	sm "deleteStateCC/statemanager"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(sm.DeleteStateCC))
	if err != nil {
		shim.NewLogger("DeleteStateCCMainLog").Errorf("Error starting DeleteStateCC: %s", err)
	}
}
