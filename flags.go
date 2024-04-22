package main

import "flag"

type Flags struct {
	Add         bool
	Remove      bool
	Generate    bool
	Category    string
	Feature     string
	NumFeatures int
}

func handleFlags() *Flags {
	add := flag.Bool("add", false, "add entry to features file")
	remove := flag.Bool("rm", false, "remove entry to features file")
	generate := flag.Bool("gen", false, "generate feature set")
	category := flag.String("cat", "", "specify category")
	feature := flag.String("ft", "", "add an entry: ")
	numFeatures := flag.Int("nf", -1, "number of random features to generate")

	flag.Parse()

	return &Flags{
		Add:         *add,
		Remove:      *remove,
		Generate:    *generate,
		Category:    *category,
		Feature:     *feature,
		NumFeatures: *numFeatures,
	}
}
