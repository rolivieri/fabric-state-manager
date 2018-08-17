# fabric-reset-world-state

This repository contains a reusable Golang chaincode component for deleting the records found under a given list of namespaces,thus, effectively wiping out the world state. This chaincode component exposes the following methods:

 - DeleteState - Deletes all records found under the namespaces that are passed in to the `Init()` method.


## Compiling and running test cases

### Platform
It is stronly recommended to use **macOS** or a **Linux** flavor (such as Ubuntu) for compiling and testing the chaincode components.

### Prerequistes
1) Before you attempt to compile and run the tests cases for the chaincode components, please make sure you have the necessary [pre-requisites](https://hyperledger-fabric.readthedocs.io/en/release-1.1/prereqs.html) on your system:

* GO programming language (v1.9.x)

2) Once you have GO and Python installed on your system, you should download the Fabric v1.1.0 files. To do so, you can access the following URL from your browser: `https://github.com/hyperledger/fabric/archive/v1.1.0.tar.gz`. As an alternative, you can execute the following command from your command line:

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
1) Once you have cloned this repository, update your `GOPATH` environment variable so that it includes the absolute path to the root folder of this repository on your local file system. For instance, after cloning this repo to my local file system, the following is the absolute path to the root folder: `/Users/olivieri/git/fabric-reset-world-state`. Therefore, my `GOPATH` environment variable was updated as follows:

```
$ echo $GOPATH
/Users/olivieri/go
$ export GOPATH=$GOPATH:/Users/olivieri/git/fabric-reset-world-state
$ echo $GOPATH
/Users/olivieri/go:/Users/olivieri/git/fabric-reset-world-state
```

2) To compile the `DeleteStateCC` chaincode component, you can simply run the following command from the root folder of this repository:

```
$ pwd
/Users/olivieri/git/fabric-reset-world-state
$ go build DeleteStateCC
```

