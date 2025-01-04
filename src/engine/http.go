package engine

import (
	"fmt"
	"log"
	"net/http"
)

func StartSite(path string, port int) {

	formattedPort := fmt.Sprintf(":%d", port)	
	server := http.FileServer(http.Dir(path))

	log.Printf("Server listening on %s....", formattedPort)
	log.Print("Site successfully started.")

	http.Handle("/", server)
	err := http.ListenAndServe(formattedPort, nil)

	if err != nil {
		log.Fatal(err)
	}
}
