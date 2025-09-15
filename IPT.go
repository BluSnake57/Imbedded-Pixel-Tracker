package main

import (
	"IPT/Imbedded-Pixel-Tracker/router"
	"IPT/Imbedded-Pixel-Tracker/routes"
	"IPT/Imbedded-Pixel-Tracker/routes_client"
	"IPT/Imbedded-Pixel-Tracker/routes_server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	//Eventually should add some authentication for system admin to limit who can access server functions

	//Eventually should add a database in case of large number of simultaneous trackers

	r := router.NewRouter(":4040") //Creates new router with specified IP

	router.AddRoute(&r, router.Receiver{
		Route:      "/ping",                    //Path to accesss through url
		RouteType:  router.RouteGet,            //Specifies type of request to expect
		Middleware: router.Client_Middleware(), //Specifies middleware to use
		Sender:     routes.Trace_Ping,          //Speicifies function to run on request
	})

	routes_client.Client_Routes(r)
	routes_server.Server_Routes(r)

	router.RunRouter(r)
}
