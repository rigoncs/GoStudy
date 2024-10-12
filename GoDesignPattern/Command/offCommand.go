package command

type OffCommand struct {
	decice Device
}

func (c *OffCommand) execute() {
	c.decice.off()
}
