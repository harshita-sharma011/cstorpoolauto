/*
Copyright 2019 The MayaData Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

const (
	// AnnotationNamespace is the common namespace used across all
	// the annotations supported in this project
	AnnotationNamespace string = "dao.mayadata.io"

	// AnnKeyCStorClusterConfigUID is the annotation that refers to
	// CStorClusterConfig UID
	AnnKeyCStorClusterConfigUID string = AnnotationNamespace + "/cstorclusterconfig-uid"

	// AnnKeyCStorClusterPlanUID is the annotation that refers to
	// CStorClusterPlan UID
	AnnKeyCStorClusterPlanUID string = AnnotationNamespace + "/cstorclusterplan-uid"

	// AnnKeyCStorClusterStorageSetUID is the annotation that refers to
	// CStorClusterStorageSet UID
	AnnKeyCStorClusterStorageSetUID string = AnnotationNamespace + "/cstorclusterstorageset-uid"
)
