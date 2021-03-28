package gountries

import (
	"strings"
)

type Continents interface {
	FindContinent(continentName string) (Continent, error)
}

type continents struct {
	continents []Continent
}

type Continent struct {
	Name string
	Code string
}

func NewContinents() Continents {
	return &continents{
		continents: []Continent{
			{Code: "AF", Name: "Africa"},
			{Code: "AN", Name: "Antarctica"},
			{Code: "AS", Name: "Asia"},
			{Code: "EU", Name: "Europe"},
			{Code: "NA", Name: "North America"},
			{Code: "OC", Name: "Oceania"},
			{Code: "SA", Name: "South America"},
		},
	}
}

func (continents *continents) FindContinent(continentName string) (Continent, error) {
	cn := strings.ToUpper(continentName)

	for _, continent := range continents.continents {
		if continent.Code == cn || strings.ToUpper(continent.Name) == cn {
			return continent, nil
		}
	}

	return Continent{}, makeError("Could not find continent with given identifier", continentName)
}

