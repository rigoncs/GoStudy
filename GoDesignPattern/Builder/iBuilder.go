package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "igloo" {
		return newIglooBuilder()
	} else if builderType == "normal" {
		return newNormalBuilder()
	}
	return nil
}
