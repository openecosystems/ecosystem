// Code generated by protoc-gen-platform go/entity-unspecified. DO NOT EDIT.
// source: platform/communication/v1beta/preference_center.proto

package communicationv1betapb

import (
	"context"
	"encoding/json"

	"errors"
	"github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
	"google.golang.org/protobuf/types/known/anypb"

	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type PreferenceCenterSpecEntity struct {
	PreferenceCenter *PreferenceCenter
}

func NewPreferenceCenterSpecEntity(specContext *specv2pb.SpecContext) (*PreferenceCenterSpecEntity, error) {

	return &PreferenceCenterSpecEntity{
		PreferenceCenter: &PreferenceCenter{},
	}, nil

}

func NewPreferenceCenterSpecEntityFromSpec(ctx context.Context, s *specv2pb.Spec) (*PreferenceCenterSpecEntity, error) {
	data := &PreferenceCenter{}
	err := sdkv2alphalib.GetDataFromSpec[*PreferenceCenter](ctx, s, data)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err)
	}

	return &PreferenceCenterSpecEntity{
		PreferenceCenter: data,
	}, nil
}

func (entity *PreferenceCenterSpecEntity) ToProto() (*PreferenceCenter, error) {

	return entity.PreferenceCenter, nil

}

func (entity *PreferenceCenterSpecEntity) ToEvent() (*string, error) {

	bytes, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	event := string(bytes)

	return &event, nil

}

func (entity *PreferenceCenterSpecEntity) FromEvent(event *string) (*PreferenceCenterSpecEntity, error) {

	bytes := []byte(*event)
	err := json.Unmarshal(bytes, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil

}

func (entity *PreferenceCenterSpecEntity) MarshalEntity() (*anypb.Any, error) {

	d, err := anypb.New(entity.PreferenceCenter)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to marshall entity"), err)
	}

	return d, nil

}

func (entity *PreferenceCenterSpecEntity) MarshalProto() (*anypb.Any, error) {

	proto, err := entity.ToProto()
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to convert entity to proto"), err)
	}

	d, err := anypb.New(proto)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to marshall proto"), err)
	}

	return d, nil

}

func (entity *PreferenceCenterSpecEntity) TypeName() string {
	return "preferenceCenter"
}

func (entity *PreferenceCenterSpecEntity) CommandTopic() string {
	return CommandDataPreferenceCenterTopic
}

func (entity *PreferenceCenterSpecEntity) EventTopic() string {
	return EventDataPreferenceCenterTopic
}

func (entity *PreferenceCenterSpecEntity) RoutineTopic() string {
	return RoutineDataPreferenceCenterTopic
}

func (entity *PreferenceCenterSpecEntity) TopicWildcard() string {
	return PreferenceCenterTypeNameEventPrefix + ">"
}

func (entity *PreferenceCenterSpecEntity) SystemName() string {
	return "communication"
}

func (entity *PreferenceCenterSpecEntity) internal() {

	var _ timestamppb.Timestamp
	//created_at

	var _ timestamppb.Timestamp
	//updated_at

}
