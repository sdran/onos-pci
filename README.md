Welcome to the open source release version 1.1.1 of ONF's SD-RAN project. Please note that this is a read only release of the source code. We will not be accepting pull requests in these repos, and the source code that is contained here cannot be used to build functional binaries because it refers back to a number of private repositories. But executables - docker containers referencing the released code - are available on Docker Hub and referenced [here].  SD-RAN is currently a member-only project, and ONF membership is required to access the most current release, the master branch, and to do development with SD-RAN.


# onos-pci
PCI xAPP for ONOS SD-RAN (µONOS Architecture)

## Overview 
The onos-pci is an xApp running over ONOS SD-RAN and as of now it supports the following features:

* Provides capability to subscribe to RC-PRE service model and receives indication messages from RAN simulator

* Provides capability to send control requests to change PCI values in RAN simulator

* Supports listing of PCI resources such metrics, neighbors, PCI, and PCI conflicts of cell(s) using CLI that is integrated with [onos-cli] 

* Detects PCI conflicts and resolves them based on an algorithm using cell neighbors information


See [README.md](docs/README.md) for details of running the onos-pci application.


[onos-cli]: https://github.com/onosproject/onos-cli
[here]: https://wiki.opennetworking.org/display/COM/SD-RAN+1.1+Release
