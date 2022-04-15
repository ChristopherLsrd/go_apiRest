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

func (l LanguageDAOBolt) FindAll() []entities.Language {
	res := b.DBgetAll("Languages")
	var languages []entities.Language
	for _, l := range res {
		
		var language entities.Language
		json.Unmarshal([]byte(l), &language)
		languages = append(languages, language)
	}
	return languages
}

func (l LanguageDAOBolt) Find(code string) entities.Language {
	var language entities.Language
	res:=b.DBget("Languages", code)
	json.Unmarshal([]byte(res), &language)
	return language
}

func (l LanguageDAOBolt) Delete(code string) bool {
	
	err:=	b.DBdelete("Languages", code)
		if err!=nil{
			log.Fatal(err)
			return false
		}
		return true
}

func (l LanguageDAOBolt) Create(language entities.Language) bool {
	res, _ := json.Marshal(language)
	if l.Exists(language.Code)==false{
		b.DBput("Languages", language.Code, string(res))
		return true
	}
	return  false
	
	
}


func(l LanguageDAOBolt) Exists(code string) bool{
	language:=b.DBget("Languages",code)
	if language !=""{
		
		return true
	}
	return false
}


func (l LanguageDAOBolt) Update(language entities.Language) bool {
	res, _ := json.Marshal(language)
	if l.Exists(language.Code)==true{
		b.DBdelete("Languages",language.Code)
		b.DBput("Languages", language.Code, string(res))
		return true
	}
	return  false
	
	
}