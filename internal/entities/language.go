package entities

import (
	"fmt"
)

type Language struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
}

func NewLanguage(code string, name string) Language {
	language := Language{
		Code: code,
		Name: name,
	}
	return language
}

func (l Language) String() string {
	return fmt.Sprintf("%s %s", l.Code, l.Name)
}
