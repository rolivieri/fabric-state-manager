[![Build Status - Master](https://travis-ci.org/rolivieri/fabric-state-manager.svg?branch=master)](https://travis-ci.org/rolivieri/fabric-state-manager/builds)

# fabric-state-manager

This repository contains a reusable Golang chaincode component, `RemoverCC`, for deleting the records found under a list of namespaces. Thus, this chaincode component can be used to wipe out the world state. The `RemoverCC` chaincode component exposes the following methods:

 - DeleteState - Deletes all records found under the namespaces that are passed in to the `Initialize()` method.

Though more than likely you won't be resetting the world state in a production environment, doing so in a test or staging environment or as part of a PoC application is more than common.

## Compiling and running test cases

### Platform
It is strongly recommended to use **macOS** or a **Linux** flavor (such as Ubuntu) for compiling and testing the `RemoverCC` chaincode component.

### Prerequisites
1) Before you attempt to compile and run the tests cases for `RemoverCC`, please make sure you have the necessary [pre-requisites](https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html) on your system:

* GO programming language (v1.11.x)

2) Once you have GO installed on your system, you should download the Fabric v1.4.2 files. To do so, you can access the following URL from your browser: `https://github.com/hyperledger/fabric/archive/v1.4.2.tar.gz`. As an alternative, you can execute the following command from your command line:

```
curl -O -L https://github.com/hyperledger/fabric/archive/v1.4.2.tar.gz
```

After downloading the Fabric files, you should untar the archive. You can do so by executing the following command:

```
tar -xvf v1.4.2.tar.gz
```

Untarring the above file results in the creation of a folder named `fabric-1.4.2`. You should now move the contents of the `fabric-1.4.2` folder into the following folder `${GOPATH}/src/github.com/hyperledger/fabric` (you may need to first create the `${GOPATH}/src/github.com/hyperledger/fabric` folder).

**Note**: Unfortunately, at the time of writing, the GO language does not have yet an official dependency manager tool. Initially, [dep](https://github.com/golang/dep) was slated to become this tool but that no longer seems to be the case. Due to the lack of such a tool, the need for downloading the Fabric v1.4.2 archive, untarring it, and copying the files to the corresponding folder.

3) Finally, you'll also need to download the following assert library before proceeding with steps in the next section:

```
go get github.com/stretchr/testify/assert
```

### Steps
1) Once you have cloned this repository into your GO workspace, navigate to its root folder and execute the commands shown below to compile and test the `RemoverCC` chaincode component:

```
$ pwd
/Users/olivieri/go/fabric-state-manager
$ ls -la
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
ok  	_/Users/olivieri/go/fabric-state-manager	0.029s
$  
```

## How to leverage this code from your chaincode component
### Development
Please remember that Hyperledger Fabric takes into account chaincode components when partitioning the data stored in a channel. In other words, a chaincode component won't be able to read and/or delete any state that has been saved by another chaincode component on the same channel. Because of this, you cannot just deploy this code as a chaincode component and expect it to have access to the data written by another chaincode. Therefore, to leverage the code in this repository, you can follow the following steps:

1. Download this repository as a dependency:

```
$ go get github.com/rolivieri/fabric-state-manager
```

2. Import the `fabric-state-manager` package into your chaincode component:

    ```
    package main

    import (

        ...

        sm "github.com/rolivieri/fabric-state-manager"

        ...
    )
    ```

3.  Use inheritance (or composition) to extend the capabilities of your code by referencing the `RemoverCC` structure (which resides in the `fabric-state-manager` package you just imported) in your chaincode component:
    
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

### Deployment
As mentioned before, hopefully there will soon be a dependency manager tool for GO. Also, once this is available, hopefully the Hyperledger Fabric platform makes use of this tool so chaincode developers can simply specify in some form their chaincode dependencies so they can be pulled down on behalf of the developer (think of `package.json` for Node.js apps or  `Package.swift` for Swift apps). In the mean time, to deploy a chaincode component that leverages the code in this repository, you will need to perform a few manual steps.

1. Install the `govendor` tool:

```
go get -u github.com/kardianos/govendor
```

For further details, please see the govendor [documentation](https://github.com/kardianos/govendor).

2. From the root folder of your chaincode component, execute the following commands:

```
govendor init
govendor add github.com/rolivieri/fabric-state-manager

```

Please note the above assumes you have the `govendor` tool in your `PATH`.

The execution of the above commands should result in the creation of the `vendor` folder and in it you should find the `fabric-state-manager` dependency:

```
$ pwd
/Users/olivieri/go/src/my-chaincode/vendor
$ ls -la
total 8
drwxr-xr-x   4 olivieri  staff  128 Aug 30 13:21 .
drwxr-xr-x  10 olivieri  staff  320 Aug 30 13:19 ..
drwxr-xr-x   3 olivieri  staff   96 Aug 30 13:21 github.com
-rw-r--r--   1 olivieri  staff  313 Aug 30 13:21 vendor.json
$ find .
.
./vendor.json
./github.com
./github.com/rolivieri
./github.com/rolivieri/fabric-state-manager
./github.com/rolivieri/fabric-state-manager/README.md
./github.com/rolivieri/fabric-state-manager/removerCC.go
$ 
```

By following the steps above, you have effectively included in your chaincode component the `fabric-state-manager` as a dependency that will be available when your chaincode is deployed and instantiated in a Hyperledger Fabric network. For further details, please see [Managing external dependencies for chaincode written in Go](https://hyperledger-fabric.readthedocs.io/en/release-1.4/chaincode4ade.html#managing-external-dependencies-for-chaincode-written-in-go).