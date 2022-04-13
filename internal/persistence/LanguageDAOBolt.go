package persistence

import (
	"encoding/json"
	"internal/entities"
	"internal/persistence/bolt"
	"log"
)

var b = bolt.DBopen("base.db")

type LanguageDAOBolt struct{}

func NewLanguageDAOBolt() LanguageDAOBolt {
	return LanguageDAOBolt{}
}

func (LanguageDAOBolt) FindAll() []entities.Language {
	res := b.DBgetAll("languages")
	var languages []entities.Language
	for _, l := range res {
		log.Fatal(res)
		var language entities.Language
		json.Unmarshal([]byte(l), &language)
		languages = append(languages, language)
	}
	return languages
}

func (LanguageDAOBolt) Find(code string) string {
	return b.DBget("languages", code)
}

func (LanguageDAOBolt) Delete(code string) bool {
	res := b.DBget("languages", code)
	if res != "" {
		b.DBdelete("languages", code)
		return true
	}

	return false
}

func (LanguageDAOBolt) Create(language entities.Language) {
	b.DBput("languages", language.Code, language.String())
}
