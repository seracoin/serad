package protowire

import (
	"github.com/seracoin/serad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SeradMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SeradMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *SeradMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
