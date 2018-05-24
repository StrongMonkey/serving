/*
Copyright 2018 Google LLC

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

package revision

import (
	"github.com/elafros/elafros/pkg/apis/ela/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// fluentdSidecarPreOutputConfig defines source and filter configurations for
// files under /var/log.
const fluentdSidecarPreOutputConfig = `
<source>
	@type tail
	path /var/log/revisions/**/*.*
	pos_file /var/log/varlog.log.pos
	tag *
	<parse>
		@type multi_format
		<pattern>
			format json
			time_key fluentd-time # fluentd-time is reserved for structured logs
			time_format %Y-%m-%dT%H:%M:%S.%NZ
		</pattern>
		<pattern>
			format none
			message_key log
		</pattern>
	</parse>
	read_from_head true
</source>

<filter var.log.**>
	@type record_transformer
	<record>
		kubernetes.container_name "#{ENV['ELA_CONTAINER_NAME']}"
		kubernetes.labels.elafros_dev/configuration "#{ENV['ELA_CONFIGURATION']}"
		kubernetes.labels.elafros_dev/revision "#{ENV['ELA_REVISION']}"
		kubernetes.namespace_name "#{ENV['ELA_NAMESPACE']}"
		kubernetes.pod_name "#{ENV['ELA_POD_NAME']}"
		stream varlog
	</record>
</filter>

`

const fluentdConfigMapName = "fluentd-varlog-config"

// MakeFluentdConfigMap creates a ConfigMap that gets mounted for fluentd
// container on the pod.
func MakeFluentdConfigMap(
	rev *v1alpha1.Revision, namespace string, fluentdSidecarOutputConfig string) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      fluentdConfigMapName,
			Namespace: namespace,
			Labels:    MakeElaResourceLabels(rev),
		},
		Data: map[string]string{
			"varlog.conf": makeFullFluentdConfig(fluentdSidecarOutputConfig),
		},
	}
}

func makeFullFluentdConfig(fluentdSidecarOutputConfig string) string {
	return fluentdSidecarPreOutputConfig + fluentdSidecarOutputConfig
}