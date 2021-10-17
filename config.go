package main

import "flag"

type Config struct {
	MainFilepath  string
	SidesFilepath string

	Verbose bool
}

func NewConfig() Config {
	config := Config{}

	flag.StringVar(&config.MainFilepath, "main", "main.json", "Path to the json file containing your main dishes")
	flag.StringVar(&config.SidesFilepath, "sides", "sides.json", "Path to the json file containing your side dishes")

	flag.BoolVar(&config.Verbose, "v", false, "Whether the program will display debug logs")

	flag.Parse()
	return config
}
