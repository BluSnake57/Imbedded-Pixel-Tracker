package routes_server

import "IPT/Imbedded-Pixel-Tracker/router"

func Server_Routes(r router.Router) {
	router.AddRoute(&r, router.Receiver{
		Route:     "/initalize_tracker",
		RouteType: router.RoutePost,
		Sender:    Initalize_Tracker,
	})

	router.AddRoute(&r, router.Receiver{
		Route:     "/get_status",
		RouteType: router.RoutePost,
		Sender:    Get_Status,
	})
}
