package main

import (
	"flag"

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
			features.StdLib = append(features.StdLib, *feature)
		case "ext":
			features.ExtLib = append(features.ExtLib, *feature)
		case "dtype":
			features.DataType = append(features.DataType, *feature)
		case "algo":
			features.Algorithm = append(features.Algorithm, *feature)
		case "tool":
			features.Tool = append(features.Tool, *feature)
		}
	}
	handlefile.UpdateFile(&features, file)
}
