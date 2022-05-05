package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Avatar struct {
	URL  string
	Size int64
}

type Client struct {
	Id   int64
	Img  Avatar
	Name string
	Person
}

func (c Client) HasAvatar() bool {
	if c.Img.URL != "" {
		return true
	}
	return false
}

func (c Client) UpdateAvatar() {
	c.Img.URL = "new_url"
}

func main() {
	client := Client{}

	fmt.Printf("%#v\n", client)
	client.UpdateAvatar()
	fmt.Printf("%#v\n", client)

	client.Name = "Client Name"
	client.Person.Name = "Person Name"
	fmt.Println(client.Name)
	fmt.Println(client.Person.Name)
}
