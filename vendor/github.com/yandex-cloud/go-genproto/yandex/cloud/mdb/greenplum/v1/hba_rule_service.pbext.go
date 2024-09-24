// Code generated by protoc-gen-goext. DO NOT EDIT.

package greenplum

func (m *CreateHBARuleRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateHBARuleRequest) SetHbaRule(v *HBARule) {
	m.HbaRule = v
}

func (m *UpdateHBARuleRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateHBARuleRequest) SetHbaRule(v *HBARule) {
	m.HbaRule = v
}

func (m *DeleteHBARuleRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteHBARuleRequest) SetPriority(v int64) {
	m.Priority = v
}

func (m *ListHBARulesRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListHBARulesAtRevisionRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListHBARulesAtRevisionRequest) SetRevision(v int64) {
	m.Revision = v
}

func (m *ListHBARulesResponse) SetHbaRules(v []*HBARule) {
	m.HbaRules = v
}

func (m *BatchUpdateHBARulesRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *BatchUpdateHBARulesRequest) SetHbaRules(v []*HBARule) {
	m.HbaRules = v
}

func (m *HBARulesMetadata) SetClusterId(v string) {
	m.ClusterId = v
}
