package main

import (
	"flag"
	"os"
	"text/template"
)

type data struct {
	Locale string
}

func main() {
	var d data
	flag.StringVar(&d.Locale, "locale", "", "The locale abbreviation this language pack is generated for")
	flag.Parse()

	t := template.Must(template.New("pack").Parse(packTemplate))
	t.Execute(os.Stdout, d)
}

// TODO: add generated template
var packTemplate = `
package {{.Locale}}

const locale = "{{.Locale}}"

// LanguagePack is an implementation of the generic golem.LanguagePack interface for {{.Locale}}
type LanguagePack struct {
}

// NewPackage creates a language pack
func NewPackage() *LanguagePack {
	return &LanguagePack{}
}

// GetResource returns the dictionary of lemmatized words
func (l *LanguagePack) GetResource() ([]byte, error) {
	return Asset("data/" + locale)
}

// GetLocale returns the language name
func (l *LanguagePack) GetLocale() string {
	return locale
}

`
