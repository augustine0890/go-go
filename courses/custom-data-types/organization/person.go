package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Handler struct {
	handle string
	name   string
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type SocialSecurityNumber string

// Return the interface
func NewSocialSecurityNumber(value string) Identifiable {
	return SocialSecurityNumber(value)
}

func (ssn SocialSecurityNumber) ID() string {
	return string(ssn)
}

type Name struct {
	first string
	last  string
}
type Employee struct {
	name Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Identifiable
}

func NewPerson(firstName, lastName string, identifiable Identifiable) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Identifiable: identifiable,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifiable: %s0", p.Identifiable.ID())
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
