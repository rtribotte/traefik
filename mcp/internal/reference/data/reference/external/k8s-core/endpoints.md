---
schema_version: 2
kind: kubernetes-core
name: Endpoints
id: k8s-core.endpoints
source: k8s-core
traefik_version: v3.7.0
extracted_from:
  - schemas/external/k8s-core/endpoints.json
summary: 'Endpoints is a collection of endpoints that implement the actual service. Example:'
fields:
  - name: subsets
    type: array
    items: object
    description: The set of all endpoints is the union of all subsets. Addresses are placed into subsets according to the IPs they share. A single address with multiple ports, some of which are ready and some of which are not (because they come from different containers) will result in the address being displayed in different subsets for the different ports. No address will appear in both Addresses and NotReadyAddresses in the same subset. Sets of addresses and ports that comprise a service.
representations:
  crd:
    apiVersion: v1
    kind: Endpoints
    spec_path: ""
---

# Endpoints

Endpoints is a collection of endpoints that implement the actual service. Example:

	 Name: "mysvc",
	 Subsets: [
	   {
	     Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],
	     Ports: [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]
	   },
	   {
	     Addresses: [{"ip": "10.10.3.3"}],
	     Ports: [{"name": "a", "port": 93}, {"name": "b", "port": 76}]
	   },
	]

Endpoints is a legacy API and does not contain information about all Service features. Use discoveryv1.EndpointSlice for complete information about Service endpoints.

Deprecated: This API is deprecated in v1.33+. Use discoveryv1.EndpointSlice.
