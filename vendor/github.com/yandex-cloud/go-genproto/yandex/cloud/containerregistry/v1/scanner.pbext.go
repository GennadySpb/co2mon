// Code generated by protoc-gen-goext. DO NOT EDIT.

package containerregistry

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *ScanResult) SetId(v string) {
	m.Id = v
}

func (m *ScanResult) SetImageId(v string) {
	m.ImageId = v
}

func (m *ScanResult) SetScannedAt(v *timestamppb.Timestamp) {
	m.ScannedAt = v
}

func (m *ScanResult) SetStatus(v ScanResult_Status) {
	m.Status = v
}

func (m *ScanResult) SetVulnerabilities(v *VulnerabilityStats) {
	m.Vulnerabilities = v
}

func (m *VulnerabilityStats) SetCritical(v int64) {
	m.Critical = v
}

func (m *VulnerabilityStats) SetHigh(v int64) {
	m.High = v
}

func (m *VulnerabilityStats) SetMedium(v int64) {
	m.Medium = v
}

func (m *VulnerabilityStats) SetLow(v int64) {
	m.Low = v
}

func (m *VulnerabilityStats) SetNegligible(v int64) {
	m.Negligible = v
}

func (m *VulnerabilityStats) SetUndefined(v int64) {
	m.Undefined = v
}

type Vulnerability_Vulnerability = isVulnerability_Vulnerability

func (m *Vulnerability) SetVulnerability(v Vulnerability_Vulnerability) {
	m.Vulnerability = v
}

func (m *Vulnerability) SetSeverity(v Vulnerability_Severity) {
	m.Severity = v
}

func (m *Vulnerability) SetPackage(v *PackageVulnerability) {
	m.Vulnerability = &Vulnerability_Package{
		Package: v,
	}
}

func (m *PackageVulnerability) SetName(v string) {
	m.Name = v
}

func (m *PackageVulnerability) SetLink(v string) {
	m.Link = v
}

func (m *PackageVulnerability) SetPackage(v string) {
	m.Package = v
}

func (m *PackageVulnerability) SetSource(v string) {
	m.Source = v
}

func (m *PackageVulnerability) SetVersion(v string) {
	m.Version = v
}

func (m *PackageVulnerability) SetFixedBy(v string) {
	m.FixedBy = v
}

func (m *PackageVulnerability) SetOrigin(v string) {
	m.Origin = v
}

func (m *PackageVulnerability) SetType(v string) {
	m.Type = v
}
