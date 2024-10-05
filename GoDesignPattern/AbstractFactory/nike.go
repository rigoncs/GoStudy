package abstractfactory

type Nike struct{}

func (n *Nike) makeShoe() IShoe {
	return &NikeShoe{
		Shoe{logo: "nike"},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt{logo: "nike"},
	}
}
