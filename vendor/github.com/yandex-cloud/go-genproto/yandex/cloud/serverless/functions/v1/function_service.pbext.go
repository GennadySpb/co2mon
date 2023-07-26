// Code generated by protoc-gen-goext. DO NOT EDIT.

package functions

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *GetFunctionRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *GetFunctionVersionRequest) SetFunctionVersionId(v string) {
	m.FunctionVersionId = v
}

func (m *GetFunctionVersionByTagRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *GetFunctionVersionByTagRequest) SetTag(v string) {
	m.Tag = v
}

func (m *ListFunctionsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListFunctionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFunctionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFunctionsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListFunctionsResponse) SetFunctions(v []*Function) {
	m.Functions = v
}

func (m *ListFunctionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateFunctionRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateFunctionRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateFunctionRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateFunctionRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateFunctionMetadata) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *UpdateFunctionRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *UpdateFunctionRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateFunctionRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateFunctionRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateFunctionRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateFunctionMetadata) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *DeleteFunctionRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *DeleteFunctionMetadata) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *ListRuntimesResponse) SetRuntimes(v []string) {
	m.Runtimes = v
}

type ListFunctionsVersionsRequest_Id = isListFunctionsVersionsRequest_Id

func (m *ListFunctionsVersionsRequest) SetId(v ListFunctionsVersionsRequest_Id) {
	m.Id = v
}

func (m *ListFunctionsVersionsRequest) SetFolderId(v string) {
	m.Id = &ListFunctionsVersionsRequest_FolderId{
		FolderId: v,
	}
}

func (m *ListFunctionsVersionsRequest) SetFunctionId(v string) {
	m.Id = &ListFunctionsVersionsRequest_FunctionId{
		FunctionId: v,
	}
}

func (m *ListFunctionsVersionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFunctionsVersionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFunctionsVersionsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListFunctionsVersionsResponse) SetVersions(v []*Version) {
	m.Versions = v
}

func (m *ListFunctionsVersionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListFunctionOperationsRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *ListFunctionOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFunctionOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFunctionOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListFunctionOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListFunctionOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

type CreateFunctionVersionRequest_PackageSource = isCreateFunctionVersionRequest_PackageSource

func (m *CreateFunctionVersionRequest) SetPackageSource(v CreateFunctionVersionRequest_PackageSource) {
	m.PackageSource = v
}

func (m *CreateFunctionVersionRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *CreateFunctionVersionRequest) SetRuntime(v string) {
	m.Runtime = v
}

func (m *CreateFunctionVersionRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateFunctionVersionRequest) SetEntrypoint(v string) {
	m.Entrypoint = v
}

func (m *CreateFunctionVersionRequest) SetResources(v *Resources) {
	m.Resources = v
}

func (m *CreateFunctionVersionRequest) SetExecutionTimeout(v *durationpb.Duration) {
	m.ExecutionTimeout = v
}

func (m *CreateFunctionVersionRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateFunctionVersionRequest) SetPackage(v *Package) {
	m.PackageSource = &CreateFunctionVersionRequest_Package{
		Package: v,
	}
}

func (m *CreateFunctionVersionRequest) SetContent(v []byte) {
	m.PackageSource = &CreateFunctionVersionRequest_Content{
		Content: v,
	}
}

func (m *CreateFunctionVersionRequest) SetVersionId(v string) {
	m.PackageSource = &CreateFunctionVersionRequest_VersionId{
		VersionId: v,
	}
}

func (m *CreateFunctionVersionRequest) SetEnvironment(v map[string]string) {
	m.Environment = v
}

func (m *CreateFunctionVersionRequest) SetTag(v []string) {
	m.Tag = v
}

func (m *CreateFunctionVersionRequest) SetConnectivity(v *Connectivity) {
	m.Connectivity = v
}

func (m *CreateFunctionVersionRequest) SetNamedServiceAccounts(v map[string]string) {
	m.NamedServiceAccounts = v
}

func (m *CreateFunctionVersionRequest) SetSecrets(v []*Secret) {
	m.Secrets = v
}

func (m *CreateFunctionVersionRequest) SetLogOptions(v *LogOptions) {
	m.LogOptions = v
}

func (m *CreateFunctionVersionRequest) SetStorageMounts(v []*StorageMount) {
	m.StorageMounts = v
}

func (m *CreateFunctionVersionMetadata) SetFunctionVersionId(v string) {
	m.FunctionVersionId = v
}

func (m *SetFunctionTagRequest) SetFunctionVersionId(v string) {
	m.FunctionVersionId = v
}

func (m *SetFunctionTagRequest) SetTag(v string) {
	m.Tag = v
}

func (m *RemoveFunctionTagRequest) SetFunctionVersionId(v string) {
	m.FunctionVersionId = v
}

func (m *RemoveFunctionTagRequest) SetTag(v string) {
	m.Tag = v
}

func (m *SetFunctionTagMetadata) SetFunctionVersionId(v string) {
	m.FunctionVersionId = v
}

func (m *RemoveFunctionTagMetadata) SetFunctionVersionId(v string) {
	m.FunctionVersionId = v
}

func (m *ListFunctionTagHistoryRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *ListFunctionTagHistoryRequest) SetTag(v string) {
	m.Tag = v
}

func (m *ListFunctionTagHistoryRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFunctionTagHistoryRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFunctionTagHistoryRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListFunctionTagHistoryResponse) SetFunctionTagHistoryRecord(v []*ListFunctionTagHistoryResponse_FunctionTagHistoryRecord) {
	m.FunctionTagHistoryRecord = v
}

func (m *ListFunctionTagHistoryResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListFunctionTagHistoryResponse_FunctionTagHistoryRecord) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *ListFunctionTagHistoryResponse_FunctionTagHistoryRecord) SetFunctionVersionId(v string) {
	m.FunctionVersionId = v
}

func (m *ListFunctionTagHistoryResponse_FunctionTagHistoryRecord) SetTag(v string) {
	m.Tag = v
}

func (m *ListFunctionTagHistoryResponse_FunctionTagHistoryRecord) SetEffectiveFrom(v *timestamppb.Timestamp) {
	m.EffectiveFrom = v
}

func (m *ListFunctionTagHistoryResponse_FunctionTagHistoryRecord) SetEffectiveTo(v *timestamppb.Timestamp) {
	m.EffectiveTo = v
}

func (m *ListScalingPoliciesRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *ListScalingPoliciesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListScalingPoliciesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListScalingPoliciesResponse) SetScalingPolicies(v []*ScalingPolicy) {
	m.ScalingPolicies = v
}

func (m *ListScalingPoliciesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *SetScalingPolicyRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *SetScalingPolicyRequest) SetTag(v string) {
	m.Tag = v
}

func (m *SetScalingPolicyRequest) SetProvisionedInstancesCount(v int64) {
	m.ProvisionedInstancesCount = v
}

func (m *SetScalingPolicyRequest) SetZoneInstancesLimit(v int64) {
	m.ZoneInstancesLimit = v
}

func (m *SetScalingPolicyRequest) SetZoneRequestsLimit(v int64) {
	m.ZoneRequestsLimit = v
}

func (m *SetScalingPolicyMetadata) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *RemoveScalingPolicyRequest) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *RemoveScalingPolicyRequest) SetTag(v string) {
	m.Tag = v
}

func (m *RemoveScalingPolicyMetadata) SetFunctionId(v string) {
	m.FunctionId = v
}
