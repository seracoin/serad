package protowire

import (
	"github.com/seracoin/serad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SeradMessage_NotifyFinalityConflictsRequest) toAppMessage() (appmessage.Message, error) {
	return &appmessage.NotifyFinalityConflictsRequestMessage{}, nil
}

func (x *SeradMessage_NotifyFinalityConflictsRequest) fromAppMessage(_ *appmessage.NotifyFinalityConflictsRequestMessage) error {
	x.NotifyFinalityConflictsRequest = &NotifyFinalityConflictsRequestMessage{}
	return nil
}

func (x *SeradMessage_NotifyFinalityConflictsResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SeradMessage_NotifyFinalityConflictsResponse is nil")
	}
	return x.NotifyFinalityConflictsResponse.toAppMessage()
}

func (x *SeradMessage_NotifyFinalityConflictsResponse) fromAppMessage(message *appmessage.NotifyFinalityConflictsResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.NotifyFinalityConflictsResponse = &NotifyFinalityConflictsResponseMessage{
		Error: err,
	}
	return nil
}

func (x *NotifyFinalityConflictsResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NotifyFinalityConflictsResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.NotifyFinalityConflictsResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *SeradMessage_FinalityConflictNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SeradMessage_FinalityConflictNotification is nil")
	}
	return x.FinalityConflictNotification.toAppMessage()
}

func (x *SeradMessage_FinalityConflictNotification) fromAppMessage(message *appmessage.FinalityConflictNotificationMessage) error {
	x.FinalityConflictNotification = &FinalityConflictNotificationMessage{
		ViolatingBlockHash: message.ViolatingBlockHash,
	}
	return nil
}

func (x *FinalityConflictNotificationMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "FinalityConflictNotificationMessage is nil")
	}
	return &appmessage.FinalityConflictNotificationMessage{
		ViolatingBlockHash: x.ViolatingBlockHash,
	}, nil
}

func (x *SeradMessage_FinalityConflictResolvedNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SeradMessage_FinalityConflictResolvedNotification is nil")
	}
	return x.FinalityConflictResolvedNotification.toAppMessage()
}

func (x *SeradMessage_FinalityConflictResolvedNotification) fromAppMessage(message *appmessage.FinalityConflictResolvedNotificationMessage) error {
	x.FinalityConflictResolvedNotification = &FinalityConflictResolvedNotificationMessage{
		FinalityBlockHash: message.FinalityBlockHash,
	}
	return nil
}

func (x *FinalityConflictResolvedNotificationMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "FinalityConflictResolvedNotificationMessage is nil")
	}
	return &appmessage.FinalityConflictResolvedNotificationMessage{
		FinalityBlockHash: x.FinalityBlockHash,
	}, nil
}
