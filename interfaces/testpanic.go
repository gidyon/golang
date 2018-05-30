package main

import (
	"errors"
	"fmt"
	"strings"
)

type Creds struct {
	email    string
	password string
}

type Profile struct{}

func (p *Creds) Login() (*Profile, error) {
	if len(p.password) > 6 && strings.HasSuffix(p.email, "@gmail.com") {
		return &Profile{}, nil
	}
	return nil, errors.New("invalid credentials")
}
func main() {
	var creds *Creds = &Creds{"Gideon", "gideon@gmailcom"}
	_, err := creds.Login()
	if err != nil {
		fmt.Println(err)
	}
}
