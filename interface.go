package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Email struct {
	Address string
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправленно по почте на адрес %v \n", msg, e.Address)
	return nil
}

type Phone struct {
	Number int
	Balance int
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправленно по телефону с номера %v \n", msg, p.Number)
	return nil
}


func Notify(s Sender) {
	err := s.Send("Notify message")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	email := &Email{"test@test.com"}
	Notify(email)

	phone := &Phone{Number: 7777, Balance: 10}
	Notify(phone)
}