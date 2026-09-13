package main

import (
	"log/slog"
	"os"
)

func main() {
	slog.Info("Just Testing, Will Remove Main")

	run(os.Args[1:])

	slog.Info("end")
}

func run(locs []string) error {

	for _, loc := range locs {
		slog.Info("loc", "blah", loc)

		info, err := os.Stat(loc)
		if err != nil {
			return err
		}

		slog.Info("info", "info", info)

	}
	return nil
}
