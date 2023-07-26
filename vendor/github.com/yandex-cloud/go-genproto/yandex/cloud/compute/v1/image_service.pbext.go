// Code generated by protoc-gen-goext. DO NOT EDIT.

package compute

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetImageRequest) SetImageId(v string) {
	m.ImageId = v
}

func (m *GetImageLatestByFamilyRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *GetImageLatestByFamilyRequest) SetFamily(v string) {
	m.Family = v
}

func (m *ListImagesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListImagesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListImagesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListImagesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListImagesRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListImagesResponse) SetImages(v []*Image) {
	m.Images = v
}

func (m *ListImagesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

type CreateImageRequest_Source = isCreateImageRequest_Source

func (m *CreateImageRequest) SetSource(v CreateImageRequest_Source) {
	m.Source = v
}

func (m *CreateImageRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateImageRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateImageRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateImageRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateImageRequest) SetFamily(v string) {
	m.Family = v
}

func (m *CreateImageRequest) SetMinDiskSize(v int64) {
	m.MinDiskSize = v
}

func (m *CreateImageRequest) SetProductIds(v []string) {
	m.ProductIds = v
}

func (m *CreateImageRequest) SetImageId(v string) {
	m.Source = &CreateImageRequest_ImageId{
		ImageId: v,
	}
}

func (m *CreateImageRequest) SetDiskId(v string) {
	m.Source = &CreateImageRequest_DiskId{
		DiskId: v,
	}
}

func (m *CreateImageRequest) SetSnapshotId(v string) {
	m.Source = &CreateImageRequest_SnapshotId{
		SnapshotId: v,
	}
}

func (m *CreateImageRequest) SetUri(v string) {
	m.Source = &CreateImageRequest_Uri{
		Uri: v,
	}
}

func (m *CreateImageRequest) SetOs(v *Os) {
	m.Os = v
}

func (m *CreateImageRequest) SetPooled(v bool) {
	m.Pooled = v
}

func (m *CreateImageMetadata) SetImageId(v string) {
	m.ImageId = v
}

func (m *UpdateImageRequest) SetImageId(v string) {
	m.ImageId = v
}

func (m *UpdateImageRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateImageRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateImageRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateImageRequest) SetMinDiskSize(v int64) {
	m.MinDiskSize = v
}

func (m *UpdateImageRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateImageMetadata) SetImageId(v string) {
	m.ImageId = v
}

func (m *DeleteImageRequest) SetImageId(v string) {
	m.ImageId = v
}

func (m *DeleteImageMetadata) SetImageId(v string) {
	m.ImageId = v
}

func (m *ListImageOperationsRequest) SetImageId(v string) {
	m.ImageId = v
}

func (m *ListImageOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListImageOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListImageOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListImageOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
