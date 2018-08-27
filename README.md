[![Build Status - Master](https://travis-ci.org/rolivieri/fabric-reset-world-state.svg?branch=master)](https://travis-ci.org/rolivieri/fabric-reset-world-state/builds)

# fabric-reset-world-state

This repository contains a reusable Golang chaincode component, `RemoverCC`, for deleting the records found under a list of namespaces. Thus, this chaincode component can be used to wipe out the world state. The `RemoverCC` chaincode component exposes the following methods:

 - DeleteState - Deletes all records found under the namespaces that are passed in to the `Initialize()` method.

Though more than likely you won't be resetting the world state in a production environment, doing so in a test or staging environment or as part of a PoC application is more than common.

## Compiling and running test cases

### Platform
It is stronly recommended to use **macOS** or a **Linux** flavor (such as Ubuntu) for compiling and testing the `RemoverCC` chaincode component.

### Prerequistes
1) Before you attempt to compile and run the tests cases for `RemoverCC`, please make sure you have the necessary [pre-requisites](https://hyperledger-fabric.readthedocs.io/en/release-1.1/prereqs.html) on your system:

* GO programming language (v1.9.x)

2) Once you have GO installed on your system, you should download the Fabric v1.1.0 files. To do so, you can access the following URL from your browser: `https://github.com/hyperledger/fabric/archive/v1.1.0.tar.gz`. As an alternative, you can execute the following command from your command line:

```
curl -O -L https://github.com/hyperledger/fabric/archive/v1.1.0.tar.gz
```

After downloading the Fabric files, you should untar the archive. You can do so by executing the following command:

```
tar -xvf v1.1.0.tar.gz
```

Untarring the above file results in the creation of a folder named `fabric-1.1.0`. You should now move the contents of the `fabric-1.1.0` folder into the following folder `${GOPATH}/src/github.com/hyperledger/fabric` (you may need to first create the `${GOPATH}/src/github.com/hyperledger/fabric` folder).

**Note**: Unfortunately, at the time of writing, the GO languge does not have yet an official dependency manager tool. Initially, [dep](https://github.com/golang/dep) was slated to become this tool but that no longer seems to be the case. Due to the lack of such a tool, the need for downloading the Fabric v1.1.0 archive, untarring it, and copying the files to the corresponding folder.

3) Finally, you'll also need to download the following assert library before proceeding with steps in the next section:

```
go get github.com/stretchr/testify/assert
```

### Steps
1) Once you have cloned this repository, navigate to its root folder and execute the commands shown below to compile and test the `RemoverCC` chaincode component:

```
$ pwd
/Users/olivieri/git/fabric-reset-world-state
Ricardos-MacBook-Pro:fabric-reset-world-state olivieri$ ls -la
total 88
drwxr-xr-x   8 olivieri  staff    256 Aug 27 13:00 .
drwxr-xr-x  39 olivieri  staff   1248 Aug 24 14:39 ..
drwxr-xr-x  16 olivieri  staff    512 Aug 27 12:58 .git
-rw-r--r--   1 olivieri  staff     25 Aug 27 09:47 .gitignore
-rw-r--r--   1 olivieri  staff    835 Aug 27 09:54 .travis.yml
-rw-r--r--   1 olivieri  staff  20192 Aug 27 13:00 README.md
-rw-r--r--   1 olivieri  staff   4292 Aug 27 12:58 main_test.go
-rw-r--r--   1 olivieri  staff   4699 Aug 27 12:58 removerCC.go
$ go build
$ go test
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Info -> INFO 001 ########### RemoverCC Init ###########
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Info -> INFO 002 ########### RemoverCC Initialize ###########
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 003 Initialize() - Namespaces provided to RemoverCC: [TestNamespace1 TestNamespace2]
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 004 - End execution -  Initialize()
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 005 - End execution -  Init()
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Info -> INFO 006 ########### RemoverCC Init ###########
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Info -> INFO 007 ########### RemoverCC Initialize ###########
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 008 Initialize() - Namespaces provided to RemoverCC: [TestNamespace1 TestNamespace2]
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 009 - End execution -  Initialize()
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 00a - End execution -  Init()
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
Inserting dummy record into namespace: TestNamespace1
Inserting dummy record into namespace: TestNamespace2
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Info -> INFO 00b ########### RemoverCC Invoke ###########
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Info -> INFO 00c ########### Calling DeleteState ###########
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 00d - Begin execution -  DeleteState()
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 00e DeleteState() - Deleting data for namespaces: 'TestNamespace1,TestNamespace2'
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 00f DeleteState() - Deleting data for namespace 'TestNamespace1'.
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 010 - Begin execution -  DeleteRecordsByPartialKey()
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 011 DeleteRecordsByPartialKey() - Starting to delete all records with namespace TestNamespace1
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 012 About to delete record with key TestNamespace10
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 013 About to delete record with key TestNamespace11
2018-08-27 13:01:09.265 EDT [RemoverCCLog] Infof -> INFO 014 About to delete record with key TestNamespace12
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 015 About to delete record with key TestNamespace13
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 016 About to delete record with key TestNamespace14
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 017 About to delete record with key TestNamespace15
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 018 About to delete record with key TestNamespace16
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 019 About to delete record with key TestNamespace17
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 01a About to delete record with key TestNamespace18
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 01b About to delete record with key TestNamespace19
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 01c DeleteRecordsByPartialKey() - Finished deleting all records found in TestNamespace1
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 01d - End execution -  DeleteRecordsByPartialKey()
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 01e DeleteState() - DeleteRecordsByPartialKey returned with total # of records deleted - 10 for namespace TestNamespace1
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 01f DeleteState() - Deleting data for namespace 'TestNamespace2'.
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 020 - Begin execution -  DeleteRecordsByPartialKey()
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 021 DeleteRecordsByPartialKey() - Starting to delete all records with namespace TestNamespace2
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 022 About to delete record with key TestNamespace20
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 023 About to delete record with key TestNamespace21
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 024 About to delete record with key TestNamespace22
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 025 About to delete record with key TestNamespace23
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 026 About to delete record with key TestNamespace24
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 027 About to delete record with key TestNamespace25
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 028 About to delete record with key TestNamespace26
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 029 About to delete record with key TestNamespace27
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 02a About to delete record with key TestNamespace28
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 02b About to delete record with key TestNamespace29
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 02c DeleteRecordsByPartialKey() - Finished deleting all records found in TestNamespace2
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 02d - End execution -  DeleteRecordsByPartialKey()
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 02e DeleteState() - DeleteRecordsByPartialKey returned with total # of records deleted - 10 for namespace TestNamespace2
2018-08-27 13:01:09.266 EDT [RemoverCCLog] Infof -> INFO 02f DeleteState() - Total number of records deleted accross all namespaces - 20
Summary: Expected number of deleted records = 20, actual number of deleted records from chain = 20 
 PASS
ok  	_/Users/olivieri/git/fabric-reset-world-state	0.029s
$  
```

## How to leverage this code from your chaincode component
In an ideal world, to use this code from your chaincode component, you'd simply need to:

* Deploy this code as a chaincode component to your channel.
* Instantiate this chaincode component by providing the namespaces whose data should deleted from the world state.
* Invoke this chaincode component from your chaincode component that was deployed to the same channel, or
* Invoke this chaincode component directly from your Fabric client.

Unfortunately, the above won't work. At the time of writing, Fabric takes into account the chaincode component when partitioning the data stored in a channel. This means that this chaincode component won't be able to read and/or delete any state that has been saved by another chaincode component on the same channel. Because of this limitation, you can follow these steps instead:

1. Download this repository as a dependency:

```
$ go get github.com/rolivieri/fabric-reset-world-state
```

2. Import the `statemanager` package into your chaincode component:

    ```
    package main

    import (

        ...

        sm "github.com/rolivieri/fabric-reset-world-state"

        ...
    )
    ```

3.  Use inheritance (or composition) to extend the capabiltiies of your code by referencing the `RemoverCC` structure (which resides in the `statemanager` package you just imported) in your chaincode component:
    
    ```
    type SampleChaincodeCC struct {      

        ...

        //Using inheritance in this sample
        sm.RemoverCC

        ...    
    }
    ```

4.  From the `Init()` method in your chaincode component, invoke the `Initialize(...)` method. The invocation to the `Initialize()` method from your chaincode should pass an array of strings that contains the namespaces whose data should be deleted from the world state.  Ex:

    ```
    // Init initializes chaincode
    func (t *SampleChaincodeCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	
        ...

        namespaces := []string{"namespace1", "namespace2", ... "namespaceN"}			
        t.Initialize(namespaces)

        ...

        return shim.Success(nil)
    }
    ```

5.  Add the `DeleteState(...)` method to the `Invoke()` method of your chaincode component.  Ex:

    ```
    // Invoke - Entry point for Invocations
    func (t *StakeholdersChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
        function, args := stub.GetFunctionAndParameters()

        switch function {
            case "Function1":
                return t.Function1(stub, args)
            case "Function2":
                return t.Function2(stub, args)
            
            ...

            case "DeleteState":		
                return t.DeleteState(stub)
            
            ...
        }

        ...

    }
    ```

6.  Whenever there is the need to reset the world state, your Fabric client application should call the `DeleteState()` method which will read the namespaces provided to the `Initialize()` method; the invocation of the `DeleteState()` method will result in the deletion of all the records found under those namespaces.

