package testmodel

import (
	"encoding/json"
	"fmt"
	"groupie/models"
	"net/http"
	"strconv"
)

const artistsUrl string = "https://groupietrackers.herokuapp.com/api/artists"

var client = &http.Client{}

func GetArtistsID(id int) *model.Artist {
	artist := &model.Artist{}
	err := get(artistsUrl+"/"+strconv.Itoa(id), &artist)
	if err != nil {
		return nil
	}
	fmt.Println(artist.Id)
	return artist

}

func get(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(target)

	if err != nil {
		return err
	}

	return nil
}
