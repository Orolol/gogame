package utils

var technologies []Technology
var actions []PlayerActionOrder
var policies []Policy

func GetTechnolgies() []Technology {

	return technologies
}
func GetActions() []PlayerActionOrder {
	return actions
}
func GetPolicies() []Policy {
	return policies
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

func GetAction(name string) PlayerActionOrder {
	var ret PlayerActionOrder
	for _, x := range actions {
		if x.ActionName == name {
			ret = x
		}
	}
	return ret
}

func GetPolicy(name string) Policy {
	var ret Policy
	for _, x := range policies {
		if x.ActionName == name {
			ret = x
		}
	}
	return ret
}

func SetBaseValueDB() {

	//POLICIES
	policies = append(policies, Policy{
		Name:          "Training Time",
		ActionName:    "setPopRecPolicy",
		Description:   "Set your recuitement policy",
		TypePolicy:    "MIL",
		PossibleValue: "{\"Full\" : 1,\"Long\" : 2,\"Hurry\" : 5,\"No time !\" : 10,\"Send everyone !\" : 30}",
		DefaultValue:  "1"})
	policies = append(policies, Policy{
		Name:          "Conscription Policy",
		ActionName:    "setConscPolicy",
		Description:   "Set your recuitement policy",
		TypePolicy:    "MIL",
		PossibleValue: "{\"Pro Army\" : 1,\"Volonteer\" : 2,\"War time\" : 5,\"All valids !\" : 10,\"Anyone who can hold a weapon\" : 30}",
		DefaultValue:  "1"})

	policies = append(policies, Policy{
		Name:          "Tax rate",
		ActionName:    "setTaxRatePolicy",
		Description:   "Set your tax rate. ",
		TypePolicy:    "ECO",
		PossibleValue: "{\"Low taxes\" : 1,\"Country effort\" : 1.5,\"War Economy\" : 2,\"Full Mobilization\" : 3,\"Total war\" : 5}",
		DefaultValue:  "1.5"})
	policies = append(policies, Policy{
		Name:          "Build Light Tank ?",
		ActionName:    "setBuildLgtTank",
		Description:   "Set your tax rate. ",
		TypePolicy:    "ECO",
		PossibleValue: "{\"Yes\" : 1,\"No\" : 0}",
		DefaultValue:  "1"})

	policies = append(policies, Policy{
		Name:          "Build Heavy Tank ?",
		ActionName:    "setBuildHvyTank",
		Description:   "Set your tax rate. ",
		TypePolicy:    "ECO",
		PossibleValue: "{\"Yes\" : 1,\"No\" : 0}",
		DefaultValue:  "1"})

	//ACTIONS
	actions = append(actions, PlayerActionOrder{
		Name:        "Civ Fact -> Lght Fact",
		ActionName:  "actionCivConvertFactoryToLightTankFact",
		Description: "Convert Civilian Factory to light Tank factory (Cost 1M) ",
		Cooldown:    10,
		Costs:       []Cost{Cost{Type: "money", Value: 1000000}},
	})

	actions = append(actions, PlayerActionOrder{
		Name:        "Civ Fact -> Hvy Fact",
		ActionName:  "actionCivConvertFactoryToHvyTankFact",
		Description: "Convert Civilian Factory to Heavy Tank factory (Cost 1M) ",
		Cooldown:    10,
		Costs:       []Cost{Cost{Type: "money", Value: 1000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Name:        "War Propaganda",
		ActionName:  "actionWarPropaganda",
		Description: "Boost morale by 15% (cost 10M) ",
		Cooldown:    10,
		Costs:       []Cost{Cost{Type: "money", Value: 10000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Name:        "Emergency recruitment",
		ActionName:  "emergencyRecruitment",
		Description: "Recruit immediatly 10% of your manpower. (cost 50M and 10% morale)",
		Cooldown:    30,
		Costs:       []Cost{Cost{Type: "money", Value: 50000000}, Cost{Type: "morale", Value: 10}},
	})
	actions = append(actions, PlayerActionOrder{
		Name:        "Purge of weak elements",
		ActionName:  "purgeSoldier",
		Description: "Executes possible traitors and weak soldier to improve morale and discipline by 15%. (cost 15% soldier and 10M)",
		Cooldown:    30,
		Costs:       []Cost{Cost{Type: "money", Value: 10000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Name:       "Buy foreign tanks",
		ActionName: "buyForeignTanks",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoIndusT1N1"},
			Constraint{Type: "tech", Value: "technoIndusT1N2"},
		},
		Description: "Get 150 light tank and 50 heavy one (cost 100M) ",
		Cooldown:    60,
		Costs:       []Cost{Cost{Type: "money", Value: 10000000}},
	})

	//TECHNOLOGY
	technologies = append(technologies, Technology{
		Name:           "Boost Civilian production",
		Description:    "Boost civilian factory production by 15%",
		Costs:          []Cost{Cost{Type: "science", Value: 100}},
		Effects:        []Effect{Effect{ModifierName: "civilianFactoryProduction", Operator: "*", Value: 1.15}},
		ActionName:     "technoIndusT1N1",
		Tier:           1,
		TypeTechnology: "INDUS",
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
		TypeTechnology: "INDUS",
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
		TypeTechnology: "INDUS",
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
		TypeTechnology: "INDUS",
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
		TypeTechnology: "INDUS",
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
		TypeTechnology: "INDUS",
	})
	technologies = append(technologies, Technology{
		Name:           "Boost Aircraft production",
		Description:    "Boost Aircraft factory production by 15%",
		Costs:          []Cost{Cost{Type: "science", Value: 100}},
		Effects:        []Effect{Effect{ModifierName: "aircraftFactoryProduction", Operator: "*", Value: 1.15}},
		ActionName:     "technoIndusT1N3",
		Tier:           1,
		TypeTechnology: "INDUS",
	})

	//MIL TECHNOLOGY

	technologies = append(technologies, Technology{
		Name:           "Boost soldier damage",
		Description:    "Boost soldier damage by 10%",
		Costs:          []Cost{Cost{Type: "science", Value: 200}},
		Effects:        []Effect{Effect{ModifierName: "soldierQuality", Operator: "*", Value: 1.10}},
		ActionName:     "technoMilT1N1",
		Tier:           1,
		TypeTechnology: "MIL",
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
		TypeTechnology: "MIL",
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
		TypeTechnology: "MIL",
	})

	technologies = append(technologies, Technology{
		Name:           "Boost light tank damage",
		Description:    "Boost light tank damage by 10%",
		Costs:          []Cost{Cost{Type: "science", Value: 200}},
		Effects:        []Effect{Effect{ModifierName: "lightTankQuality", Operator: "*", Value: 1.10}},
		ActionName:     "technoMilT1N2",
		Tier:           1,
		TypeTechnology: "MIL",
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
		TypeTechnology: "MIL",
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
		TypeTechnology: "MIL",
	})

}
