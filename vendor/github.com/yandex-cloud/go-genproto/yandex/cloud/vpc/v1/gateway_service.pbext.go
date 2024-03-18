// Code generated by protoc-gen-goext. DO NOT EDIT.

package vpc

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetGatewayRequest) SetGatewayId(v string) {
	m.GatewayId = v
}

func (m *ListGatewaysRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListGatewaysRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListGatewaysRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListGatewaysRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListGatewaysResponse) SetGateways(v []*Gateway) {
	m.Gateways = v
}

func (m *ListGatewaysResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListGatewayOperationsRequest) SetGatewayId(v string) {
	m.GatewayId = v
}

func (m *ListGatewayOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListGatewayOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListGatewayOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListGatewayOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

type CreateGatewayRequest_Gateway = isCreateGatewayRequest_Gateway

func (m *CreateGatewayRequest) SetGateway(v CreateGatewayRequest_Gateway) {
	m.Gateway = v
}

func (m *CreateGatewayRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateGatewayRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateGatewayRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateGatewayRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateGatewayRequest) SetSharedEgressGatewaySpec(v *SharedEgressGatewaySpec) {
	m.Gateway = &CreateGatewayRequest_SharedEgressGatewaySpec{
		SharedEgressGatewaySpec: v,
	}
}

func (m *CreateGatewayMetadata) SetGatewayId(v string) {
	m.GatewayId = v
}

type UpdateGatewayRequest_Gateway = isUpdateGatewayRequest_Gateway

func (m *UpdateGatewayRequest) SetGateway(v UpdateGatewayRequest_Gateway) {
	m.Gateway = v
}

func (m *UpdateGatewayRequest) SetGatewayId(v string) {
	m.GatewayId = v
}

func (m *UpdateGatewayRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateGatewayRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateGatewayRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateGatewayRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateGatewayRequest) SetSharedEgressGatewaySpec(v *SharedEgressGatewaySpec) {
	m.Gateway = &UpdateGatewayRequest_SharedEgressGatewaySpec{
		SharedEgressGatewaySpec: v,
	}
}

func (m *UpdateGatewayMetadata) SetGatewayId(v string) {
	m.GatewayId = v
}

func (m *DeleteGatewayRequest) SetGatewayId(v string) {
	m.GatewayId = v
}

func (m *DeleteGatewayMetadata) SetGatewayId(v string) {
	m.GatewayId = v
}

func (m *MoveGatewayRequest) SetGatewayId(v string) {
	m.GatewayId = v
}

func (m *MoveGatewayRequest) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *MoveGatewayMetadata) SetGatewayId(v string) {
	m.GatewayId = v
}