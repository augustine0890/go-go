package models

import "temp-service/data"

var (
	beachVacationThreshold float64 = 22
	skipVacationThreshold  float64 = -2
)

type city struct {
	name        string
	tempC       float64
	hasBeach    bool
	hasMountain bool
}

// CityTemp interface defines all things city temperature related
type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64
	BeachVacationReady() bool
	SkiVacationReady() bool
}

// NewCity creates a new city instance with the given name and Celsius temperature.
func NewCity(name string, tempC float64, hasBeach bool, hasMountain bool) CityTemp {
	return &city{
		name:        name,
		tempC:       tempC,
		hasBeach:    hasBeach,
		hasMountain: hasMountain,
	}
}

func (c city) Name() string {
	return c.name
}

func (c city) TempC() float64 {
	return c.tempC
}

func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

func (c city) BeachVacationReady() bool {
	return c.hasBeach && c.tempC > beachVacationThreshold
}

func (c city) SkiVacationReady() bool {
	return c.hasMountain && c.tempC < skipVacationThreshold
}

type cities struct {
	cityMap map[string]CityTemp
}

type Cities interface {
	ListAll() []CityTemp
}

// NewCities initialises the Cities data structure by calling the
// ReadData method to read information from file.
func NewCities(reader data.DataReader) (Cities, error) {
	data, err := reader.ReadData()
	if err != nil {
		return nil, err
	}

	cmap := make(map[string]CityTemp)
	for _, r := range data {
		cmap[r.Id] = NewCity(r.Name, r.TempC, r.HasBeach, r.HasMountain)
	}

	return &cities{cityMap: cmap}, nil
}

func (c cities) ListAll() []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		cs = append(cs, rc)
	}
	return cs
}
