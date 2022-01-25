package organization

import (
	"errors"
	"fmt"
	"strconv"
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

type Citizen interface {
	Identifiable
	Country() string
}

type SocialSecurityNumber string

// Return the interface
func NewSocialSecurityNumber(value string) Identifiable {
	return SocialSecurityNumber(value)
}

func (ssn SocialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn SocialSecurityNumber) Country() string {
	return "United States of America"
}

type EuropeanUnionIdentifier struct {
	id      string
	country []string
}

// Return the interface
func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return EuropeanUnionIdentifier{
			id:      v,
			country: []string{country},
		}
	case int:
		return EuropeanUnionIdentifier{
			id:      strconv.Itoa(v),
			country: []string{country},
		}
	case EuropeanUnionIdentifier:
		return v
	case Person:
		euID, _ := v.Citizen.(EuropeanUnionIdentifier)
		return euID
	default:
		panic("using an invalid type ot initalize EU Identifier")
	}
}

func (eui EuropeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui EuropeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
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
	Citizen
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifiable: %s", p.Citizen.ID())
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
