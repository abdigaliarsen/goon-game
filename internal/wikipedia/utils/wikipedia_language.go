package utils

import "fmt"

// languageToDomain https://en.wikipedia.org/wiki/List_of_Wikipedias
var languageToDomain = map[string]string{
	"en": "https://en.wikipedia.org",
	"de": "https://de.wikipedia.org",
	"fr": "https://fr.wikipedia.org",
	"es": "https://es.wikipedia.org",
	"ja": "https://ja.wikipedia.org",
	"ru": "https://ru.wikipedia.org",
	"pt": "https://pt.wikipedia.org",
	"it": "https://it.wikipedia.org",
	"zh": "https://zh.wikipedia.org",
}

func GetWikipediaDomainByLanguage(language string) string {
	domain, ok := languageToDomain[language]
	if ok {
		return domain
	}

	return fmt.Sprintf("https://%s.wikipedia.org", language)
}
