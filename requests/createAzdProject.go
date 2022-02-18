package requests

import (
	"io/ioutil"
	"log"
	"net/http"
)

func CreateAzdProjects() (responseBody string) {
	var username string = "username"
	var password string = "PAT"

	client := &http.Client{}

	request, err := http.NewRequest("POST", "https://dev.azure.com/{organization}/_apis/projects?api-version=2.0", nil)
	request.SetBasicAuth(username, password)
	// fmt.Println(response.StatusCode)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseBody = string(responseData)

	return responseBody
}
