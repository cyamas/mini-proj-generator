package handlefile

import (
	"encoding/json"
	"log"
	"os"
)

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
		if err := json.Unmarshal(features, ds); err != nil {
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
