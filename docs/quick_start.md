# Quick Start

## Prerequisite
Since ONOS SD-RAN has multiple micro-services running on the Kubernetes platform, 
onos-pci can run on the Kubernetes along with other ONOS SD-RAN micro-services. In order to deploy onos-kpimon, a Helm chart is necessary, which is in the 
[sdran-helm-charts] repository. 
Note that this application should be running together with the other SD-RAN micro-services (e.g., Atomix, onos-e2t, onos-e2sub, and onos-cli). sd-ran umbrella chart can be used
to deploy all essential micro-services and onos-pci.




[sdran-helm-charts]: https://github.com/onosproject/sdran-helm-charts