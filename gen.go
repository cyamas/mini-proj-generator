package main

import (
	"fmt"
	"log"
	"math/rand"
)

func handleGen(numFeatures int) {
	var features map[string][]string
	ReadAndUnmarshal(file, &features)
	cats := getCategoryNames(&features)
	if len(cats) < numFeatures {
		log.Println("not enough categories for given number of features")
		return
	}

	fmt.Println("\nBuild a small project that uses these features: ")
	for i := 0; i < numFeatures; i++ {
		cat, ft := selectRandoFeature(features)
		fmt.Printf("%v: %v", cat, ft)
	}
}

func getCategoryNames(features *map[string][]string) []string {
	var cats []string
	for key := range *features {
		cats = append(cats, key)
	}
	return cats
}

func selectRandoFeature(features map[string][]string) (string, string) {
	cats := getCategoryNames(&features)
	randoCatIndex := 0
	if len(cats) > 1 {
		randoCatIndex = rand.Intn(len(cats))
	}
	entries := features[cats[randoCatIndex]]
	randoEntryIndex := rand.Intn(len(entries))

	delete(features, cats[randoCatIndex])
	return cats[randoCatIndex], entries[randoEntryIndex]
}
