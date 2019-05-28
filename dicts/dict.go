package dicts

// LanguagePack is what each language should implement
type LanguagePack interface {
	GetResource() ([]byte, error)
	GetLocale() string
}
