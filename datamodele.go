package main

type Items struct {
	Items []Country `json:"items"`
}

//City用構造体
type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	District   string `json:"district"`
	Population int    `json:"population"`
}

type Country struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Continent      string  `json:"continent"`
	Region         string  `json:"region"`
	SurfaceArea    float64 `json:"surfaceArea"`
	IndepYear      int     `json:"indepYear"`
	Population     int     `jason:"population"`
	LifeExpectancy float64 `json:"lifeExpectancy"`
	GNP            float64 `json:"gnp"`
	GNPOld         float64 `json:"gnpOld"`
	LocalName      string  `json:localName`
	GovernmentForm string  `json:"governmentForm"`
	HeadOfState    string  `json:"headOfState"`
	Capital        int     `json:"capital"`
	Code2          string  `json:"code2"`
	Id             int     `json:"id"`
	CityName       string  `json:"cityName"`
	District       string  `json:"district"`
	CityPopulation int     `json:"cityPopulation"`
	Language       string  `json:"language"`
	IsOfficial     string  `json:"isOfficial"`
	Percentage     float64 `json:"percentage"`
}

type Language struct {
	CountryCode string  `json:"countryCode"`
	Language    string  `json:"language"`
	IsOfficial  string  `json:"isOfficial"`
	Percentage  float64 `json:"percentage"`
}
