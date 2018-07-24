package utils

var Countries []Country

func GetCountries() *[]Country {
	return &Countries
}

func GetCountry(name string) Country {
	var ret Country
	for _, x := range Countries {
		if x.Name == name {
			ret = x
		}
	}
	return ret
}

func SetBaseValueCountries() {

	//Countries

	Countries = append(Countries, Country{
		Name:        "France",
		Flag:        "fra.png",
		Description: "Specialized in aircraft ",
		Effects: []Effect{
			Effect{ModifierType: "Modifiers", ModifierName: "dmgAerialBonus", Operator: "+", Value: 0.1, Target: "Player", ActionName: "dmgAerialBonus"},
			Effect{ModifierType: "Modifiers", ModifierName: "dmgBombBonus", Operator: "+", Value: 0.1, Target: "Player", ActionName: "dmgBombBonus"},
		},
	})
	Countries = append(Countries, Country{
		Name:        "Germany",
		Flag:        "ger.png",
		Description: "Specialized in heavy tanks and high quality army.",
		Effects: []Effect{
			Effect{ModifierType: "Army", ModifierName: "Quality", Operator: "+", Value: 5, Target: "Player", ActionName: "addQuality"},
			Effect{ModifierType: "Modifiers", ModifierName: "heavyTankQuality", Operator: "+", Value: 0.1, Target: "Player", ActionName: "heavyTankQuality"},
		},
	})

}
