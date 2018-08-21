package main_test

import (
	ws "deletestatecc"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
)

// TxID is just a dummuy transactional ID for test cases
const TxID = "mockTxID"

// MockStubUUID is an UUID value used accross all invocations to MockInit() method
const MockStubUUID = "d2490ad8-3901-11e8-b467-0ed5f89f718a"

var TestNamespaces = []string{"TestNamespace1", "TestNamespace2"}

// CheckInvoke is a common utilities method for test cases
func CheckInvoke(t *testing.T, stub *shim.MockStub, function string, args []byte) pb.Response {
	mockInvokeArgs := [][]byte{[]byte(function), args}
	res := stub.MockInvoke(MockStubUUID, mockInvokeArgs)
	if res.Status != shim.OK {
		fmt.Println("Invocation of", function, "function failed: ", string(res.Message))
		t.FailNow()
	}
	return res
}

func TestInit(t *testing.T) {
	scc := new(ws.DeleteStateCC)
	stub := shim.NewMockStub("TestInit", scc)
	res := initStub(stub)
	if res.Status != shim.OK {
		fmt.Println("Initialization of DeleteStateCC chaincode failed: ", string(res.Message))
		t.FailNow()
	}
	assert.True(t, reflect.DeepEqual(scc.Namespaces, TestNamespaces))
}

func initStub(stub *shim.MockStub) pb.Response {
	namespacesAsBytes := [][]byte{[]byte("init")}
	for _, namespace := range TestNamespaces {
		namespacesAsBytes = append(namespacesAsBytes, []byte(namespace))
	}
	res := stub.MockInit(MockStubUUID, namespacesAsBytes)
	return res
}

// TestResetWorldState tests the ResetWorldState method
func TestDeleteState(t *testing.T) {
	scc := new(ws.DeleteStateCC)
	stub := shim.NewMockStub("TestDeleteState", scc)
	initStub(stub)
	dummyRecord := `{"id": "{0}", "Company Code": "IBM"}`
	numberOfRecordsPerNamespace := 10
	// Store dummy data into world state
	for id := 0; id < numberOfRecordsPerNamespace; id++ {
		recordID := strconv.Itoa(id)
		record := strings.Replace(dummyRecord, "{0}", recordID, 1)
		recordAsBytes := []byte(record)
		for _, namespace := range TestNamespaces {
			recordCompositeKey, compositeErr := stub.CreateCompositeKey(namespace, []string{recordID})
			if compositeErr != nil {
				fmt.Println("Failed to generate composite key for record with id " + recordID + ". Error: " + compositeErr.Error())
				t.FailNow()
			}
			fmt.Println("Inserting dummy reconrd into namespace:", namespace)

			// Need a dummy transaction before we can call the stub.PutState() method
			stub.MockTransactionStart(TxID)
			stub.PutState(recordCompositeKey, recordAsBytes)
			// Insert additional data but using this time a non-composite key (these records should not be deleted)
			stub.PutState(recordID, recordAsBytes)
			stub.MockTransactionEnd(TxID)
		}
	}

	// Now we are ready to test our API to ensure it can delete records as expected
	rawNumberOfRecordsDeleted := CheckInvoke(t, stub, "DeleteState", []byte{})
	actualNumberOfRecordsDeleted := string(rawNumberOfRecordsDeleted.GetPayload())

	for id := 0; id < numberOfRecordsPerNamespace; id++ {
		recordID := strconv.Itoa(id)
		// Create composite key using namespace (prefix) for record
		for _, namespace := range TestNamespaces {

			recordCompositeKey, compositeErr := stub.CreateCompositeKey(namespace, []string{recordID})
			if compositeErr != nil {
				fmt.Println("Failed to generate composite key for record with id " + recordID + ".  Error: " + compositeErr.Error())
				t.FailNow()
			}

			// Get invoice data directly from world state (there should be none)
			recordAsBytesByCompositeKey := stub.State[recordCompositeKey]
			if recordAsBytesByCompositeKey != nil {
				fmt.Println("Failed to delete record with key:", recordCompositeKey)
				t.FailNow()
			}

			recordAsBytesBySimpleKey := stub.State[recordID]
			if recordAsBytesBySimpleKey == nil {
				fmt.Println("Failed to read record with key:", recordID)
				t.FailNow()
			}
		}
	}

	expectedNumberOfRecordsDeleted := numberOfRecordsPerNamespace * len(TestNamespaces)
	fmt.Printf("Summary: Expected number of deleted records = %d, actual number of deleted records from chain = %s \n ", expectedNumberOfRecordsDeleted, actualNumberOfRecordsDeleted)
	assert.Equal(t, strconv.Itoa(expectedNumberOfRecordsDeleted), actualNumberOfRecordsDeleted, "Number of deleted records do NOT match.")
}
