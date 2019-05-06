
package sv

const locale = "sv"

// LanguagePack is an implementation of the generic golem.LanguagePack interface for sv
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

