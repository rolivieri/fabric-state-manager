package main

import (
	ws "deletestatecc"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(ws.DeleteStateCC))
	if err != nil {
		shim.NewLogger("DeleteStateCCMainLog").Errorf("Error starting DeleteStateCC: %s", err)
	}
}
