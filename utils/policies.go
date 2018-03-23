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

}
