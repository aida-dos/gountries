package gountries

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContinentByCodeOk(t *testing.T) {
	continentCode := "eu"

	continent, err := continentsTest.FindContinent(continentCode)
	if err != nil {
		t.Fail()
	}

	assert.Equal(
		t, "Europe", continent.Name,
		fmt.Sprintf("Search %s continent with code '%s' returned %s",
			"Europe", continentCode, continent.Code,
		),
	)
}

func TestFindContinentByNameOk(t *testing.T) {
	continentName := "Asia"

	continent, err := continentsTest.FindContinent(continentName)

	if err != nil {
		t.Fail()
	}

	assert.Equal(
		t, continentName, continent.Name,
		fmt.Sprintf("Search %s continent returned %s",
			continentName, continent.Code,
		),
	)
}

func TestFindContinentByNameError(t *testing.T) {
	continentName := "Portugal"

	_, err := continentsTest.FindContinent(continentName)
	if err != nil {
		assert.Equal(t, "gountries error. Could not find continent with given identifier: " + continentName, err.Error())
	} else {
		t.Fail()
	}
}
