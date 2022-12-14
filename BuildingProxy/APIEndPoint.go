package BuildingProxy

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Message string
	Status  string
}

func PostToEndPoint(endPointAddress string) {
	res, err := http.Post(
		endPointAddress,
		"application/json",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}
	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Printf("%s -> %s\n", status.Status, status.Message)
}
