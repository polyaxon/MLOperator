/*
Copyright 2018-2021 Polyaxon, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kinds

const (
	// KFAPIVersion api version
	KFAPIVersion = "kubeflow.org/v1"

	// MPIJobKind kind
	MPIJobKind = "MPIJob"

	// TFJobKind kind
	TFJobKind = "TFJob"

	// PytorchJobKind kind
	PytorchJobKind = "PyTorchJob"

	// MXJobKind kind
	MXJobKind = "MXJob"

	// XGBoostJobKind tfjob kind
	XGBoostJobKind = "XGBoostJob"

	// IstioAPIVersion istio networing api version
	IstioAPIVersion = "networking.istio.io/v1alpha3"

	// IstioVirtualServiceKind istio virtual service kind
	IstioVirtualServiceKind = "VirtualService"

	// SparkAPIVersion Spark operator api version
	SparkAPIVersion = "sparkoperator.k8s.io/v1beta2"

	// SparkApplicationKind Spark application kind
	SparkApplicationKind = "SparkApplication"
)
