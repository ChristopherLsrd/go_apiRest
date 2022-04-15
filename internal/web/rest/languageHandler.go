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
	log.Fatal("ccc")
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
	//languageDAOMem.Create(language)
	languageDAOBolt.Create(language)
	json.NewEncoder(w).Encode(language)
}

func DeleteLanguageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	
	//languageDAOMem.Delete(code)
	languageDAOBolt.Delete(code)

}

func PutLanguageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var language entities.Language

	json.Unmarshal(reqBody, &language)
	//languageDAOMem.Update(language)
	languageDAOBolt.Update(language)

}

