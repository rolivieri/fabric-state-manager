package main_test

import (
	main "DeleteStateCC"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// TxID is just a dummuy transactional ID for test cases
const TxID = "mockTxID"

// MockStubUUID is an UUID value used accross all invocations to MockInit() method
const MockStubUUID = "d2490ad8-3901-11e8-b467-0ed5f89f718a"

const TestNamespace = "TestNamespace"

func TestInit(t *testing.T) {
	scc := new(main.DeleteStateCC)
	stub := shim.NewMockStub("TestInit", scc)
	namespaces := [][]byte{[]byte("init"), []byte(TestNamespace)}
	res := stub.MockInit(MockStubUUID, namespaces)
	if res.Status != shim.OK {
		fmt.Println("Initialization of chaincode failed: ", string(res.Message))
		t.FailNow()
	}
	// TODO: Assertions
}

// TestResetWorldState tests the ResetWorldState method
func TestDeleteState(t *testing.T) {
	scc := new(main.DeleteStateCC)
	stub := shim.NewMockStub("TestDeleteState", scc)
	dummyRecord := `{"id": "{0}", "Company Code": "IBM"}`
	expectedNumOfRecordsDeleted := 10
	// Store dummy data into world state
	for id := 0; id < expectedNumOfRecordsDeleted; id++ {
		recordID := strconv.Itoa(id)
		record := strings.Replace(dummyRecord, "{0}", recordID, 1)
		recordAsBytes := []byte(record)
		recordCompositeKey, compositeErr := stub.CreateCompositeKey(TestNamespace, []string{recordID})
		if compositeErr != nil {
			fmt.Println("Failed to generate composite key for record with id " + recordID + ".  Error: " + compositeErr.Error())
			t.FailNow()
		}

		// Need a dummy transaction before we can call the stub.PutState() method
		stub.MockTransactionStart(TxID)
		stub.PutState(recordCompositeKey, recordAsBytes)
		// Insert additional data but using this time a non-composite key (these records should not be deleted)
		stub.PutState(recordID, recordAsBytes)
		stub.MockTransactionEnd(TxID)
	}

	//TODO: Assertions
}
