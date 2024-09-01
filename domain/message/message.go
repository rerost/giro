package message

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rerost/giro/domain/grpcreflectiface"
	"github.com/rerost/giro/domain/messagename"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type JSON []byte
type Binary []byte

type MessageService interface {
	EmptyJSON(ctx context.Context, messageName messagename.MessageName) (JSON, error)
	RequestExample(ctx context.Context, serviceName string, methodName string) (JSON, error)
	ToJSON(ctx context.Context, messageName messagename.MessageName, binary Binary) (JSON, error)
	ToBinary(ctx context.Context, messageName messagename.MessageName, json JSON) (Binary, error)
	// NOTE: For internal.
	ToDynamicMessage(ctx context.Context, messageName messagename.MessageName, json JSON) (proto.Message, error)
}

type messageServiceImpl struct {
	grpcreflectClient   grpcreflectiface.Client
	messageNameResolver messagename.MessageNameResolver
	jsonMarshaler       protojson.MarshalOptions
}

func NewMessageService(client grpcreflectiface.Client, messageNameResolver messagename.MessageNameResolver) MessageService {
	return &messageServiceImpl{
		grpcreflectClient:   client,
		messageNameResolver: messageNameResolver,
		jsonMarshaler: protojson.MarshalOptions{
			EmitDefaultValues: true,
		},
	}
}

func (ms messageServiceImpl) EmptyJSON(ctx context.Context, messageName messagename.MessageName) (JSON, error) {
	md, err := ms.grpcreflectClient.ResolveMessage(string(messageName))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	json, err := ms.jsonMarshaler.Marshal(md)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return JSON(json), nil
}

func (ms messageServiceImpl) RequestExample(ctx context.Context, serviceName string, methodName string) (JSON, error) {
	requestMessageName, err := ms.messageNameResolver.RequestMessageName(ctx, serviceName, methodName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := ms.EmptyJSON(ctx, requestMessageName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}

func (ms messageServiceImpl) ToJSON(ctx context.Context, messageName messagename.MessageName, binary Binary) (JSON, error) {
	md, err := ms.grpcreflectClient.ResolveMessage(string(messageName))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	json, err := ms.jsonMarshaler.Marshal(md)
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

	bin, err := proto.Marshal(md)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bin, nil
}

func (ms messageServiceImpl) ToDynamicMessage(ctx context.Context, messageName messagename.MessageName, json JSON) (proto.Message, error) {
	md, err := ms.grpcreflectClient.ResolveMessage(string(messageName))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := protojson.Unmarshal(json, md); err != nil {
		return nil, errors.WithStack(err)
	}

	return md, nil
}
