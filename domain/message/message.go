package message

import (
	"context"

	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/pkg/errors"
	"github.com/rerost/giro/domain/grpcreflectiface"
	"github.com/rerost/giro/domain/messagename"
)

type JSON []byte
type Binary []byte

type MessageService interface {
	EmptyJSON(ctx context.Context, messageName messagename.MessageName) (JSON, error)
	ToJSON(ctx context.Context, messageName messagename.MessageName, binary Binary) (JSON, error)
	ToBinary(ctx context.Context, messageName messagename.MessageName, json JSON) (Binary, error)
	// NOTE: For internal.
	ToDynamicMessage(ctx context.Context, messageName messagename.MessageName, json JSON) (*dynamic.Message, error)
	DynamicMessageToJSON(ctx context.Context, dm *dynamic.Message) (JSON, error)
}

type messageServiceImpl struct {
	grpcreflectClient grpcreflectiface.Client
	jsonMarshaler     *jsonpb.Marshaler
}

func NewMessageService(client grpcreflectiface.Client) MessageService {
	return &messageServiceImpl{
		grpcreflectClient: client,
		jsonMarshaler: &jsonpb.Marshaler{
			EmitDefaults: true,
		},
	}
}

func (ms messageServiceImpl) EmptyJSON(ctx context.Context, messageName messagename.MessageName) (JSON, error) {
	md, err := ms.grpcreflectClient.ResolveMessage(string(messageName))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dMessage := dynamic.NewMessageFactoryWithDefaults().NewDynamicMessage(md)

	json, err := dMessage.MarshalJSONPB(ms.jsonMarshaler)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return JSON(json), nil
}

func (ms messageServiceImpl) ToJSON(ctx context.Context, messageName messagename.MessageName, binary Binary) (JSON, error) {
	md, err := ms.grpcreflectClient.ResolveMessage(string(messageName))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dMessage := dynamic.NewMessageFactoryWithDefaults().NewDynamicMessage(md)

	err = dMessage.Unmarshal(binary)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	json, err := dMessage.MarshalJSONPB(ms.jsonMarshaler)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return json, nil
}

func (ms messageServiceImpl) ToBinary(ctx context.Context, messageName messagename.MessageName, json JSON) (Binary, error) {
	md, err := ms.grpcreflectClient.ResolveMessage(string(messageName))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dMessage := dynamic.NewMessageFactoryWithDefaults().NewDynamicMessage(md)

	err = dMessage.UnmarshalJSON(json)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	bin, err := dMessage.Marshal()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bin, nil
}

func (ms messageServiceImpl) ToDynamicMessage(ctx context.Context, messageName messagename.MessageName, json JSON) (*dynamic.Message, error) {
	md, err := ms.grpcreflectClient.ResolveMessage(string(messageName))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	dMessage := dynamic.NewMessageFactoryWithDefaults().NewDynamicMessage(md)

	err = dMessage.UnmarshalJSON(json)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dMessage, nil
}

func (ms messageServiceImpl) DynamicMessageToJSON(ctx context.Context, dm *dynamic.Message) (JSON, error) {
	json, err := dm.MarshalJSONPB(ms.jsonMarshaler)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return json, nil
}
