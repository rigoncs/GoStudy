package abstractfactory

type IShirt interface {
	setLogo(logo string)
	getLogo() string
}

type Shirt struct {
	logo string
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) getLogo() string {
	return s.logo
}
