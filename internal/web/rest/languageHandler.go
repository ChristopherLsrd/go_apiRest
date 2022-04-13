package rest

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/persistence"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var languageDAOMem = persistence.NewLanguageDAOMem()
var languageDAOBolt = persistence.NewLanguageDAOBolt()

func LanguageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := (vars["code"])

	//data := getLanguage(code)
	//data := languageDAOMem.Find(code)
	data := languageDAOBolt.Find(code)

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "%s", j)
}

func LanguagesHandler(w http.ResponseWriter, r *http.Request) {
	//j, err := json.Marshal(languageDAOMem.FindAll())
	j, err := json.Marshal(languageDAOBolt.FindAll())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s \n", j)
}

func PostLanguageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var language entities.Language
	json.Unmarshal(reqBody, &language)
	languageDAOMem.Create(language)

	json.NewEncoder(w).Encode(language)
}

func DeleteLanguageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	languageDAOMem.Delete(code)
	/*for index, language := range languages {
	    if language.Code == code {
	        languages = append(languages[:index], languages[index+1:]...)

	    }
	}*/

}

func PutLanguageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var language entities.Language

	json.Unmarshal(reqBody, &language)
	languageDAOMem.Update(language)
	/*for i,languageFor:=range languages{
		if language.Code==languageFor.Code{
			languages[i]=language
		}
	}*/
}

/*
func getLanguage(code string) entities.Language{

	for _,language:= range languages{
		if language.Code == code{
			return language
		}
	}

	return entities.NewLanguage("","")
}*/
