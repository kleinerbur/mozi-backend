package artmozi

const (
	corvin = 1447
	puskin = 1448
	toldi  = 1449
	muvesz = 1450
	taban  = 1451
	kino   = 1452
)

var locationName = map[int]string{
	corvin: "Corvin",
	puskin: "Puskin",
	toldi:  "Toldi",
	muvesz: "Művész",
	taban:  "Tabán",
	kino:   "Kino Café",
}

var baseUrl = map[int]string{
	corvin: "https://corvinmozi.hu/api/",
	puskin: "https://puskinmozi.hu/api/",
	toldi:  "https://toldimozi.hu/api/",
	muvesz: "https://muveszmozi.hu/api/",
	taban:  "https://tabanartmozi.hu/api/",
	kino:   "https://kinocafemozi.hu/api/",
}
