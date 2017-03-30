/*
Copyright 2016 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Some of the code below came from https://github.com/coreos/etcd-operator
which also has the apache 2.0 license.
*/
package cluster

type Spec struct {
	// The namespace where the the rook resources will all be created.
	Namespace string `json:"namespace"`

	// Version is the expected version of the rook container to run in the cluster.
	// The operator will eventually make the rook cluster version
	// equal to the expected version.
	Version string `json:"version"`

	// Paused is to pause the control of the operator for the rook cluster.
	Paused bool `json:"paused,omitempty"`

	// Whether to consume all the storage devices found on a machine
	UseAllDevices bool `json:"useAllDevices"`

	// A regular expression to allow more fine-grained selection of devices on nodes across the cluster
	DeviceFilter string `json:"deviceFilter"`

	// The path on the host where config and data can be persisted.
	DataDirHostPath string `json:"dataDirHostPath"`
}

type PoolSpec struct {
	// The name of the pool. Defined in the spec since the object metadata name must be unique across all namespaces,
	// while you could have the same pool name created in multiple instances of rook.
	Name string `json:"name"`

	// The namespace where the pool will be created (required). A Rook cluster must be running in this namespace.
	Namespace string `json:"namespace"`

	// The replication settings
	Replication ReplicationSpec `json:"replication"`

	// The erasure code setteings
	ErasureCoding ErasureCodeSpec `json:"erasureCode"`
}

type ReplicationSpec struct {
	// Number of copies per object in a replicated storage pool, including the object itself (required for replicated pool type)
	Count uint `json:"count"`
}

type ErasureCodeSpec struct {
	// Number of coding chunks per object in an erasure coded storage pool (required for erasure-coded pool type)
	CodingChunks uint `json:"codingChunks"`

	// Number of data chunks per object in an erasure coded storage pool (required for erasure-coded pool type)
	DataChunks uint `json:"dataChunks"`
}
