package main

import (
	"gpd/db/couchdb"
	"gpd/puzzler"

	"flag"
)

func app(is_daemon bool, n_clues uint8) {
	for {
		db := couchdb.New()
		g := db.Solution()
		p := puzzler.Make(g, n_clues)
		db.StorePuzzle(p)

		if !is_daemon {
			break
		}

	}

}

func main() {
	num_clues := flag.Uint("n", 38, "number of clues (default 38)")
	is_daemon := flag.Bool("d", false, "run as daemon")
	flag.Parse()
	app(*is_daemon, uint8(*num_clues))
}
