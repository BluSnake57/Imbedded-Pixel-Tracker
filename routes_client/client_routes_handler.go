package routes_client

import (
	"IPT/Imbedded-Pixel-Tracker/router"
)

func Client_Routes(r router.Router) {
	router.AddRoute(&r, router.Receiver{
		Route:     "/trace_client",
		RouteType: router.RouteGet,
		Sender:    Trace_Client,
	})
}
