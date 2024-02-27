package protowire

import (
	"github.com/seracoin/serad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SeradMessage_GetSubnetworkRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SeradMessage_GetSubnetworkRequest is nil")
	}
	return x.GetSubnetworkRequest.toAppMessage()
}

func (x *SeradMessage_GetSubnetworkRequest) fromAppMessage(message *appmessage.GetSubnetworkRequestMessage) error {
	x.GetSubnetworkRequest = &GetSubnetworkRequestMessage{
		SubnetworkId: message.SubnetworkID,
	}
	return nil
}

func (x *GetSubnetworkRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetSubnetworkRequestMessage is nil")
	}
	return &appmessage.GetSubnetworkRequestMessage{
		SubnetworkID: x.SubnetworkId,
	}, nil
}

func (x *SeradMessage_GetSubnetworkResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SeradMessage_GetSubnetworkResponse is nil")
	}
	return x.GetSubnetworkResponse.toAppMessage()
}

func (x *SeradMessage_GetSubnetworkResponse) fromAppMessage(message *appmessage.GetSubnetworkResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.GetSubnetworkResponse = &GetSubnetworkResponseMessage{
		GasLimit: message.GasLimit,
		Error:    err,
	}
	return nil
}

func (x *GetSubnetworkResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetSubnetworkResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	if rpcErr != nil && x.GasLimit != 0 {
		return nil, errors.New("GetSubnetworkResponseMessage contains both an error and a response")
	}

	return &appmessage.GetSubnetworkResponseMessage{
		GasLimit: x.GasLimit,
		Error:    rpcErr,
	}, nil
}
