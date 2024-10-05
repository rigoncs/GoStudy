package factorymethod

import "fmt"

func getGun(gunType string) (iGun, error) {
	if gunType == "awm" {
		return newAwm(), nil
	} else if gunType == "ak47" {
		return newAk47(), nil
	} else {
		return nil, fmt.Errorf("Wrong gun type passed")
	}
}
