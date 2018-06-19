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
				Name:         "Surface",
				Type:         "Territory",
				LowAlert:     50,
				VeryLowAlert: 15,
			},
			DisplayInfos{
				Name: "SmallCities",
				Type: "Territory",
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
				Name: "Factory",
				Type: "PlayerInformations",
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
				Name: "NbSoldier",
				Type: "Army",
			},
			DisplayInfos{
				Name: "NbLigtTank",
				Type: "Army",
			},
			DisplayInfos{
				Name: "NbHvyTank",
				Type: "Army",
			},
			DisplayInfos{
				Name: "NbArt",
				Type: "Army",
			},
			DisplayInfos{
				Name: "",
				Type: "Separator",
			},
			DisplayInfos{
				Name: "NbAirSup",
				Type: "Army",
			},
			DisplayInfos{
				Name: "NbAirBomb",
				Type: "Army",
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
				Name: "Ammunition",
				Type: "Army",
			},
			DisplayInfos{
				Name: "InfantryEquipment",
				Type: "Army",
			},
		},
	})

}
