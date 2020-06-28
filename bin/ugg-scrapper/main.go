package main

import (
	"flag"
	"log"
	"os"

	"github.com/j4rv/ugg-scrapper/pkg/uggscrapper"
	"github.com/wcharczuk/go-chart"
)

func main() {
	cfg := configFromFlags()
	winrates := uggscrapper.GetWRsByRank(cfg)
	var data []chart.Value
	for _, rank := range uggscrapper.Ranks {
		data = append(data, chart.Value{Label: rank, Value: winrates[rank]})
	}
	renderGraph(cfg, data)
}

func configFromFlags() uggscrapper.Config {
	var cfg uggscrapper.Config
	flag.StringVar(&cfg.Champ, "champ", "aatrox", "The champ to get the winrates graph for. Examples: \"Aurelion Sol\", aatrox, Zoe")
	flag.StringVar(&cfg.Role, "role", "default", "The role for the champ, empty if you want their most used role. Valid values: top, jungle, middle, bot, support")
	flag.StringVar(&cfg.Patch, "patch", "latest", "The LoL patch to get the winrates from. Examples: 10_11, 10_9")
	flag.Parse()

	log.Println("Champ:", cfg.Champ)
	log.Println("Role", cfg.Role)
	log.Println("Patch", cfg.Patch)

	return cfg
}

func renderGraph(cfg uggscrapper.Config, data []chart.Value) {
	graph := chart.BarChart{
		Title:    cfg.Champ + " winrates by rank, role: " + cfg.Role + ", patch: " + cfg.Patch,
		Height:   512,
		BarWidth: 60,
		Bars:     data,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 60,
			},
		},
	}

	f, err := os.Create(cfg.Champ + "-" + cfg.Role + "-" + cfg.Patch + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		panic(err)
	}
}
