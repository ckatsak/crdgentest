package v1alpha1

import (
	"encoding/base64"
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

	Bitset *BitSet `json:"bitset"`
}

func (dsp *DisttateSpec) GetCoolName() string {
	return "Cool " + dsp.Name
}

func (dsp *DisttateSpec) GetCoolRingSize() int {
	return dsp.RingSize % 42
}

// +k8s:deepcopy-gen=true

type BitSet struct {
	Bits []byte `json:"bits"`
}

var (
	_ json.Marshaler   = NewBitSet(0)
	_ json.Unmarshaler = NewBitSet(0)
)

func NewBitSet(cap int) *BitSet {
	return &BitSet{Bits: make([]byte, 0, cap)}
}

func (bs *BitSet) MarshalJSON() ([]byte, error) {
	// base64 encoding to string, and then json marshal as string.
	return json.Marshal(base64.URLEncoding.EncodeToString(bs.Bits))
}

func (bs *BitSet) UnmarshalJSON(in []byte) error {
	var s string
	// First unmarshall the input as a json string, to get the URL base64
	// encoded string...
	err := json.Unmarshal(in, &s)
	if err != nil {
		return err
	}
	// ...and then decode the URL base64 encoded string.
	bs.Bits = make([]byte, 0, len(s))
	bs.Bits, err = base64.URLEncoding.DecodeString(s)
	return err
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DisttateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Disttate `json:"items"`
}
