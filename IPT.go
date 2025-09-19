package main

import (
	Database "IPT/Imbedded-Pixel-Tracker/database"
	"IPT/Imbedded-Pixel-Tracker/router"
	"IPT/Imbedded-Pixel-Tracker/routes"
	"IPT/Imbedded-Pixel-Tracker/routes_client"
	"IPT/Imbedded-Pixel-Tracker/routes_server"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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

	go command_line_interface()
	router.RunRouter(r)

}

// IK all the cli code is garbage I added it last second if you want feel free to make it better
func command_line_interface() {
	reader := bufio.NewReader(os.Stdin)
	for {
		print(">")
		var command string
		var clientID string

		input, _ := reader.ReadString('\n')
		input = strings.Split(input, "\n")[0]

		//n, err := fmt.Scan(&input)
		//log.Print(n, err)
		log.Println(input)

		input_list := strings.Split(input, " ")
		command = input_list[0]
		if len(strings.Split(input, " ")) > 1 {
			clientID = input_list[1]
		}

		switch strings.ToLower(command) {
		case "help":
			fmt.Println("\nCommands are:")
			fmt.Println("Create [ClientID]")
			fmt.Println("Status [ClientID]")
			fmt.Println("Delete [ClientID]")
		case "create":
			url, err := create_tracker(clientID)
			if err != nil {
				fmt.Println("\nCommand Failed because:", err)
			} else {
				fmt.Println("Tracker URL:", url)
			}
		case "status":
			tracker, err := tracker_status(clientID)
			if err != nil {
				fmt.Println("\nCommand Failed because:", err)
			} else {
				fmt.Println(tracker)
			}
		case "delete":
			err := delete_tracker(clientID)
			if err != nil {
				fmt.Println("\nCommand Failed because:", err)
			}
		}
	}
}

func create_tracker(clientID string) (string, error) {
	url, err := Database.Generate_Tracker(clientID)

	return url, err
}

func tracker_status(clientID string) (Database.Tracker, error) {
	tracker, err := Database.Tracker_Status(clientID)
	return tracker, err
}

func delete_tracker(clientID string) error {
	err := Database.Remove_Tracker(clientID)
	return err
}
