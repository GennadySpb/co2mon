// Code generated by protoc-gen-goext. DO NOT EDIT.

package vision

func (m *ClassAnnotation) SetProperties(v []*Property) {
	m.Properties = v
}

func (m *Property) SetName(v string) {
	m.Name = v
}

func (m *Property) SetProbability(v float64) {
	m.Probability = v
}
