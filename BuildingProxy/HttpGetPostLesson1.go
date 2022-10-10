package BuildingProxy

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetWebPage(webAddress string) {

	resp, err := http.Get(webAddress)
	if err != nil {
		log.Panicln(err)
	}
	// Print HTTP Status
	fmt.Println(resp.Status)
	// Read and display response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}
