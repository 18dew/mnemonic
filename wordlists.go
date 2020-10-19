package mnemonic

// Language ...
type Language string

// Language const
const (
	LanguageEnglish            Language = "english"
)

// GetWordList ...
func GetWordList(language Language) ([]string, error) {
	var wordlist = []string{}
	switch language {
	case LanguageEnglish:
		wordlist = WordlistEnglish
	default:
		wordlist = WordlistEnglish
	}
	return wordlist, nil
}