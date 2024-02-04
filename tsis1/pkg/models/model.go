package models

import "strings"

type Actor struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	Role     string `json:"role"`
	Image    string `json:"image"`
}

var actors = []Actor{
	{Name: "Leonardo DiCaprio", Age: "47", Role: "Cobb", Image: "https://example.com/leonardo.jpg"},
	{Name: "Joseph Gordon-Levitt", Age: "40", Role: "Arthur", Image: "https://example.com/joseph.jpg"},
	{Name: "Ellen Page", Age: "35", Role: "Ariadne", Image: "https://example.com/ellen.jpg"},
	{Name: "Tom Hardy", Age: "44", Role: "Eames", Image: "https://example.com/tom.jpg"},
	{Name: "Ken Watanabe", Age: "61", Role: "Saito", Image: "https://example.com/ken.jpg"},
	
}

func GetAllActors() []Actor {
	return actors
}

func GetActorByName(name string) Actor {
	for i := 0; i < len(actors); i++ {
		actor := actors[i]
		nameWithoutSpaces := strings.ReplaceAll(actor.Name, " ", "")
		nameWithoutSpaces = strings.ReplaceAll(nameWithoutSpaces, "/", "")
		if nameWithoutSpaces == name {
			return actor
		}
	}
	return Actor{}
}
