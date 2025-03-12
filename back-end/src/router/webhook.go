package router

import (
	"log"
	"observer-go/src/handler"

	"github.com/gorilla/mux"
)

const (
	INFO  = "[INFO] "
	ERROR = "[ERROR] "
)

type Webhooks struct {
	Log     *log.Logger
	Router  *mux.Router
	Handler *handler.Webhook
}

func (i *Webhooks) Build() {
	//TODO: Implement the webhook routes
	i.Log.Printf("%sBuild Webhooks routes finished", INFO)
}
