// Code generated by protoc-gen-goext. DO NOT EDIT.

package agent

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetTestRequest) SetTestId(v string) {
	m.TestId = v
}

func (m *UpdateTestRequest) SetTestId(v string) {
	m.TestId = v
}

func (m *UpdateTestRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateTestRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateTestRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateTestRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateTestRequest) SetFavorite(v bool) {
	m.Favorite = v
}

func (m *UpdateTestRequest) SetTargetVersion(v string) {
	m.TargetVersion = v
}

func (m *UpdateTestRequest) SetImbalancePoint(v int64) {
	m.ImbalancePoint = v
}

func (m *UpdateTestRequest) SetImbalanceTs(v int64) {
	m.ImbalanceTs = v
}

func (m *UpdateTestRequest) SetImbalanceComment(v string) {
	m.ImbalanceComment = v
}

func (m *UpdateTestMetadata) SetTestId(v string) {
	m.TestId = v
}