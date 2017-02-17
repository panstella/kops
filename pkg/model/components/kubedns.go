/*
Copyright 2016 The Kubernetes Authors.

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

package components

import (
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/upup/pkg/fi/loader"
)

// KubeDnsOptionsBuilder adds options for kube-dns
type KubeDnsOptionsBuilder struct {
	Context *OptionsContext
}

var _ loader.OptionsBuilder = &KubeDnsOptionsBuilder{}

func (b *KubeDnsOptionsBuilder) BuildOptions(o interface{}) error {
	clusterSpec := o.(*kops.ClusterSpec)

	if clusterSpec.KubeDNS == nil {
		clusterSpec.KubeDNS = &kops.KubeDNSConfig{}
	}

	clusterSpec.KubeDNS.Replicas = 2
	ip, err := WellKnownServiceIP(clusterSpec, 10)
	if err != nil {
		return err
	}
	clusterSpec.KubeDNS.ServerIP = ip.String()
	clusterSpec.KubeDNS.Domain = clusterSpec.ClusterDNSDomain
	// TODO: Once we start shipping more images, start using them
	clusterSpec.KubeDNS.Image = "gcr.io/google_containers/kubedns-amd64:1.3"

	return nil
}