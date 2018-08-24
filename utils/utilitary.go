package utils

// func updateOrCreateInfos(name string, player *PlayerInGame, value float32, v string, subtypeInfo string){
// 	var found = false
// 	if val, ok := player.PlayerInformations[name]; ok {
// 		player.PlayerInformations[name].Value = value
// 	} else {
// 		player.PlayerInformations[name] = PlayerInformation{
// 			Name: name,
// 			Type: typeInfo,
// 			SubType: subtypeInfo,
// 			Value: value,

// 		}
// 	}
//     //do something here
// }
// }

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
