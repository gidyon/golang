package main

import (
	"encoding/json"
	"os"
	"reflect"
	"unicode/utf8"
)

type Profile struct {
	max string `json:name,`
	age string `json:age,`
}

func (p *Profile) MarshalJSON() (json []byte, err error) {
	json = append(json, '{')
	value := reflect.ValueOf(p)
	for i := 0; i < value.NumField(); i++ {
		str := value.Field(i).String()
		for len(str) > 0 {
			r, size := utf8.DecodeLastRuneInString(str)
			json = append(json, byte(r))
			str = str[:len(str)-size]
		}
	}
	json = append(json, '}')
	return
}

func main() {
	json.NewEncoder(os.Stdout).Encode(Profile{"max", "ten"})
}
