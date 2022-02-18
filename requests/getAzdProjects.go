package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// A Response struct to map the Entire Response
type Response struct {
	Count uint8   `json:"count"`
	Value []Value `json:"value"`
}

// A Project Struct to map every project to.
type Value struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

//GetAzdProjects returns a list of all projects in Azure DevOps
func GetAzdProjects() (responseBody string) {
	var username string = "username"
	var password string = "PAT"

	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://dev.azure.com/{organization}/_apis/projects?api-version=2.0", nil)
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

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Value); i++ {
		fmt.Println(responseObject.Value[i].Id)
		fmt.Println(responseObject.Value[i].Name)
		fmt.Println(responseObject.Value[i].Url)
	}
	return responseBody
}
