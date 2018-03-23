package utils

var events []PlayerEvent

func GetEvents() *[]PlayerEvent {
	return &events
}

func GetEvent(name string) PlayerEvent {
	var ret PlayerEvent
	for _, x := range events {
		if x.ActionName == name {
			ret = x
		}
	}
	return ret
}
func GetEventsByType(typeEvent string) []PlayerEvent {

	var ret []PlayerEvent
	for _, x := range events {
		if x.Type == typeEvent {
			ret = append(ret, x)
		}
	}
	return ret
}

func SetBaseValueEvents() {

	//EVENTS

	events = append(events, PlayerEvent{
		Name:        "",
		Description: "",
		ActionName:  "event0",
		Weight:      0,
		Type:        "NullEvent",
	})
	events = append(events, PlayerEvent{
		Name:        "Famous Colonel die",
		Description: "One your most famous Colonel just died while defending position",
		Effects:     []Effect{Effect{ModifierType: "Army", ModifierName: "Morale", Operator: "*", Value: 0.95}},
		ActionName:  "event1",
		Weight:      5,
		Type:        "Single",
	})
	events = append(events, PlayerEvent{
		Name:        "A new hero",
		Description: "A soldier just accomplish an incredible feat of heroism.",
		Effects:     []Effect{Effect{ModifierType: "Army", ModifierName: "Morale", Operator: "*", Value: 1.05}},
		ActionName:  "event2",
		Weight:      5,
		Type:        "Single",
	})
	events = append(events, PlayerEvent{
		Name:        "Eureka !",
		Description: "Your scientists just made a fantastic breakthrough !",
		Effects:     []Effect{Effect{ModifierName: "NbResearchPoint", Operator: "+", Value: 250, ModifierType: "Civilian"}},
		ActionName:  "event3",
		Weight:      5,
		Type:        "Single",
	})
	events = append(events, PlayerEvent{
		Name:        "War enthusiasm",
		Description: "People rush your conscription centers ! You have a surge in fresh soldier and a morale boost",
		Effects: []Effect{
			Effect{ModifierName: "Morale", Operator: "+", Value: 0.05, ModifierType: "Army"},
			Effect{ModifierName: "NbSoldier", Operator: "+", Value: 5000, ModifierType: "Army"},
		},
		ActionName: "event4",
		Weight:     5,
		Type:       "Single",
	})

}
