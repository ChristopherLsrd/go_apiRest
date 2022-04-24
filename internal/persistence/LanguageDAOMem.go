package persistence

import (
	"internal/entities"
)

var _ LanguageDAO = (*LanguageDAOMem)(nil)

var languages []entities.Language = []entities.Language{
	entities.NewLanguage("21", "Go"), entities.NewLanguage("12", "Python"),
}

type LanguageDAOMem struct{}

func NewLanguageDAOMem() LanguageDAOMem {

	return LanguageDAOMem{}
}

func (LanguageDAOMem) FindAll() []entities.Language {
	return languages
}

func (LanguageDAOMem) Find(code string) entities.Language {
	for i, languageFor := range languages {
		if languageFor.Code == code {
			return languages[i]
		}
	}
	return entities.Language{"", ""}
}

func (LanguageDAOMem) Exists(code string) bool {
	for _, languageFor := range languages {
		if languageFor.Code == code {
			return true
		}
	}
	return false
}

func (LanguageDAOMem) Delete(code string) bool {
	for i, languageFor := range languages {
		if languageFor.Code == code {
			languages = append(languages[:i], languages[i+1:]...)
			return true
		}
	}
	return false
}

func (LanguageDAOMem) Create(language entities.Language) bool {
	for _, languageFor := range languages {
		if languageFor.Code == language.Code {
			return false
		}
	}
	languages = append(languages, language)
	return true
}

func (LanguageDAOMem) Update(language entities.Language) bool {
	for i, languageFor := range languages {
		if languageFor.Code == language.Code {
			languages[i] = language
			return true
		}
	}

	return false
}
