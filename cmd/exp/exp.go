package main

import (
	"encoding/json"
	"fmt"

	"movieexample.com/metadata/pkg/model"
)

func main() {
	m := model.Metadata{
		ID:          "1",
		Title:       "Game of throne",
		Description: "Game of trhone",
		Director:    "Robert Martin",
	}
	s, _ := json.Marshal(m)
	fmt.Println(string(s))
	var struc model.Metadata
	json.Unmarshal([]byte(s), &struc)
	fmt.Println(struc)
}
