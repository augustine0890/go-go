package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Name struct {
	first string
	last  string
}
type Employee struct {
	name Name
}

type Person struct {
	ID string
	Name
	twitterHandler string
}

func NewPerson(firstName, lastName string) Person {
	return Person{
		ID: "1123",
		Name: Name{
			first: firstName,
			last:  lastName,
		},
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

func (p *Person) SetTwitterHandler(handler string) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(handler, "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() string {
	return p.twitterHandler
}
