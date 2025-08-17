package cinemacity

const baseUrl = "https://www.cinemacity.hu/hu/data-api-service/v1/quickbook/10102/"

const (
	ARENA     = 1132
	ALLEE     = 1133
	WESTEND   = 1137
	CAMPONA   = 1139
	DUNAPLAZA = 1141
	MAMMUT    = 1144
)

var locationNames = map[int]string{
	ALLEE:     "Allee",
	ARENA:     "Aréna",
	CAMPONA:   "Campona",
	DUNAPLAZA: "Duna Pláza",
	MAMMUT:    "Mammut",
	WESTEND:   "WestEnd",
}
