package protowire

import (
	"github.com/shatll-s/nexelliad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NexelliadMessage_RequestPruningPointAndItsAnticone) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NexelliadMessage_RequestPruningPointAndItsAnticone is nil")
	}
	return &appmessage.MsgRequestPruningPointAndItsAnticone{}, nil
}

func (x *NexelliadMessage_RequestPruningPointAndItsAnticone) fromAppMessage(_ *appmessage.MsgRequestPruningPointAndItsAnticone) error {
	return nil
}
