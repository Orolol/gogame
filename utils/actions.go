package utils

var actions []PlayerActionOrder

func GetActions() []PlayerActionOrder {
	return actions
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

func SetBaseValueActions() {

	//ACTIONS

	actions = append(actions, PlayerActionOrder{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "War Propaganda",
		ActionName:  "actionWarPropaganda",
		Description: "Boost morale by 15%",
		Cooldown:    10,
		Effects:     []Effect{Effect{ModifierType: "Army", ModifierName: "Morale", Operator: "+", Value: 10, Target: "Player", ActionName: "addMorale"}},
		Costs:       []Cost{Cost{Type: "money", Value: 15000000}},
		Selector:    "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Promote war heroes",
		ActionName:  "promoteWarHearoes",
		Description: "Boost morale by 5%, quality by 5%",
		Cooldown:    15,
		Effects: []Effect{
			Effect{ModifierType: "Army", ModifierName: "Morale", Operator: "+", Value: 5, Target: "Player", ActionName: "addMorale"},
			Effect{ModifierType: "Army", ModifierName: "Quality", Operator: "+", Value: 2, Target: "Player", ActionName: "addQuality"},
		},
		Costs:    []Cost{Cost{Type: "money", Value: 15000000}},
		Selector: "fixed",
	})

	actions = append(actions, PlayerActionOrder{
		Type:        "MILITARY",
		SubType:     "COMMANDMENT",
		Name:        "Purge of weak elements",
		ActionName:  "purgeSoldier",
		Description: "Executes possible traitors and weak soldier (10% of your soldiers) to improve quality by 10%.",
		Cooldown:    30,
		Effects: []Effect{
			Effect{ModifierType: "Army", ModifierName: "NbSoldier", Operator: "*", Value: 0.90, Target: "Player", ActionName: "sacrificeSoldier", ToolTipValue: 10},
			Effect{ModifierType: "Army", ModifierName: "Quality", Operator: "+", Value: 15, Target: "Player", ActionName: "addQuality"},
		},
		Costs:    []Cost{Cost{Type: "money", Value: 10000000}},
		Selector: "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "MILITARY",
		SubType:     "RECRUITMENT",
		Name:        "Emergency recruitment",
		ActionName:  "emergencyRecruitment",
		Description: "Recruit immediatly 15 000 soldier.",
		Cooldown:    30,
		Effects:     []Effect{Effect{ModifierType: "Army", ModifierName: "NbSoldier", Operator: "+", Value: 15000, Target: "Player"}},
		Costs:       []Cost{Cost{Type: "money", Value: 50000000}, Cost{Type: "morale", Value: 10}},
		Selector:    "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "MILITARY",
		SubType:     "RECRUITMENT",
		Name:        "Build additionnal barracks",
		ActionName:  "buildBarracks",
		Description: "Build two barrack",
		Cooldown:    5,
		Effects:     []Effect{Effect{ModifierType: "Territory", ModifierName: "Barracks", Operator: "+", Value: 2, Target: "Player"}},
		Costs:       []Cost{Cost{Type: "money", Value: 20000000}},
		Selector:    "fixed",
	})

	actions = append(actions, PlayerActionOrder{
		Type:       "MILITARY",
		SubType:    "RECRUITMENT",
		Name:       "Buy foreign tanks",
		ActionName: "buyForeignTanks",
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoIndusT1N1"},
			Constraint{Type: "tech", Value: "technoIndusT1N2"},
		},
		Description: "Get 150 light tank and 50 heavy one (cost 100M) ",
		Cooldown:    60,
		Costs:       []Cost{Cost{Type: "money", Value: 10000000}},
		Selector:    "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "INTELLIGENCE",
		SubType:     "SABOTAGE",
		Name:        "Send spys to assassinate key scientists",
		ActionName:  "assassinateScientist",
		Description: "Send spys to assassinate key scientists",
		Cooldown:    30,
		Effects:     []Effect{Effect{ModifierType: "Civilian", ModifierName: "NbScientist", Operator: "-", Value: 30, Target: "Opponent"}},
		Costs:       []Cost{Cost{Type: "money", Value: 45000000}},
		Selector:    "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "INTELLIGENCE",
		SubType:     "SABOTAGE",
		Name:        "Assassinate scientists and destroy laboratories",
		ActionName:  "advancedAssassinateScientist",
		Description: "Send spys to assassinate key scientists and destroy laboratories",
		Cooldown:    30,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbScientist", Operator: "-", Value: 50, Target: "Opponent"},
			Effect{ModifierType: "Modifiers", ModifierName: "labDestroyed", Operator: "+", Value: 10, Target: "Opponent"},
		},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N2"},
		},
		Costs:    []Cost{Cost{Type: "money", Value: 75000000}},
		Selector: "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "INTELLIGENCE",
		SubType:     "SABOTAGE",
		Name:        "Sabotage factories",
		ActionName:  "sabotageFactories",
		Description: "Pay foreign workers to sabotage factories",
		Cooldown:    30,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbCivilianFactory", Operator: "-", Value: 2, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbLightFactory", Operator: "-", Value: 1, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbHeavyFactory", Operator: "-", Value: 1, Target: "Opponent"},
		},
		Costs:    []Cost{Cost{Type: "money", Value: 60000000}},
		Selector: "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "INTELLIGENCE",
		SubType:     "SABOTAGE",
		Name:        "Workers Strikes",
		ActionName:  "sabotageFactoriesAdvanced",
		Description: "Pay foreign workers to sabotage factories and go on extensive stike",
		Cooldown:    30,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbCivilianFactory", Operator: "-", Value: 5, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbLightFactory", Operator: "-", Value: 2, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbHeavyFactory", Operator: "-", Value: 2, Target: "Opponent"},
		},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N2"},
		},
		Costs:    []Cost{Cost{Type: "money", Value: 90000000}},
		Selector: "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "INTELLIGENCE",
		SubType:     "SABOTAGE",
		Name:        "Worker Uprising",
		ActionName:  "sabotageFactoriesAdvanced",
		Description: "Pay foreign workers to sabotage factories and provoke country wise worker uprising",
		Cooldown:    30,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbCivilianFactory", Operator: "-", Value: 5, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbLightFactory", Operator: "-", Value: 2, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbHeavyFactory", Operator: "-", Value: 2, Target: "Opponent"},
		},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N2"},
		},
		Costs:    []Cost{Cost{Type: "money", Value: 90000000}},
		Selector: "fixed",
	})

	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "TAX",
		Name:        "Nationalize private factory",
		ActionName:  "nationalizeTwoFact",
		Description: "Convert Civilian Factory to military production. 2 new factories are available",
		Cooldown:    10,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbCivilianFactory", Operator: "+", Value: 2, Target: "Player"},
		},
		Costs:    []Cost{Cost{Type: "money", Value: 1000000}},
		Selector: "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "TAX",
		Name:        "Nationalize private factory",
		ActionName:  "nationalizeFiveFact",
		Description: "Convert Civilian Factory to military production. 5 new factories are available",
		Cooldown:    10,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbCivilianFactory", Operator: "+", Value: 5, Target: "Player"},
		},
		Costs: []Cost{Cost{Type: "money", Value: 3000000}},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N1"},
		},
		Selector: "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "Light tank production",
		ActionName:  "LightTankProduction",
		Description: "Select the percentage of factories production which go toward light tank.",
		Cooldown:    0,
		Constraints: []Constraint{
			Constraint{Type: "custom", Value: "0", Operator: ">"},
			Constraint{Type: "custom", Value: "200", Operator: "<"},
			Constraint{Type: "linked", Value: "production"},
		},
		Effects: []Effect{
			Effect{ModifierType: "Economy", ModifierName: "LightTankProduction", Operator: "=", Value: 0, Target: "Player"},
		},
		Selector:  "range",
		BaseValue: 100,
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "Heavy tank production",
		ActionName:  "HeavyTankProduction",
		Description: "Select the percentage of factories production which go toward Heavy tank.",
		Cooldown:    0,
		Constraints: []Constraint{
			Constraint{Type: "custom", Value: "0", Operator: ">"},
			Constraint{Type: "custom", Value: "200", Operator: "<"},
			Constraint{Type: "linked", Value: "production"},
		},
		Effects: []Effect{
			Effect{ModifierType: "Economy", ModifierName: "HeavyTankProduction", Operator: "=", Value: 0, Target: "Player"},
		},
		Selector:  "range",
		BaseValue: 100,
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "Infantry equipment tank production",
		ActionName:  "InfantryEquipmentProduction",
		Description: "Select the percentage of factories production which go toward Infantry equipmentProduction.",
		Cooldown:    0,
		Constraints: []Constraint{
			Constraint{Type: "custom", Value: "0", Operator: ">"},
			Constraint{Type: "custom", Value: "200", Operator: "<"},
			Constraint{Type: "linked", Value: "production"},
		},
		Effects: []Effect{
			Effect{ModifierType: "Economy", ModifierName: "InfantryEquipmentProduction", Operator: "=", Value: 0, Target: "Player"},
		},
		Selector:  "range",
		BaseValue: 100,
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "Artillery production",
		ActionName:  "ArtilleryProduction",
		Description: "Select the percentage of factories production which go toward Artillery.",
		Cooldown:    0,
		Constraints: []Constraint{
			Constraint{Type: "custom", Value: "0", Operator: ">"},
			Constraint{Type: "custom", Value: "200", Operator: "<"},
			Constraint{Type: "linked", Value: "production"},
		},
		Effects: []Effect{
			Effect{ModifierType: "Economy", ModifierName: "ArtilleryProduction", Operator: "=", Value: 0, Target: "Player"},
		},
		Selector:  "range",
		BaseValue: 100,
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "Fighter production",
		ActionName:  "FighterProduction",
		Description: "Select the percentage of factories production which go toward figther.",
		Cooldown:    0,
		Constraints: []Constraint{
			Constraint{Type: "custom", Value: "0", Operator: ">"},
			Constraint{Type: "custom", Value: "200", Operator: "<"},
			Constraint{Type: "linked", Value: "production"},
		},
		Effects: []Effect{
			Effect{ModifierType: "Economy", ModifierName: "FighterProduction", Operator: "=", Value: 0, Target: "Player"},
		},
		Selector:  "range",
		BaseValue: 100,
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "Bomber production",
		ActionName:  "BomberProduction",
		Description: "Select the percentage of factories production which go toward Bomber.",
		Cooldown:    0,
		Constraints: []Constraint{
			Constraint{Type: "custom", Value: "0", Operator: ">"},
			Constraint{Type: "custom", Value: "200", Operator: "<"},
			Constraint{Type: "linked", Value: "production"},
		},
		Effects: []Effect{
			Effect{ModifierType: "Economy", ModifierName: "BomberProduction", Operator: "=", Value: 0, Target: "Player"},
		},
		Selector:  "range",
		BaseValue: 100,
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "Factory production",
		ActionName:  "FactoryProduction",
		Description: "Select the percentage of factories production which go toward Factory.",
		Cooldown:    0,
		Constraints: []Constraint{
			Constraint{Type: "custom", Value: "0", Operator: ">"},
			Constraint{Type: "custom", Value: "200", Operator: "<"},
			Constraint{Type: "linked", Value: "production"},
		},
		Effects: []Effect{
			Effect{ModifierType: "Economy", ModifierName: "FactoryProduction", Operator: "=", Value: 0, Target: "Player"},
		},
		Selector:  "range",
		BaseValue: 100,
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "INDUSTRIAL",
		Name:        "ammunition production",
		ActionName:  "AmmunitionProduction",
		Description: "Select the percentage of factories production which go toward ammunition.",
		Cooldown:    0,
		Constraints: []Constraint{
			Constraint{Type: "custom", Value: "0", Operator: ">"},
			Constraint{Type: "custom", Value: "200", Operator: "<"},
			Constraint{Type: "linked", Value: "production"},
		},
		Effects: []Effect{
			Effect{ModifierType: "Economy", ModifierName: "AmmunitionProduction", Operator: "=", Value: 0, Target: "Player"},
		},
		Selector:  "range",
		BaseValue: 100,
	})

	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "TAX",
		Name:        "Take Loan",
		ActionName:  "TakeLoan",
		Description: "Take a loan. You gain immediatly 50M, but you will pay 50M in 20 turns and 1M each turns.",
		Cooldown:    20,
		Costs:       []Cost{Cost{Type: "money", Value: 0}},
		Effects: []Effect{
			Effect{
				ModifierType: "Modifier",
				ModifierName: "loanT1",
				Operator:     "turn+",
				Value:        20,
				Target:       "Player",
				Callbacks: []CallbackEffect{
					CallbackEffect{
						Constraints: []Constraint{Constraint{Type: "ModifierTurn", Operator: ">", Key: "loanT1"}},
						Effects: []Effect{
							Effect{ModifierType: "Economy", ModifierName: "Money", Operator: "-", Value: 50000000, Target: "Player"},
							Effect{ModifierType: "Economy", ModifierName: "Loans", Operator: "-", Value: 1, Target: "Player"},
						},
					},
				},
			},
			Effect{ModifierType: "Economy", ModifierName: "Money", Operator: "+", Value: 50000000, Target: "Player"},
			Effect{ModifierType: "Economy", ModifierName: "Loans", Operator: "+", Value: 1, Target: "Player"},
		},
		Selector: "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "TAX",
		Name:        "Take large Loan",
		ActionName:  "TakeLoanT2",
		Description: "Take a loan. You gain immediatly 150M, but you will pay 150M in 30 turns and 2M each turns.",
		Cooldown:    20,
		Costs:       []Cost{Cost{Type: "money", Value: 0}},
		Effects: []Effect{
			Effect{
				ModifierType: "Modifier",
				ModifierName: "loanT2",
				Operator:     "turn+",
				Value:        30,
				Target:       "Player",
				Callbacks: []CallbackEffect{
					CallbackEffect{
						Constraints: []Constraint{Constraint{Type: "ModifierTurn", Operator: ">", Key: "loanT1"}},
						Effects: []Effect{
							Effect{ModifierType: "Economy", ModifierName: "Money", Operator: "-", Value: 150000000, Target: "Player"},
							Effect{ModifierType: "Economy", ModifierName: "Loans", Operator: "-", Value: 2, Target: "Player"},
						},
					},
				},
			},
			Effect{ModifierType: "Economy", ModifierName: "Money", Operator: "+", Value: 150000000, Target: "Player"},
			Effect{ModifierType: "Economy", ModifierName: "Loans", Operator: "+", Value: 2, Target: "Player"},
		},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N3"},
		},
		Selector: "fixed",
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMY",
		SubType:     "TAX",
		Name:        "Take massive Loan",
		ActionName:  "TakeLoanT3",
		Description: "Take a loan. You gain immediatly 500M, but you will pay 500M in 40 turns and 5M each turns.",
		Cooldown:    20,
		Costs:       []Cost{Cost{Type: "money", Value: 0}},
		Effects: []Effect{
			Effect{
				ModifierType: "Modifier",
				ModifierName: "loanT3",
				Operator:     "turn+",
				Value:        40,
				Target:       "Player",
				Callbacks: []CallbackEffect{
					CallbackEffect{
						Constraints: []Constraint{Constraint{Type: "ModifierTurn", Operator: ">", Key: "loanT1"}},
						Effects: []Effect{
							Effect{ModifierType: "Economy", ModifierName: "Money", Operator: "-", Value: 500000000, Target: "Player"},
							Effect{ModifierType: "Economy", ModifierName: "Loans", Operator: "-", Value: 5, Target: "Player"},
						},
					},
				},
			},
			Effect{ModifierType: "Economy", ModifierName: "Money", Operator: "+", Value: 500000000, Target: "Player"},
			Effect{ModifierType: "Economy", ModifierName: "Loans", Operator: "+", Value: 5, Target: "Player"},
		},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT2N3"},
		},
		Selector: "fixed",
	})

}
