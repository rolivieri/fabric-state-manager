# fabric-reset-world-state

This repository contains a reusable Golang chaincode component, `DeleteStateCC`, for deleting the records found under a given list of namespaces. Thus, this chaincode component can be used to wipe out the world state. The `DeleteStateCC` chaincode component exposes the following methods:

 - DeleteState - Deletes all records found under the namespaces that are passed in to the `Init()` method.

## Compiling and running test cases

### Platform
It is stronly recommended to use **macOS** or a **Linux** flavor (such as Ubuntu) for compiling and testing the `DeleteStateCC` chaincode component.

### Prerequistes
1) Before you attempt to compile and run the tests cases for `DeleteStateCC`, please make sure you have the necessary [pre-requisites](https://hyperledger-fabric.readthedocs.io/en/release-1.1/prereqs.html) on your system:

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

**Note**: Unfortunately, at the time of writing, the GO languge does not have yet an official dependency manager tool. Hopefully, [dep](https://github.com/golang/dep) will soon become this tool. However, last time we tested `dep` for downloading Fabric, it did not quite yield the desired outcome. Hence, the need for the steps described here for downloading the Fabric v1.1.0 archive, untarring it, and copying the files to the corresponding folder.

3) Finally, you'll also need to download the following assert library before proceeding with steps in the next section:

```
go get github.com/stretchr/testify/assert
```

### Steps
1) Once you have cloned this repository, you should update your `GOPATH` environment variable so that it includes the absolute path to the root folder of this repository on your local file system. For instance, after cloning this repo to my local file system, the following is the absolute path to the root folder: `/Users/olivieri/git/fabric-reset-world-state`. Therefore, my `GOPATH` environment variable was updated as follows:

```
$ echo $GOPATH
/Users/olivieri/go
$ export GOPATH=$GOPATH:/Users/olivieri/git/fabric-reset-world-state
$ echo $GOPATH
/Users/olivieri/go:/Users/olivieri/git/fabric-reset-world-state
```

2) To compile and test the `deleteStateCC` chaincode component, you can simply run the following commands:

```
$ pwd
/Users/olivieri/git/fabric-reset-world-state/src/deleteStateCC
$ go build
$ go test
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Info -> INFO 001 ########### DeleteStateCC Init ###########
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 002 Namespsaces provided to DeleteStateCC: [TestNamespace1 TestNamespace2]
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 003 - End execution -  Init()
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Info -> INFO 004 ########### DeleteStateCC Init ###########
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 005 Namespsaces provided to DeleteStateCC: [TestNamespace1 TestNamespace2]
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 006 - End execution -  Init()
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
Inserting dummy reconrd into namespace: TestNamespace1
Inserting dummy reconrd into namespace: TestNamespace2
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Info -> INFO 007 ########### DeleteStateCC Invoke ###########
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Info -> INFO 008 ########### Calling DeleteState ###########
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 009 - Begin execution -  DeleteState()
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 00a DeleteState() - Deleting data for namespaces: 'TestNamespace1,TestNamespace2'
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 00b DeleteState() - Deleting data for namespace 'TestNamespace1'.
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 00c - Begin execution -  DeleteRecordsByPartialKey()
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 00d Starting to delete all records with namespace TestNamespace1
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 00e About to delete record with key TestNamespace10
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 00f About to delete record with key TestNamespace11
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 010 About to delete record with key TestNamespace12
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 011 About to delete record with key TestNamespace13
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 012 About to delete record with key TestNamespace14
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 013 About to delete record with key TestNamespace15
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 014 About to delete record with key TestNamespace16
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 015 About to delete record with key TestNamespace17
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 016 About to delete record with key TestNamespace18
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 017 About to delete record with key TestNamespace19
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 018 Finished deleting all records found in TestNamespace1
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 019 - End execution -  DeleteRecordsByPartialKey()
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 01a - DeleteRecordsByPartialKey returned with total # of records deleted - 10 for namespace TestNamespace1
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 01b DeleteState() - Deleting data for namespace 'TestNamespace2'.
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 01c - Begin execution -  DeleteRecordsByPartialKey()
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 01d Starting to delete all records with namespace TestNamespace2
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 01e About to delete record with key TestNamespace20
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 01f About to delete record with key TestNamespace21
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 020 About to delete record with key TestNamespace22
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 021 About to delete record with key TestNamespace23
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 022 About to delete record with key TestNamespace24
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 023 About to delete record with key TestNamespace25
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 024 About to delete record with key TestNamespace26
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 025 About to delete record with key TestNamespace27
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 026 About to delete record with key TestNamespace28
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 027 About to delete record with key TestNamespace29
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 028 Finished deleting all records found in TestNamespace2
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 029 - End execution -  DeleteRecordsByPartialKey()
2018-08-17 11:34:37.313 EDT [DeleteStateCCLog] Infof -> INFO 02a - DeleteRecordsByPartialKey returned with total # of records deleted - 10 for namespace TestNamespace2
2018-08-17 11:34:37.314 EDT [DeleteStateCCLog] Infof -> INFO 02b - Total number of records deleted accross all namespaces - 20
Summary: Expected number of deleted records = 20, actual number of deleted records from chain = 20 
 PASS
ok  	deleteStateCC	0.029s
$ 
```

## How to leverage this code from your chaincode component
Ideally to use this code, you'd simply need to:

* Deploy this code as a chaincode component to your channel.
* Instantiate the chaincode by providing the namespaces.
* Invoking this chaincode component from your chaincode component that is also deployed to the same channel.

Unfortunately, the above won't work. Fabric takes into account the chaincode component when partitioning the data. This means that this chaincode component won't be able to read and/or delete any state that has been saved by another chaincode component on the same channel. Because of this limitation, you'd need to follow these steps instead:

TODO: Complete documentation

