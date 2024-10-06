package adapter

import "fmt"

type Client struct{}

func (c *Client) InsertIntoLightningPort(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
