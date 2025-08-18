package pannonia

const (
	mom    = 0
	polus  = 1
	lurdy  = 2
	gobuda = 3
)

var baseUrl = map[int]string{
	mom:    "https://cinemamom.hu/",
	polus:  "https://poluscentermozi.hu/",
	lurdy:  "https://lurdymozi.hu/",
	gobuda: "https://gobudamozi.hu/",
}

var locationName = map[int]string{
	mom:    "Cinema MOM",
	polus:  "Pólus",
	lurdy:  "Lurdy",
	gobuda: "GOBUDA",
}

var titlePostfixes = []string{
	" - Original language with Hungarian subtitles",
	" - A film",
	" - Digitálisan felújított",
	" 4K",
}

var excludedKeywords = []string{
	"bábszínház",
	"bábmese",
	"Royal Ballet and Opera",
}
