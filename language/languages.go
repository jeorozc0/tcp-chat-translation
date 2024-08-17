package lang

const (
	English  = "english"
	Spanish  = "spanish"
	Italian  = "italian"
	Japanese = "japanese"
	French   = "french"
)

var Languages = []string{English, Spanish, Italian, Japanese, French}

func GetLanguages() []string {
	return Languages
}

// IsValidLanguage checks if a given language is supported
func IsValidLanguage(language string) bool {
	for _, l := range Languages {
		if l == language {
			return true
		}
	}
	return false
}
