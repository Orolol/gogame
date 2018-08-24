package utils

var news []News

func GetNews() []News {
	return news
}

func SetBaseValueNews() {

	news = append(news, News{
		Date:      "23 Aug 2018",
		Title:     "Alpha Launch",
		Paragrahs: []string{"Get ready for the **Alpha** (v0.1) launch of this fucking game !"},
	})

}
