package adapter

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertIntoLightningPort(mac)

	windowMachine := &Windows{}
	windowMachineAdapter := &WindowsAdapter{
		windowMachine: windowMachine,
	}
	client.InsertIntoLightningPort(windowMachineAdapter)
}
