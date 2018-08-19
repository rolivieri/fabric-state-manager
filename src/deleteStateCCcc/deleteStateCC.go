package ws

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("DeleteStateCCLog")

// DeleteStateCC chaincode structure
type DeleteStateCC struct {
	// Namespaces array variable
	Namespaces []string
}

// Init initializes chaincode
func (t *DeleteStateCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### DeleteStateCC Init ###########")
	methodName := "Init()"
	_, args := stub.GetFunctionAndParameters()
	for _, namespace := range args {
		t.Namespaces = append(t.Namespaces, namespace)
	}

	if len(t.Namespaces) == 0 {
		warningMsg := fmt.Sprintf("%s - No namespaces were provided to DeleteStateCC.", methodName)
		logger.Warning(warningMsg)
	}

	logger.Infof("Namespaces provided to DeleteStateCC: %v", t.Namespaces)
	logger.Infof("- End execution -  %s\n", methodName)
	return shim.Success(nil)
}

// Invoke is the entry point for all invocations
func (t *DeleteStateCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, _ := stub.GetFunctionAndParameters()
	logger.Info("########### DeleteStateCC Invoke ###########")

	switch function {
	case "DeleteState":
		logger.Info("########### Calling DeleteState ###########")
		return t.DeleteState(stub)
	}

	errorMsg := fmt.Sprintf("Could not find function named '%s' in DeleteStateCC.", function)
	logger.Errorf(errorMsg)
	return shim.Error(errorMsg)
}

// DeleteState deletes all data found under each one of the namespaces provided in the Init() method
func (t *DeleteStateCC) DeleteState(stub shim.ChaincodeStubInterface) pb.Response {
	methodName := "DeleteState()"
	logger.Infof("- Begin execution -  %s", methodName)

	totalRecordsDeleted := 0
	logger.Infof("%s - Deleting data for namespaces: '%s'", methodName, strings.Join(t.Namespaces, ","))

	// Delete records/state in each namespace
	for _, namespace := range t.Namespaces {
		logger.Infof("%s - Deleting data for namespace '%s'.", methodName, namespace)
		recordsDeleted, err := t.DeleteRecordsByPartialKey(stub, namespace)
		if err != nil {
			return shim.Error(err.Error())
		}
		totalRecordsDeleted += recordsDeleted
		logger.Infof("- DeleteRecordsByPartialKey returned with total # of records deleted - %d for namespace %s", recordsDeleted, namespace)
	}

	logger.Infof("- Total number of records deleted accross all namespaces - %d", totalRecordsDeleted)
	totalDeleteCountInBytes := []byte(strconv.Itoa(totalRecordsDeleted))
	return shim.Success(totalDeleteCountInBytes)
}

// DeleteRecordsByPartialKey deletes records using a partial composite key (helper function used by DeleteState)
func (t *DeleteStateCC) DeleteRecordsByPartialKey(stub shim.ChaincodeStubInterface, namespace string) (int, error) {
	methodName := "DeleteRecordsByPartialKey()"
	logger.Infof("- Begin execution -  %s", methodName)
	var recordsDeletedCount = 0
	// Create composite partial key for namespace
	iterator, err := stub.GetStateByPartialCompositeKey(namespace, []string{})
	if err != nil {
		errorMsg := "Failed to get iterator for partial composite key:" + namespace + ". Error: " + err.Error()
		logger.Error(errorMsg)
		return recordsDeletedCount, err
	}

	// Once we are done with the iterator, we must close it
	defer iterator.Close()
	logger.Infof("Starting to delete all records with namespace %s", namespace)

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
		logger.Debugf("Successfully deleted record '%d' with composite key: %s", recordsDeletedCount, recordKey)
	}

	logger.Infof("Finished deleting all records found in %s", namespace)
	logger.Infof("- End execution -  %s", methodName)
	return recordsDeletedCount, nil
}
