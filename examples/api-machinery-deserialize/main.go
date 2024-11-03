package main

import (
	"fmt"
	"io"
	"os"

	starwarsv1 "github.com/okontajneroch/starwars/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

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

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	o, gvk, err := serializer.Decode(data, nil, nil)
	if err != nil {
		panic(err)
	}

	// vypišeme Group-VersionKind informácie o objekte
	fmt.Printf("Group: %s\n", gvk.Group)
	fmt.Printf("Version: %s\n", gvk.Version)
	fmt.Printf("Kind: %s\n\n", gvk.Kind)

	// vypišeme 'spec' informácie o objekte
	xwing := o.(*starwarsv1.Starfighter)

	fmt.Printf("Faction: %s\n", xwing.Spec.Faction)
	fmt.Printf("Type: %s\n", xwing.Spec.Type)
	fmt.Printf("Pilot: %s\n", xwing.Spec.Pilot)
}
