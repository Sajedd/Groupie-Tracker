package testmodel

import (
	"encoding/json"
	"fmt"
	"groupie/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetDates(id int) (*model.Dates, error) {
	client = &http.Client{}
	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/dates"+"/"+strconv.Itoa(id), nil)
	
	if err != nil {
		fmt.Print(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject model.Dates
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Println(responseObject.Dates)
	return &responseObject, err
}
