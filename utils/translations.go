package utils

var translations []Translation

func GetTranslations() *[]Translation {
	return &translations
}

func GetTranslation(name string) Translation {
	var ret Translation
	for _, x := range translations {
		if x.ActionName == name {
			ret = x
		}
	}
	return ret
}
func GetTranslationsByLanguage(language string) []Translation {

	var ret []Translation
	for _, x := range translations {
		if x.Language == language {
			ret = append(ret, x)
		}
	}
	return ret
}

func SetBaseValueTranslations() {

	//EVENTS

	translations = append(translations, Translation{
		Language:    "en",
		ActionName:  "NbSoldier",
		ShortName:   "Soldiers",
		LongName:    "Number of soldiers",
		Description: "How many soldier are fighting for you",
	})
	translations = append(translations, Translation{
		Language:    "en",
		ActionName:  "NbLigtTank",
		ShortName:   "Light tanks",
		LongName:    "Number of light tanks",
		Description: "How many light tanks are fighting for you",
	})
	translations = append(translations, Translation{
		Language:    "en",
		ActionName:  "NbHvyTank",
		ShortName:   "Heavy tanks",
		LongName:    "Number of heavy tanks",
		Description: "How many heavy tanks are fighting for you",
	})
	translations = append(translations, Translation{
		Language:    "en",
		ActionName:  "NbArt",
		ShortName:   "Artillery",
		LongName:    "Number of artillery",
		Description: "How many artillery are fighting for you",
	})
	translations = append(translations, Translation{
		Language:    "en",
		ActionName:  "NbAirSup",
		ShortName:   "Fighter aircraft",
		LongName:    "Number of fighter aircraft",
		Description: "How many fighter aircraft are fighting for you",
	})
	translations = append(translations, Translation{
		Language:    "en",
		ActionName:  "NbAirBomb",
		ShortName:   "Bomber aircraft",
		LongName:    "Number of bomber aircraft",
		Description: "How many bomber aircraft are fighting for you",
	})

}
