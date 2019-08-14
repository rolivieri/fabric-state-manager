package statemanager

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("RemoverCCLog")

// Key for storing namespaces on the blockchain
const Key = "fabric-state-manager"

// RemoverCC chaincode structure
type RemoverCC struct {
}

// Init initializes chaincode
func (t *RemoverCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### RemoverCC Init ###########")
	methodName := "Init()"
	_, args := stub.GetFunctionAndParameters()
	t.Initialize(stub, args)
	logger.Infof("- End execution -  %s\n", methodName)
	return shim.Success(nil)
}

// Initialize initializes chaincode
func (t *RemoverCC) Initialize(stub shim.ChaincodeStubInterface, namespaces []string) pb.Response {
	logger.Info("########### RemoverCC Initialize ###########")
	methodName := "Initialize()"

	if len(namespaces) == 0 {
		warningMsg := fmt.Sprintf("%s - No namespaces were provided to RemoverCC.", methodName)
		logger.Warning(warningMsg)
	}

	// Store namespaces on the blockchain
	namespacesAsBytes := []byte(strings.Join(namespaces, ","))
	err := stub.PutState(Key, namespacesAsBytes)
	if err != nil {
		errorMsg := fmt.Sprintf("Error storing namespaces: %s", err.Error())
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}

	logger.Infof("%s - Namespaces provided to RemoverCC: %v", methodName, namespaces)
	logger.Infof("- End execution -  %s\n", methodName)
	return shim.Success(nil)
}

// Invoke is the entry point for all invocations
func (t *RemoverCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### RemoverCC Invoke ###########")
	function, _ := stub.GetFunctionAndParameters()

	switch function {
	case "DeleteState":
		logger.Info("########### Calling DeleteState ###########")
		return t.DeleteState(stub)
	}

	errorMsg := fmt.Sprintf("Could not find function named '%s' in RemoverCC.", function)
	logger.Errorf(errorMsg)
	return shim.Error(errorMsg)
}

// DeleteState deletes all data found under each one of the namespaces provided in the Init() method
func (t *RemoverCC) DeleteState(stub shim.ChaincodeStubInterface) pb.Response {
	methodName := "DeleteState()"
	logger.Infof("- Begin execution -  %s", methodName)

	totalRecordsDeleted := 0
	namespacesAsBytes, err := stub.GetState(Key)
	if err != nil {
		errorMsg := fmt.Sprintf("Error reading namespaces: %s", err.Error())
		logger.Error(errorMsg)
		return shim.Error(errorMsg)
	}
	namespaces := strings.Split(string(namespacesAsBytes), ",")

	logger.Infof("%s - Deleting data for namespaces: '%s'", methodName, strings.Join(namespaces, ","))

	// Delete records/state in each namespace
	for _, namespace := range namespaces {
		logger.Infof("%s - Deleting data for namespace '%s'.", methodName, namespace)
		recordsDeleted, err := t.DeleteRecordsByPartialKey(stub, namespace)
		if err != nil {
			return shim.Error(err.Error())
		}
		totalRecordsDeleted += recordsDeleted
		logger.Infof("%s - DeleteRecordsByPartialKey returned with total # of records deleted - %d for namespace %s", methodName, recordsDeleted, namespace)
	}

	logger.Infof("%s - Total number of records deleted accross all namespaces - %d", methodName, totalRecordsDeleted)
	totalDeleteCountInBytes := []byte(strconv.Itoa(totalRecordsDeleted))
	return shim.Success(totalDeleteCountInBytes)
}

// DeleteRecordsByPartialKey deletes records using a partial composite key (helper function used by DeleteState)
func (t *RemoverCC) DeleteRecordsByPartialKey(stub shim.ChaincodeStubInterface, namespace string) (int, error) {
	methodName := "DeleteRecordsByPartialKey()"
	logger.Infof("- Begin execution -  %s", methodName)
	var recordsDeletedCount = 0
	// Create composite partial key for namespace
	iterator, err := stub.GetStateByPartialCompositeKey(namespace, []string{})
	if err != nil {
		errorMsg := fmt.Sprintf("%s - Failed to get iterator for partial composite key: %s. Error: %s", methodName, namespace, err.Error())
		logger.Error(errorMsg)
		return recordsDeletedCount, err
	}

	// Once we are done with the iterator, we must close it
	defer iterator.Close()
	logger.Infof("%s - Starting to delete all records with namespace %s", methodName, namespace)

	for iterator.HasNext() {
		responseRange, err := iterator.Next()
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to get next record from iterator: %s", err.Error())
			logger.Error(errorMsg)
			return recordsDeletedCount, err
		}

		recordKey := responseRange.GetKey()
		logger.Infof("About to delete record with key %s", recordKey)
		err = stub.DelState(recordKey)
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to delete record '%d' with key %s: %s", recordsDeletedCount, recordKey, err.Error())
			logger.Error(errorMsg)
			return recordsDeletedCount, err
		}
		recordsDeletedCount++
		logger.Debugf("%s - Successfully deleted record '%d' with composite key: %s", methodName, recordsDeletedCount, recordKey)
	}

	logger.Infof("%s - Finished deleting all records found in %s", methodName, namespace)
	logger.Infof("- End execution -  %s", methodName)
	return recordsDeletedCount, nil
}
