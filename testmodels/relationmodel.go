package testmodel

import (
	"fmt"
	"groupie/models"
	"strconv"
)

func GetRelations(id int) model.Relation {
	var relations model.Relation

	err := get("https://groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(id), &relations)
	if err != nil {
		return relations
	}
	fmt.Println(relations)
	return relations
}
