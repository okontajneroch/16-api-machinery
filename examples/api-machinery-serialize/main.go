package main

import (
	"os"

	starwarsv1 "github.com/okontajneroch/starwars/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

func createXWing() *starwarsv1.Starfighter {
	xwing := &starwarsv1.Starfighter{}

	xwing.APIVersion = "starwars.okontajneroch.sk/v1"
	xwing.Kind = "Starfighter"
	xwing.ObjectMeta.Name = "x-wing-1"

	xwing.Spec = starwarsv1.StarfighterSpec{
		Faction: starwarsv1.Rebellion,
		Type:    "x-wing",
		Pilot:   "Luke Skywalker",
	}

	return xwing
}

func main() {
	scheme := runtime.NewScheme()
	starwarsv1.AddToScheme(scheme)

	serializer := json.NewSerializerWithOptions(
		json.SimpleMetaFactory{},
		scheme,
		scheme,
		json.SerializerOptions{
			Yaml: true,
		},
	)

	xwing := createXWing()
	serializer.Encode(xwing, os.Stdout)
}
