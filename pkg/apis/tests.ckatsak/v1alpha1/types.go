package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Disttate is the distributed state.
type Disttate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec *DisttateSpec `json:"spec"`
}

// DisttateSpec is the Spec of the distributed state.
type DisttateSpec struct {
	Name     string `json:"name"`
	RingSize int    `json:"ringsize"`
}

func (dsp *DisttateSpec) GetCoolName() string {
	return "Cool " + dsp.Name
}

func (dsp *DisttateSpec) GetCoolRingSize() int {
	return dsp.RingSize % 42
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DisttateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Disttate `json:"items"`
}
