package main

import "fmt"

type Avatar struct {
	URL  string
	Size int64
}

type Client struct {
	ID   int64
	Name string
	Age  int
	IMG  Avatar
}

func (c *Client) updateAvatar() {
	c.IMG.URL = "new_url"
}

func (c Client) hasAvatar() bool {
	if c.IMG.URL != "" {
		return true
	}
	return false
}

func main() {
	client := Client{
		ID:   7,
		Name: "Андрей",
		Age:  20,
		//IMG:  new(Avatar),
		IMG: Avatar{
			URL:  "some_url",
			Size: 10,
		},
	}
	fmt.Println(client.hasAvatar())
	fmt.Println(client)
	client.updateAvatar()
	fmt.Printf("%#v", client)
}
