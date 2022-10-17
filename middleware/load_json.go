package middleware

import (
	"encoding/json"
	"io/ioutil"
	"pokemon/model"
)

func GetDataFromJson() *[]model.Pokemon {
	file, _ := ioutil.ReadFile("pokedex.json")
	data := make([]model.Pokemon, 1)
	_ = json.Unmarshal(file, &data)
	return &data
}
