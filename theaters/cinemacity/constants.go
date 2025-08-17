package cinemacity

const baseUrl = "https://www.cinemacity.hu/hu/data-api-service/v1/quickbook/10102/"

const (
	arena     = 1132
	allee     = 1133
	westend   = 1137
	campona   = 1139
	dunaplaza = 1141
	mammut    = 1144
)

var locationName = map[int]string{
	arena:     "Aréna",
	allee:     "Allee",
	westend:   "WestEnd",
	campona:   "Campona",
	dunaplaza: "Duna Pláza",
	mammut:    "Mammut",
}
