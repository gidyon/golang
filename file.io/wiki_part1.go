package main

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	return ioutil.WriteFile(p.Title, p.Body, 0666)
}

func (p *Page) Load(title string) ([]byte, error) {
	return ioutil.ReadFile(title)
}
