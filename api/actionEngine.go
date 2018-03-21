package main

import (
	"github.com/orolol/gogame/utils"
)

func getDefaultPolicies() []utils.Policy {
	return utils.GetPolicies()

}

func getDefaultActions() []utils.PlayerActionOrder {
	return utils.GetActions()
}

func getDefaultTech() []utils.Technology {
	return utils.GetTechnolgies()
}

func getDefaultEvents() *[]utils.PlayerEvent {
	return utils.GetEvents()
}
