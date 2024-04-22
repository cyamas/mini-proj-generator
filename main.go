package main

var file = "features.json"

func main() {
	flags := handleFlags()

	if flags.Generate {
		handleGen(flags.NumFeatures)
		return
	}

	if flags.Add {
		handleEdit("add", flags.Category, flags.Feature)
		return
	}

	if flags.Remove {
		handleEdit("remove", flags.Category, flags.Feature)
		return
	}
}
