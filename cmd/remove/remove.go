package main

import (
	"errors"
	"flag"
	"log"

	handlefile "github.com/cyamas/minproge/pkg"
)

var file = "features.json"

func main() {
	var features handlefile.FeaturesFile
	handlefile.ReadAndUnmarshal(file, features)

	category := flag.String("cat", "", "specify category")
	feature := flag.String("ft", "", "add an entry: ")
	flag.Parse()

	if *feature != "" {
		switch *category {
		case "std":
			features.StdLib = removeEntry(*feature, features.StdLib)
		case "ext":
			features.ExtLib = removeEntry(*feature, features.ExtLib)
		case "tool":
			features.Tool = removeEntry(*feature, features.Tool)
		case "dtype":
			features.DataType = removeEntry(*feature, features.DataType)
		case "algo":
			features.Algorithm = removeEntry(*feature, features.Algorithm)
		}
	}

	handlefile.UpdateFile(&features, file)
}

func removeEntry(feat string, cat []string) []string {
	index, err := getEntryIndex(feat, cat)
	if err != nil {
		log.Println("error in retrieving index: ", err)
		return cat
	}
	return removeEntryByIndex(index, cat)
}

func getEntryIndex(entry string, features []string) (int, error) {
	if len(features) == 0 {
		err := errors.New("category has no entries")
		return -1, err
	}
	for i, feature := range features {
		if feature == entry {
			return i, nil
		}
	}
	err := errors.New("feature does not exist in this category")
	return -1, err
}

func removeEntryByIndex(index int, features []string) []string {
	if len(features) == 1 {
		return []string{}
	}
	if index == len(features)-1 {
		return features[:len(features)-1]
	}
	features[index] = features[len(features)-1]
	return features[:len(features)-1]
}
