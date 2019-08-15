package overwatch

import (
	"net/http"
	"net/http/httptest"
	"service/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalHeroResponse(t *testing.T) {
	// Define our test model
	expected := &models.Hero{
		ID:       1,
		Name:     "Ana",
		RealName: "Ana Amari",
		Health:   200,
		Abilities: []*models.Ability{
			{
				ID:          1,
				Name:        "Biotic Rifle",
				Description: "Ana’s rifle shoots darts that can restore health to her allies or deal ongoing damage to her enemies. She can use the rifle’s scope to zoom in on targets and make highly accurate shots.",
				Ultimate:    false,
			},
			{
				ID:          2,
				Name:        "Sleep Dart",
				Description: "Ana fires a dart from her sidearm, rendering an enemy unconscious (though any damage will rouse them).",
				Ultimate:    false,
			},
			{
				ID:          3,
				Name:        "Biotic Grenade",
				Description: "Ana tosses a biotic bomb that deals damage to enemies and heals allies in a small area of effect. Affected allies briefly receive increased healing from all sources, while enemies caught in the blast cannot be healed for a few moments.",
				Ultimate:    false,
			},
			{
				ID:          4,
				Name:        "Nano Boost",
				Description: "After Ana hits one of her allies with a combat boost, they temporarily move faster, deal more damage, and take less damage from enemies’ attacks.",
				Ultimate:    true,
			},
		},
	}

	// Set up our mock response and marshal it to JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//json.NewEncoder(w).Encode(expected)
		//w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":1,"name":"Ana","description":"Ana’s versatile arsenal allows her to affect heroes all over the battlefield. Her Biotic Rifle rounds and Biotic Grenades heal allies and damage or impair enemies; her sidearm tranquilizes key targets, and Nano Boost gives one of her comrades a considerable increase in power.","health":200,"armour":0,"shield":0,"real_name":"Ana Amari","age":60,"abilities":[{"id":1,"name":"Biotic Rifle","description":"Ana’s rifle shoots darts that can restore health to her allies or deal ongoing damage to her enemies. She can use the rifle’s scope to zoom in on targets and make highly accurate shots.","is_ultimate":false,"url":"https://overwatch-api.net/api/v1/ability/1"},{"id":2,"name":"Sleep Dart","description":"Ana fires a dart from her sidearm, rendering an enemy unconscious (though any damage will rouse them).","is_ultimate":false,"url":"https://overwatch-api.net/api/v1/ability/2"},{"id":3,"name":"Biotic Grenade","description":"Ana tosses a biotic bomb that deals damage to enemies and heals allies in a small area of effect. Affected allies briefly receive increased healing from all sources, while enemies caught in the blast cannot be healed for a few moments.","is_ultimate":false,"url":"https://overwatch-api.net/api/v1/ability/3"},{"id":4,"name":"Nano Boost","description":"After Ana hits one of her allies with a combat boost, they temporarily move faster, deal more damage, and take less damage from enemies’ attacks.","is_ultimate":true,"url":"https://overwatch-api.net/api/v1/ability/4"}]}`))
	}))
	defer server.Close()

	t.Log(server.URL)
	// re-assign the base url
	baseURI = server.URL

	c, err := New()
	assert.Nil(t, err)

	result, err := c.Hero(0)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
