package rpchandlers

import (
	"github.com/shatll-s/nexelliad/app/appmessage"
	"github.com/shatll-s/nexelliad/app/rpc/rpccontext"
	"github.com/shatll-s/nexelliad/infrastructure/network/netadapter/router"
)

// HandleNotifyVirtualDaaScoreChanged handles the respectively named RPC command
func HandleNotifyVirtualDaaScoreChanged(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateVirtualDaaScoreChangedNotifications()

	response := appmessage.NewNotifyVirtualDaaScoreChangedResponseMessage()
	return response, nil
}
