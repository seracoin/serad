package protowire

import (
	"github.com/seracoin/serad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SeradMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SeradMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *SeradMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
