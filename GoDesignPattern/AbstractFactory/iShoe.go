package abstractfactory

type IShoe interface {
	setLogo(logo string)
	getLogo() string
}

type Shoe struct {
	logo string
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) getLogo() string {
	return s.logo
}
