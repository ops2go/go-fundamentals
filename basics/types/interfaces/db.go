package interfaces

type Contact struct {
	name  string
	phone string
	email string
}

type Contacts []Contact

var phonebook []string

func (c Contacts) printemail() string {
	for _, Contact := range Contacts {
		phonebook == Contact
	}
	return phonebook
}

type search interface {
}
