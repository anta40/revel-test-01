package controllers

import (
	"revel-test-01/app/models"

	"github.com/revel/revel"
)

type Beer struct {
	*revel.Controller
}

var beers = []models.Beer{
	models.Beer{1, "La Trappe Quadrupel Oak Aged", "Ale", "Bierbrouwerij De Koningshoeven"},
	models.Beer{2, "Cocoa Psycho", "Porter", "BrewDog"},
	models.Beer{3, "American Dream", "Lager", "Mikkeller"},
	models.Beer{4, "Javanese Heritage", " Wheat Beer", "Jogja123 Brewery"},
}

func (c Beer) List() revel.Result {
	return c.RenderJSON(beers)
}

func (c Beer) Show(beerID int) revel.Result {
	var res models.Beer

	for _, beer := range beers {
		if beer.ID == beerID {
			res = beer
		}
	}

	if res.ID == 0 {
		return c.NotFound("Could not find beer")
	}

	return c.RenderJSON(res)
}
