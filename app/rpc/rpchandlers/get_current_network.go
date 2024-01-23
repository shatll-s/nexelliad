package rpchandlers

import (
	"github.com/shatll-s/nexelliad/app/appmessage"
	"github.com/shatll-s/nexelliad/app/rpc/rpccontext"
	"github.com/shatll-s/nexelliad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
