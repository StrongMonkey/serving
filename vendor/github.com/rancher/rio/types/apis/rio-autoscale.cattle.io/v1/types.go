package v1

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceScaleRecommendation struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceScaleRecommendationSpec   `json:"spec,omitempty"`
	Status ServiceScaleRecommendationStatus `json:"status,omitempty"`
}

type ServiceScaleRecommendationSpec struct {
	ServiceNameToRead   string `json:"serviceNameToRead,omitempty"`
	ServiceNameToChange string `json:"serviceNameToChange,omitempty"`
	ZeroScaleService    string `json:"zeroScaleService,omitempty"`
	DeploymentName      string `json:"deploymentName,omitempty"`
	MinScale            int32  `json:"minScale,omitempty"`
	MaxScale            int32  `json:"maxScale,omitempty"`
	Concurrency         int    `json:"concurrency,omitempty"`
	PrometheusURL       string `json:"prometheusURL,omitempty"`
}

type ServiceScaleRecommendationStatus struct {
	DesiredScale *int32 `json:"desiredScale,omitempty"`
}
