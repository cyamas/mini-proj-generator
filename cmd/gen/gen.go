package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"

	handlefile "github.com/cyamas/minproge/pkg"
)

var file = "features.json"

func main() {
	var features map[string][]string
	handlefile.ReadAndUnmarshal(file, &features)

	numFeatures := flag.Int("nf", -1, "number of random features to generate")
	flag.Parse()

	cats := getCategoryNames(&features)
	if len(cats) < *numFeatures {
		log.Println("not enough categories for given number of features")
		return
	}

	fmt.Println("\nBuild a small project that uses these features: ")
	for i := 0; i < *numFeatures; i++ {
		getRandoFeature(features)
	}
}

func getCategoryNames(features *map[string][]string) []string {
	var cats []string
	for key := range *features {
		cats = append(cats, key)
	}
	return cats
}

func getRandoFeature(features map[string][]string) {
	cats := getCategoryNames(&features)
	randoCatIndex := rand.Intn(len(cats) - 1)
	entries := features[cats[randoCatIndex]]
	randoEntryIndex := rand.Intn(len(entries))

	fmt.Printf("%v: %v \n", cats[randoCatIndex], entries[randoEntryIndex])
	delete(features, cats[randoCatIndex])

}
