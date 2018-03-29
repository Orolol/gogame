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
		Type:        "ECONOMIC",
		Name:        "Civ Fact -> Lght Fact",
		ActionName:  "actionCivConvertFactoryToLightTankFact",
		Description: "Convert Civilian Factory to light Tank factory (Cost 1M) ",
		Cooldown:    10,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbCivilianFactory", Operator: "-", Value: 1, Target: "Player"},
			Effect{ModifierType: "Civilian", ModifierName: "NbLightFactory", Operator: "+", Value: 1, Target: "Player"},
		},
		Costs: []Cost{Cost{Type: "money", Value: 1000000}},
	})

	actions = append(actions, PlayerActionOrder{
		Type:        "ECONOMIC",
		Name:        "Civ Fact -> Hvy Fact",
		ActionName:  "actionCivConvertFactoryToHvyTankFact",
		Description: "Convert Civilian Factory to Heavy Tank factory (Cost 1M) ",
		Cooldown:    10,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbCivilianFactory", Operator: "-", Value: 1, Target: "Player"},
			Effect{ModifierType: "Civilian", ModifierName: "NbHeavyFactory", Operator: "+", Value: 1, Target: "Player"},
		},
		Costs: []Cost{Cost{Type: "money", Value: 1000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "MILITARY",
		Name:        "War Propaganda",
		ActionName:  "actionWarPropaganda",
		Description: "Boost morale by 20% (cost 10M) ",
		Cooldown:    10,
		Effects:     []Effect{Effect{ModifierType: "Army", ModifierName: "Morale", Operator: "+", Value: 20, Target: "Player"}},
		Costs:       []Cost{Cost{Type: "money", Value: 10000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "MILITARY",
		Name:        "Emergency recruitment",
		ActionName:  "emergencyRecruitment",
		Description: "Recruit immediatly 15 000 soldier. (cost 50M and 10% morale)",
		Cooldown:    30,
		Effects:     []Effect{Effect{ModifierType: "Army", ModifierName: "NbSoldier", Operator: "+", Value: 15000, Target: "Player"}},
		Costs:       []Cost{Cost{Type: "money", Value: 50000000}, Cost{Type: "morale", Value: 10}},
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "MILITARY",
		Name:        "Purge of weak elements",
		ActionName:  "purgeSoldier",
		Description: "Executes possible traitors and weak soldier to improve morale and discipline by 15%. (cost 15% soldier and 10M)",
		Cooldown:    30,
		Effects: []Effect{
			Effect{ModifierType: "Army", ModifierName: "NbSoldier", Operator: "*", Value: 0.85, Target: "Player"},
			Effect{ModifierType: "Army", ModifierName: "Quality", Operator: "+", Value: 0.15, Target: "Player"},
		},
		Costs: []Cost{Cost{Type: "money", Value: 10000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "SABOTAGE",
		Name:        "Send spys to assassinate key scientists",
		ActionName:  "assassinateScientist",
		Description: "Send spys to assassinate key scientists (cost 45M)",
		Cooldown:    30,
		Effects:     []Effect{Effect{ModifierType: "Civilian", ModifierName: "NbScientist", Operator: "-", Value: 30, Target: "Opponent"}},
		Costs:       []Cost{Cost{Type: "money", Value: 45000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "SABOTAGE",
		Name:        "Send spys to assassinate key scientists and destroy laboratories",
		ActionName:  "advancedAssassinateScientist",
		Description: "Send spys to assassinate key scientists (cost 75M)",
		Cooldown:    30,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbScientist", Operator: "-", Value: 50, Target: "Opponent"},
			Effect{ModifierType: "Modifiers", ModifierName: "labDestroyed", Operator: "+", Value: 10, Target: "Opponent"},
		},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N2"},
		},
		Costs: []Cost{Cost{Type: "money", Value: 75000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "SABOTAGE",
		Name:        "Sabotage factories",
		ActionName:  "sabotageFactories",
		Description: "Pay foreign workers to sabotage factories(cost 60M)",
		Cooldown:    30,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbCivilianFactory", Operator: "-", Value: 2, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbLightFactory", Operator: "-", Value: 1, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbHeavyFactory", Operator: "-", Value: 1, Target: "Opponent"},
		},
		Costs: []Cost{Cost{Type: "money", Value: 60000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "SABOTAGE",
		Name:        "Worker Uprising",
		ActionName:  "sabotageFactoriesAdvanced",
		Description: "Pay foreign workers to sabotage factories and provoke country wise worker uprising(cost 90M)",
		Cooldown:    30,
		Effects: []Effect{
			Effect{ModifierType: "Civilian", ModifierName: "NbCivilianFactory", Operator: "-", Value: 5, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbLightFactory", Operator: "-", Value: 2, Target: "Opponent"},
			Effect{ModifierType: "Civilian", ModifierName: "NbHeavyFactory", Operator: "-", Value: 2, Target: "Opponent"},
		},
		Constraints: []Constraint{
			Constraint{Type: "tech", Value: "technoEcoT1N2"},
		},
		Costs: []Cost{Cost{Type: "money", Value: 90000000}},
	})
	actions = append(actions, PlayerActionOrder{
		Type:        "TEST",
		Name:        "CBTEST",
		ActionName:  "CBTEST",
		Description: "CBTEST",
		Cooldown:    3,
		Costs:       []Cost{Cost{Type: "money", Value: 0}},
		Effects: []Effect{
			Effect{
				ModifierType: "Modifier",
				ModifierName: "TEST",
				Operator:     "turn+",
				Value:        3,
				Target:       "Player",
				Callbacks: []CallbackEffect{
					CallbackEffect{
						Constraints: []Constraint{Constraint{Type: "ModifierTurn", Operator: ">", Key: "TEST"}},
						Effects:     []Effect{Effect{ModifierType: "Army", ModifierName: "Morale", Operator: "+", Value: 30, Target: "Player"}},
					},
				},
			},
			Effect{ModifierType: "Army", ModifierName: "Morale", Operator: "-", Value: 30, Target: "Player"},
		},
	})
	actions = append(actions, PlayerActionOrder{
		Type:       "MILITARY",
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

}
