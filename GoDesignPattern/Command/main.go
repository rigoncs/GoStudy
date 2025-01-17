package command

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		decice: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
