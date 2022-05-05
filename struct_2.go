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

func main() {
	i := new(int64)
	_ = i

	client := Client{}
	fmt.Printf("%#v\n", client)
	// main.Client{ID:0, Name:"", Age:0, IMG:main.Avatar{URL:"", Size:0}}

	// client.IMG = new(Avatar)
	updateAvatar(&client)
	fmt.Printf("%#v\n", client)
	updateClient(client)
}

func updateAvatar(client *Client) {
	client.IMG.URL = "update_url"
}

func updateClient(client Client) {
	client.Name = "Артём"
}
