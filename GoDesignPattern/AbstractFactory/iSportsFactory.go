package abstractfactory

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	} else if brand == "nike" {
		return &Nike{}, nil
	} else {
		return nil, fmt.Errorf("Wrong brand type passed")
	}
}
