package utils

var infosCat []DisplayInfoCat

func GetInfos() *[]DisplayInfoCat {
	return &infosCat
}

func SetBaseValueInfos() {

	//EVENTS

	infosCat = append(infosCat, DisplayInfoCat{
		Category: "generalInfos",
		Infos: []DisplayInfos{
			DisplayInfos{
				Name:         "SmallCities",
				Type:         "Territory",
				LowAlert:     50,
				VeryLowAlert: 15,
			},
			DisplayInfos{
				Name: "MediumCities",
				Type: "Territory",
			},
			DisplayInfos{
				Name: "BigCities",
				Type: "Territory",
			},
			DisplayInfos{
				Name: "Barracks",
				Type: "Territory",
			},
			DisplayInfos{
				Name: "",
				Type: "Separator",
			},
			DisplayInfos{
				Name:         "Money",
				Type:         "Economy",
				LowAlert:     10000000,
				VeryLowAlert: 1000000,
			},
			DisplayInfos{
				Name:       "Factory",
				Type:       "PlayerInformations",
				GrowthName: "factoryProduction",
				GrowthType: "PlayerInformations",
			},
			DisplayInfos{
				Name: "NbResearchPoint",
				Type: "Civilian",
			},
			DisplayInfos{
				Name: "NbScientist",
				Type: "Civilian",
			},
			DisplayInfos{
				Name: "",
				Type: "Separator",
			},
			DisplayInfos{
				Name:       "NbSoldier",
				Type:       "Army",
				GrowthName: "SoldierRecruit",
				GrowthType: "PlayerInformations",
			},
			DisplayInfos{
				Name:       "NbLigtTank",
				Type:       "Army",
				GrowthName: "lightTankProduction",
				GrowthType: "PlayerInformations",
			},
			DisplayInfos{
				Name:       "NbHvyTank",
				Type:       "Army",
				GrowthName: "heavyTankProduction",
				GrowthType: "PlayerInformations",
			},
			DisplayInfos{
				Name:       "NbArt",
				Type:       "Army",
				GrowthName: "artilleryProduction",
				GrowthType: "PlayerInformations",
			},
			DisplayInfos{
				Name: "",
				Type: "Separator",
			},
			DisplayInfos{
				Name:       "NbAirSup",
				Type:       "Army",
				GrowthName: "fighterProduction",
				GrowthType: "PlayerInformations",
			},
			DisplayInfos{
				Name:       "NbAirBomb",
				Type:       "Army",
				GrowthName: "bomberProduction",
				GrowthType: "PlayerInformations",
			},
			DisplayInfos{
				Name: "",
				Type: "Separator",
			},
			DisplayInfos{
				Name: "Morale",
				Type: "Army",
			},
			DisplayInfos{
				Name: "Quality",
				Type: "Army",
			},
			DisplayInfos{
				Name: "",
				Type: "Separator",
			},
			DisplayInfos{
				Name:       "NbManpower",
				Type:       "Civilian",
				GrowthName: "ManpowerGrowth",
				GrowthType: "PlayerInformations",
			},
			DisplayInfos{
				Name:       "Ammunition",
				Type:       "Army",
				GrowthName: "ammunitionProduction",
				GrowthType: "PlayerInformations",
			},
			DisplayInfos{
				Name:       "InfantryEquipment",
				Type:       "Army",
				GrowthName: "infantryEquipmentProduction",
				GrowthType: "PlayerInformations",
			},
		},
	})

}
