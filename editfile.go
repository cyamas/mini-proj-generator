package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func handleEdit(editType string, cat string, feat string) {
	var features FeaturesFile
	ReadAndUnmarshal(file, &features)
	if feat != "" {
		if editType == "add" {
			features.Add(cat, feat)
		}
		if editType == "remove" {
			features.Remove(cat, feat)
		}
		UpdateFile(&features, file)
	}
}

type FeaturesFile struct {
	Tool      []string `json:"tool"`
	StdLib    []string `json:"stdLib"`
	ExtLib    []string `json:"extLib"`
	DataType  []string `json:"dataType"`
	Algorithm []string `json:"algorithm"`
}

func ReadAndUnmarshal(file string, ds interface{}) {
	features, err := os.ReadFile(file)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Could not read file: ", err)
	}
	if len(features) > 0 {
		if err := json.Unmarshal(features, &ds); err != nil {
			log.Fatal("Could not unmarshal file: ", err)
		}
	}
}

func UpdateFile(new *FeaturesFile, file string) {
	old, err := os.Create(file)
	if err != nil {
		log.Fatal("Could not read file: ", err)
	}
	defer old.Close()
	if err := json.NewEncoder(old).Encode(new); err != nil {
		log.Fatal("Could not write json to file: ", err)
	}
}

func (features *FeaturesFile) Add(cat string, feat string) {
	switch cat {
	case "std":
		features.StdLib = append(features.StdLib, feat)
	case "ext":
		features.ExtLib = append(features.ExtLib, feat)
	case "dtype":
		features.DataType = append(features.DataType, feat)
	case "algo":
		features.Algorithm = append(features.Algorithm, feat)
	case "tool":
		features.Tool = append(features.Tool, feat)
	}
}

func (features *FeaturesFile) Remove(cat string, feat string) {
	switch cat {
	case "std":
		features.StdLib = removeEntry(feat, features.StdLib)
	case "ext":
		features.ExtLib = removeEntry(feat, features.ExtLib)
	case "tool":
		features.Tool = removeEntry(feat, features.Tool)
	case "dtype":
		features.DataType = removeEntry(feat, features.DataType)
	case "algo":
		features.Algorithm = removeEntry(feat, features.Algorithm)
	}
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
