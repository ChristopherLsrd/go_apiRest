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

// swagger:operation GET /languages/{code} Languages GetLanguage
// ---
// summary: Return a language.
// description: Return a language if code provided match with one language in the database.
// parameters:
// - name: code
//   in: path
//   description: language code
//   type: string
//   required: true
// responses:
//   200: Language returned
func LanguageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := (vars["code"])

	//data := languageDAOMem.Find(code)
	data := languageDAOBolt.Find(code)

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "%s", j)
}

// swagger:operation GET /languages Languages GetLanguages
// ---
// summary: Return all the languages.
// description: Return all the languages.
// responses:
//   "200": Languages returned
func LanguagesHandler(w http.ResponseWriter, r *http.Request) {
	//j, err := json.Marshal(languageDAOMem.FindAll())
	j, err := json.Marshal(languageDAOBolt.FindAll())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s \n", j)
}

// swagger:operation POST /languages Languages PostLanguage
// ---
// summary: Add new language.
// description: Add a new language if code provided not existing.
// parameters:
// - name: language
//   description: language to add
//   in: body
//   required: true
//   schema:
//     "$ref": "#/internal/entities/Language"
// responses:
//   "200": Language added
func PostLanguageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var language entities.Language
	json.Unmarshal(reqBody, &language)
	//languageDAOMem.Create(language)
	languageDAOBolt.Create(language)
	json.NewEncoder(w).Encode(language)
}

// swagger:operation DELETE /languages/{code} Languages DeleteLanguage
// ---
// summary: Delete a language.
// description: Delete an existing language.
// parameters:
// - name: code
//   description: language code
//   in: path
//   required: true
// responses:
//   "200": Language deleted
func DeleteLanguageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	//languageDAOMem.Delete(code)
	languageDAOBolt.Delete(code)

}

// swagger:operation PUT /languages Languages PutLanguage
// ---
// summary: Modify an existing language.
// description: Modify an existing language.
// parameters:
// - name: language
//   description: language to modify
//   in: body
//   required: true
//   schema:
//     "$ref": "#/internal/entities/Language"
// responses:
//   "200": Student modified
func PutLanguageHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var language entities.Language

	json.Unmarshal(reqBody, &language)
	//languageDAOMem.Update(language)
	languageDAOBolt.Update(language)

}
