package artmozi

const (
	CORVIN = 1447
	PUSKIN = 1448
	TOLDI  = 1449
	MUVESZ = 1450
	TABAN  = 1451
	KINO   = 1452
)

var locationNames = map[int]string{
	CORVIN: "Corvin",
	PUSKIN: "Puskin",
	TOLDI:  "Toldi",
	MUVESZ: "Művész",
	TABAN:  "Tabán",
	KINO:   "Kino Café",
}

var baseUrls = map[int]string{
	CORVIN: "https://corvinmozi.hu/api/",
	PUSKIN: "https://puskinmozi.hu/api/",
	TOLDI:  "https://toldimozi.hu/api/",
	MUVESZ: "https://muveszmozi.hu/api/",
	TABAN:  "https://tabanartmozi.hu/api/",
	KINO:   "https://kinocafemozi.hu/api/",
}
