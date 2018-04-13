package utils

var technologies []Technology

func GetTechnolgies() []Technology {

	return technologies
}

func GetTechnolgy(name string) Technology {
	var ret Technology
	for _, x := range technologies {
		if x.ActionName == name {
			ret = x
		}
	}
	return ret
}

func SetBaseValueTechnologies() {

	//TECHNOLOGY
	technologies = append(technologies, Technology{
		Name:           "Boost Civilian production",
		Description:    "Boost civilian factory production by 15%",
		Costs:          []Cost{Cost{Type: "science", Value: 100}},
		Effects:        []Effect{Effect{ModifierName: "civilianFactoryProduction", Operator: "*", Value: 1.15}},
		ActionName:     "technoIndusT1N1",
		Tier:           1,
		TypeTechnology: "INDUSTRIAL",
	})

	technologies = append(technologies, Technology{
		Name:        "Boost Civilian production T2",
		Description: "Boost civilian factory production by 15%",
		Costs:       []Cost{Cost{Type: "science", Value: 450}},
		Effects:     []Effect{Effect{ModifierName: "civilianFactoryProduction", Operator: "*", Value: 1.15}},
		ActionName:  "technoIndusT2N1",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoIndusT1N1"},
		},
		Tier:           2,
		TypeTechnology: "INDUSTRIAL",
	})

	technologies = append(technologies, Technology{
		Name:        "Boost Civilian production T3",
		Description: "Boost civilian factory production by 25%",
		Costs:       []Cost{Cost{Type: "science", Value: 1200}},
		Effects:     []Effect{Effect{ModifierName: "civilianFactoryProduction", Operator: "*", Value: 1.25}},
		ActionName:  "technoIndusT3N1",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoIndusT2N1"},
		},
		Tier:           3,
		TypeTechnology: "INDUSTRIAL",
	})

	technologies = append(technologies, Technology{
		Name:        "Boost Tank production",
		Description: "Boost Tank factory production by 15%",
		Costs:       []Cost{Cost{Type: "science", Value: 100}},
		Effects: []Effect{
			Effect{ModifierName: "lightTankFactoryProduction", Operator: "*", Value: 1.15},
			Effect{ModifierName: "heavyTankFactoryProduction", Operator: "*", Value: 1.15},
		},
		ActionName:     "technoIndusT1N2",
		Tier:           1,
		TypeTechnology: "INDUSTRIAL",
	})
	technologies = append(technologies, Technology{
		Name:        "Boost Tank production T2",
		Description: "Boost Tank factory production by 15%",
		Costs:       []Cost{Cost{Type: "science", Value: 300}},
		Effects: []Effect{
			Effect{ModifierName: "lightTankFactoryProduction", Operator: "*", Value: 1.15},
			Effect{ModifierName: "heavyTankFactoryProduction", Operator: "*", Value: 1.15},
		},
		ActionName: "technoIndusT2N2",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoIndusT1N2"},
		},
		Tier:           2,
		TypeTechnology: "INDUSTRIAL",
	})

	technologies = append(technologies, Technology{
		Name:        "Boost Tank production T3",
		Description: "Boost Tank factory production by 15%",
		Costs:       []Cost{Cost{Type: "science", Value: 1000}},
		Effects: []Effect{
			Effect{ModifierName: "lightTankFactoryProduction", Operator: "*", Value: 1.15},
			Effect{ModifierName: "heavyTankFactoryProduction", Operator: "*", Value: 1.15},
		},
		ActionName: "technoIndusT3N2",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoIndusT2N2"},
		},
		Tier:           3,
		TypeTechnology: "INDUSTRIAL",
	})
	technologies = append(technologies, Technology{
		Name:           "Boost Aircraft production",
		Description:    "Boost Aircraft factory production by 15%",
		Costs:          []Cost{Cost{Type: "science", Value: 100}},
		Effects:        []Effect{Effect{ModifierName: "aircraftFactoryProduction", Operator: "*", Value: 1.15}},
		ActionName:     "technoIndusT1N3",
		Tier:           1,
		TypeTechnology: "INDUSTRIAL",
	})

	//MIL TECHNOLOGY

	technologies = append(technologies, Technology{
		Name:           "Boost soldier damage",
		Description:    "Boost soldier damage by 10%",
		Costs:          []Cost{Cost{Type: "science", Value: 200}},
		Effects:        []Effect{Effect{ModifierName: "soldierQuality", Operator: "*", Value: 1.10}},
		ActionName:     "technoMilT1N1",
		Tier:           1,
		TypeTechnology: "MILITARY",
	})
	technologies = append(technologies, Technology{
		Name:        "Boost soldier damage",
		Description: "Boost soldier damage by 15%",
		Costs:       []Cost{Cost{Type: "science", Value: 800}},
		Effects:     []Effect{Effect{ModifierName: "soldierQuality", Operator: "*", Value: 1.15}},
		ActionName:  "technoMilT2N1",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoMilT1N1"},
		},
		Tier:           2,
		TypeTechnology: "MILITARY",
	})
	technologies = append(technologies, Technology{
		Name:        "Boost soldier damage",
		Description: "Boost soldier damage by 20%",
		Costs:       []Cost{Cost{Type: "science", Value: 1600}},
		Effects:     []Effect{Effect{ModifierName: "soldierQuality", Operator: "*", Value: 1.20}},
		ActionName:  "technoMilT3N1",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoMilT2N1"},
		},
		Tier:           3,
		TypeTechnology: "MILITARY",
	})

	technologies = append(technologies, Technology{
		Name:           "Boost light tank damage",
		Description:    "Boost light tank damage by 10%",
		Costs:          []Cost{Cost{Type: "science", Value: 200}},
		Effects:        []Effect{Effect{ModifierName: "lightTankQuality", Operator: "*", Value: 1.10}},
		ActionName:     "technoMilT1N2",
		Tier:           1,
		TypeTechnology: "MILITARY",
	})
	technologies = append(technologies, Technology{
		Name:        "Boost light tank damage",
		Description: "Boost light tank damage by 15%",
		Costs:       []Cost{Cost{Type: "science", Value: 800}},
		Effects:     []Effect{Effect{ModifierName: "lightTankQuality", Operator: "*", Value: 1.15}},
		ActionName:  "technoMilT2N2",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoMilT1N2"},
		},
		Tier:           2,
		TypeTechnology: "MILITARY",
	})
	technologies = append(technologies, Technology{
		Name:        "Boost light tank damage",
		Description: "Boost light tank damage by 20%",
		Costs:       []Cost{Cost{Type: "science", Value: 1600}},
		Effects:     []Effect{Effect{ModifierName: "lightTankQuality", Operator: "*", Value: 1.20}},
		ActionName:  "technoMilT3N2",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoMilT2N2"},
		},
		Tier:           3,
		TypeTechnology: "MILITARY",
	})

	//ECO
	technologies = append(technologies, Technology{
		Name:        "Administration Reform - Tax",
		Description: "Reform tax administration to ensure you gather tax more efficiently",
		Costs:       []Cost{Cost{Type: "science", Value: 150}},
		Effects: []Effect{
			Effect{ModifierName: "taxEfficiency", Operator: "=", Value: 1.15},
		},
		ActionName:     "technoEcoT1N1",
		Tier:           1,
		TypeTechnology: "ECONOMIC",
	})
	technologies = append(technologies, Technology{
		Name:        "Administration Reform - Reasearch",
		Description: "Reform Reasearch administration to ensure your scientist work in better conditions",
		Costs:       []Cost{Cost{Type: "science", Value: 600}},
		Effects: []Effect{
			Effect{ModifierName: "researchEfficiency", Operator: "=", Value: 1.15},
		},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N1"},
		},
		ActionName:     "technoEcoT2N1",
		Tier:           2,
		TypeTechnology: "ECONOMIC",
	})

	technologies = append(technologies, Technology{
		Name:           "Intelligence deployment",
		Description:    "Unlock some intelligence action, like more effective sabotage",
		Costs:          []Cost{Cost{Type: "science", Value: 350}},
		ActionName:     "technoEcoT1N2",
		Tier:           1,
		TypeTechnology: "ECONOMIC",
	})
	technologies = append(technologies, Technology{
		Name:        "Advanced intelligence deployment",
		Description: "Unlock some intelligence action, like more effective sabotage and assasination",
		Costs:       []Cost{Cost{Type: "science", Value: 1000}},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N2"},
		},
		ActionName:     "technoEcoT2N2",
		Tier:           2,
		TypeTechnology: "ECONOMIC",
	})

	technologies = append(technologies, Technology{
		Name:           "Financial devlopment",
		Description:    "Unlock some financial action, like larger loans",
		Costs:          []Cost{Cost{Type: "science", Value: 400}},
		ActionName:     "technoEcoT1N3",
		Tier:           1,
		TypeTechnology: "ECONOMIC",
	})

	technologies = append(technologies, Technology{
		Name:        "Advanced Financial devlopment",
		Description: "Unlock some financial action, like larger loans",
		Costs:       []Cost{Cost{Type: "science", Value: 800}},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N3"},
		},
		ActionName:     "technoEcoT2N3",
		Tier:           2,
		TypeTechnology: "ECONOMIC",
	})

}
