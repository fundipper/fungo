package conf

import (
	"log"
)

type Theme map[string]interface{}

func NewTheme() *Theme {
	data := Theme{}
	err := v.UnmarshalKey(MODEL_THMEM, &data)
	if err != nil {
		log.Fatal(err)
	}
	return &data
}
