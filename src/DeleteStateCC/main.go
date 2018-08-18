package deletestate

import "github.com/hyperledger/fabric/core/chaincode/shim"

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(DeleteStateCC))
	if err != nil {
		logger.Errorf("Error starting DeleteStateCC: %s", err)
	}
}
