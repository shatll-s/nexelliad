package protowire

import (
	"github.com/Nexellia-Network/nexelliad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *NexelliadMessage_DoneHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NexelliadMessage_DoneHeaders is nil")
	}
	return &appmessage.MsgDoneHeaders{}, nil
}

func (x *NexelliadMessage_DoneHeaders) fromAppMessage(_ *appmessage.MsgDoneHeaders) error {
	return nil
}
