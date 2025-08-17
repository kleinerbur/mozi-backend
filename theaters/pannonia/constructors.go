package pannonia

func new(id int) (*Pannonia, error) {
	p := &Pannonia{
		id,
		baseUrl[id],
		locationName[id],
	}
	err := p.init()
	return p, err
}

func Mom() (*Pannonia, error) {
	return new(mom)
}

func Polus() (*Pannonia, error) {
	return new(polus)
}

func Lurdy() (*Pannonia, error) {
	return new(lurdy)
}

func GoBuda() (*Pannonia, error) {
	return new(gobuda)
}
