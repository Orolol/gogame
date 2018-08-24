package utils

var policies []Policy

func GetPolicies() []Policy {
	return policies
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

func SetBaseValuePolicies() {

	//POLICIES
	policies = append(policies, Policy{
		Name:        "Training Time",
		ActionName:  "setPopRecPolicy",
		Description: "Set your recuitement policy",
		TypePolicy:  "MILITARY",
		SubType:     "RECRUITMENT",
		MaxChange:   1,
		PossibleValue2: []PolicyValue{
			PolicyValue{Name: "Full", Value: 1, Description: "Take time to train well soldiers", ActionName: "setPopRecPolicy", Constraints: nil,
				Effects: []Effect{
					Effect{ModifierType: "Policy", Value: 1, Operator: "=", ModifierName: "TrainingPolicy"},
				}, IsDefault: true,
			},
			PolicyValue{Name: "Long", Value: 2, Description: "Make sure everyone know how to fight", ActionName: "setPopRecPolicy", Constraints: nil,
				Effects: []Effect{
					Effect{ModifierType: "Policy", Value: 1.5, Operator: "=", ModifierName: "TrainingPolicy"},
					Effect{ModifierType: "Modifiers", Value: 0.98, Operator: "=", ModifierName: "soldierRecruitmentExperience"},
				},
			},
			PolicyValue{Name: "Hurry", Value: 3, Description: "Army need fresh recruit !", ActionName: "setPopRecPolicy", Constraints: []Constraint{
				Constraint{Type: "isWar"},
			}, Effects: []Effect{
				Effect{ModifierType: "Policy", Value: 2.5, Operator: "=", ModifierName: "TrainingPolicy"},
				Effect{ModifierType: "Modifiers", Value: 0.95, Operator: "=", ModifierName: "soldierRecruitmentExperience"},
			},
			},
			PolicyValue{Name: "No time !", Value: 4, Description: "If they can handle a rifle, send them !", ActionName: "setPopRecPolicy", Constraints: []Constraint{
				Constraint{Type: "isWar"},
			}, Effects: []Effect{
				Effect{ModifierType: "Policy", Value: 5, Operator: "=", ModifierName: "TrainingPolicy"},
				Effect{ModifierType: "Modifiers", Value: 0.90, Operator: "=", ModifierName: "soldierRecruitmentExperience"},
			},
			},
			PolicyValue{Name: "Send everyone !", Value: 5, Description: "Drag the full country", ActionName: "setPopRecPolicy", Constraints: []Constraint{
				Constraint{Type: "isWar"},
			}, Effects: []Effect{
				Effect{ModifierType: "Policy", Value: 10, Operator: "=", ModifierName: "TrainingPolicy"},
				Effect{ModifierType: "Modifiers", Value: 0.80, Operator: "=", ModifierName: "soldierRecruitmentExperience"},
			},
			},
		}})
	policies = append(policies, Policy{
		Name:          "Conscription Policy",
		ActionName:    "setConscPolicy",
		Description:   "Set your conscription policy. The more your mobilize your population, the less the workers will be productive",
		TypePolicy:    "MILITARY",
		SubType:       "RECRUITMENT",
		MaxChange:     1,
		PossibleValue: "{\"Pro Army\" : 1,\"Volonteer\" : 2,\"War time\" : 5,\"All valids !\" : 10,\"Anyone who can hold a weapon\" : 30}",
		PossibleValue2: []PolicyValue{
			PolicyValue{Name: "Professionnal army", Value: 1, Description: "Take time to train well soldiers", ActionName: "setConscPolicy", Constraints: nil,
				Effects: []Effect{
					Effect{ModifierType: "Policy", Value: 0.005, Operator: "=", ModifierName: "ManpowerSizePolicy"},
				}, IsDefault: true,
			},
			PolicyValue{Name: "Volunteer", Value: 2, Description: "Make sure everyone know how to fight", ActionName: "setConscPolicy", Constraints: nil,
				Effects: []Effect{
					Effect{ModifierType: "Policy", Value: 0.01, Operator: "=", ModifierName: "ManpowerSizePolicy"},
					Effect{ModifierType: "Modifiers", Value: 0.98, Operator: "=", ModifierName: "workersConcrptionEfficiency"},
				},
			},
			PolicyValue{Name: "Conscription", Value: 3, Description: "Army need fresh recruit !", ActionName: "setConscPolicy", Constraints: []Constraint{
				Constraint{Type: "isWar"},
			}, Effects: []Effect{
				Effect{ModifierType: "Policy", Value: 0.02, Operator: "=", ModifierName: "ManpowerSizePolicy"},
				Effect{ModifierType: "Modifiers", Value: 0.95, Operator: "=", ModifierName: "workersConcrptionEfficiency"},
			},
			},
			PolicyValue{Name: "All men valid", Value: 4, Description: "If they can handle a rifle, send them !", ActionName: "setConscPolicy", Constraints: []Constraint{
				Constraint{Type: "isWar"},
			}, Effects: []Effect{
				Effect{ModifierType: "Policy", Value: 0.05, Operator: "=", ModifierName: "ManpowerSizePolicy"},
				Effect{ModifierType: "Modifiers", Value: 0.90, Operator: "=", ModifierName: "workersConcrptionEfficiency"},
			},
			},
			PolicyValue{Name: "All men and women", Value: 5, Description: "Drag the full country", ActionName: "setConscPolicy", Constraints: []Constraint{
				Constraint{Type: "isWar"},
			}, Effects: []Effect{
				Effect{ModifierType: "Policy", Value: 0.01, Operator: "=", ModifierName: "ManpowerSizePolicy"},
				Effect{ModifierType: "Modifiers", Value: 0.80, Operator: "=", ModifierName: "workersConcrptionEfficiency"},
			},
			},
		}})

	policies = append(policies, Policy{
		Name:          "Tax rate",
		ActionName:    "setTaxRatePolicy",
		Description:   "Set your tax rate. ",
		TypePolicy:    "ECONOMY",
		SubType:       "TAX",
		MaxChange:     1,
		PossibleValue: "{\"Low taxes\" : 1,\"Country effort\" : 1.5,\"War Economy\" : 2,\"Full Mobilization\" : 3,\"Total war\" : 5}",
		PossibleValue2: []PolicyValue{
			PolicyValue{Name: "Low taxes", Value: 1, Description: "Low taxes rates", Constraints: nil, ActionName: "setTaxRatePolicy",
				Effects: []Effect{
					Effect{ModifierType: "Economy", Value: 1, Operator: "=", ModifierName: "TaxRate"},
				}, IsDefault: true,
			},
			PolicyValue{Name: "Country effort", Value: 1.1, Description: "Raise taxes to prepare for war", Constraints: nil, ActionName: "setTaxRatePolicy",
				Effects: []Effect{
					Effect{ModifierType: "Economy", Value: 1.1, Operator: "=", ModifierName: "TaxRate"},
					Effect{ModifierType: "Modifiers", Value: 0.98, Operator: "=", ModifierName: "TaxEffectOnIndustry"},
				},
			},
			PolicyValue{Name: "War Economy", Value: 1.2, Description: "War is upon us, country must be ready", ActionName: "setTaxRatePolicy", Constraints: []Constraint{
				Constraint{Type: "isWar"},
			}, Effects: []Effect{
				Effect{ModifierType: "Economy", Value: 1.2, Operator: "=", ModifierName: "TaxRate"},
				Effect{ModifierType: "Modifiers", Value: 0.95, Operator: "=", ModifierName: "TaxEffectOnIndustry"},
			},
			},
			PolicyValue{Name: "Full Mobilization", Value: 1.3, Description: "Everyone should make effort", ActionName: "setTaxRatePolicy", Constraints: []Constraint{
				Constraint{Type: "isWar"},
			}, Effects: []Effect{
				Effect{ModifierType: "Economy", Value: 1.3, Operator: "=", ModifierName: "TaxRate"},
				Effect{ModifierType: "Modifiers", Value: 0.90, Operator: "=", ModifierName: "TaxEffectOnIndustry"},
			},
			},
			PolicyValue{Name: "Total war", Value: 1.5, Description: "Who need schools or food anyway ?", ActionName: "setTaxRatePolicy", Constraints: []Constraint{
				Constraint{Type: "isWar"},
			}, Effects: []Effect{
				Effect{ModifierType: "Economy", Value: 1.5, Operator: "=", ModifierName: "TaxRate"},
				Effect{ModifierType: "Modifiers", Value: 0.80, Operator: "=", ModifierName: "TaxEffectOnIndustry"},
			},
			},
		}})

	policies = append(policies, Policy{
		Name:        "Bomber combat style",
		ActionName:  "bomberCombatStyle",
		Description: "Bomber combat style",
		TypePolicy:  "MILITARY",
		SubType:     "COMMANDMENT",
		MaxChange:   10,
		PossibleValue2: []PolicyValue{
			PolicyValue{Name: "Target army", Value: 1, Description: "Bomber will focus the ground support", Constraints: nil, ActionName: "bomberCombatStyle",
				Effects: []Effect{
					Effect{ModifierType: "Modifiers", Value: 1, Operator: "=", ModifierName: "bomberTargetArmy"},
					Effect{ModifierType: "Modifiers", Value: 0, Operator: "=", ModifierName: "bomberTargetFactory"},
					Effect{ModifierType: "Modifiers", Value: 0, Operator: "=", ModifierName: "bomberTargetPopulation"},
				}, IsDefault: true,
			},
			PolicyValue{Name: "Target factories", Value: 0, Description: "Bomber will focus the enemy production", Constraints: nil, ActionName: "bomberCombatStyle",
				Effects: []Effect{
					Effect{ModifierType: "Modifiers", Value: 0, Operator: "=", ModifierName: "bomberTargetArmy"},
					Effect{ModifierType: "Modifiers", Value: 1, Operator: "=", ModifierName: "bomberTargetFactory"},
					Effect{ModifierType: "Modifiers", Value: 0, Operator: "=", ModifierName: "bomberTargetPopulation"},
				}, IsDefault: false,
			},
		}})

	policies = append(policies, Policy{
		Name:        "Engage aerial force",
		ActionName:  "engageAerialForce",
		Description: "Engage aerial force",
		TypePolicy:  "MILITARY",
		SubType:     "COMMANDMENT",
		MaxChange:   10,
		PossibleValue2: []PolicyValue{
			PolicyValue{Name: "Send every planes on fight !", Value: 2, Description: "Send every planes on fight !", Constraints: nil, ActionName: "engageAerialForce",
				Effects: []Effect{
					Effect{ModifierType: "Modifiers", Value: 2, Operator: "=", ModifierName: "engageAerialForce"},
					Effect{ModifierType: "Modifiers", Value: 1, Operator: "=", ModifierName: "engageFighter"},
					Effect{ModifierType: "Modifiers", Value: 1, Operator: "=", ModifierName: "engageBomber"},
				}, IsDefault: false,
			},
			PolicyValue{Name: "Send the fighter !", Value: 1, Description: "Send the fighter !", Constraints: nil, ActionName: "engageAerialForce",
				Effects: []Effect{
					Effect{ModifierType: "Modifiers", Value: 1, Operator: "=", ModifierName: "engageAerialForce"},
					Effect{ModifierType: "Modifiers", Value: 1, Operator: "=", ModifierName: "engageFighter"},
					Effect{ModifierType: "Modifiers", Value: 0, Operator: "=", ModifierName: "engageBomber"},
				}, IsDefault: true,
			},
			PolicyValue{Name: "All planes at base", Value: 0, Description: "Bomber will focus the enemy production", Constraints: nil, ActionName: "engageAerialForce",
				Effects: []Effect{
					Effect{ModifierType: "Modifiers", Value: 0, Operator: "=", ModifierName: "engageAerialForce"},
					Effect{ModifierType: "Modifiers", Value: 0, Operator: "=", ModifierName: "engageFighter"},
					Effect{ModifierType: "Modifiers", Value: 0, Operator: "=", ModifierName: "engageBomber"},
				}, IsDefault: true,
			},
		}})

}
