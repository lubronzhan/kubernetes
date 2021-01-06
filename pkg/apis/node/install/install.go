/*
Copyright 2019 The Kubernetes Authors.

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

// Package install adds the node API group, making it available as
// an option to all of the API encoding/decoding machinery.
package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/apis/node"
	v1 "k8s.io/kubernetes/pkg/apis/node/v1"
	"k8s.io/kubernetes/pkg/apis/node/v1alpha1"
	"k8s.io/kubernetes/pkg/apis/node/v1beta1"
)

func init() {
	Install(legacyscheme.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(node.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(v1beta1.AddToScheme(scheme))
	utilruntime.Must(v1.AddToScheme(scheme))

	// TODO (SergeyKanzhelev): priority should change after 1.21. See https://github.com/kubernetes/kubernetes/pull/95718#discussion_r520969477
	// This is what controls the preferred serialization version. Add both v1beta1 and v1 here, and prefer v1beta1 over v1 until 1.21. See the comment on test/integration/etcd around serialized version.
	//
	// Details on why we can't advance the storage version for a release are at https://kubernetes.io/docs/reference/using-api/deprecation-policy/:
	//
	// > Rule #4b: The "preferred" API version and the "storage version" for a given group may not advance until after a release has been made that supports both the new version and the previous version
	utilruntime.Must(scheme.SetVersionPriority(v1beta1.SchemeGroupVersion, v1.SchemeGroupVersion))
}
