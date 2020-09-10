package main

import (
	"gpd/db/couchdb"
	"gpd/puzzler"

	"flag"
)

func main() {

	num_clues := flag.Uint("n", 38, "number of clues (default 38)")
	flag.Parse()
	db := couchdb.New()
	g := db.Solution()
	p := puzzler.Make(g, uint8(*num_clues))

	db.StorePuzzle(p)

}
