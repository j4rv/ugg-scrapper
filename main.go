package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/wcharczuk/go-chart"
)

// U.GG rank query strings
const (
	Iron        = "iron"
	Bronze      = "bronze"
	Silver      = "silver"
	Gold        = "gold"
	Platinum    = "platinum"
	Diamond     = "diamond"
	Master      = "master"
	Grandmaster = "grandmaster"
	Challenger  = "challenger"
	PlatPlus    = "platinum_plus"
	DiamPlus    = "diamond_plus"
	MasterPlus  = "master_plus"
	All         = "overall"
)

// U.GG position query strings
const (
	Default = ""
	Top     = "top"
	Jungle  = "jungle"
	Mid     = "middle"
	Bot     = "bot"
	Supp    = "support"
)

// To prevent u.gg from banning my ip for spamming requests...
const sleepTimeBetweenRequests = 1 * time.Second

// Ranks from U.GG ordered from lower to higher
var Ranks = []string{Iron, Bronze, Silver, Gold, Platinum, Diamond, MasterPlus}

type config struct {
	Champ string
	Role  string
	//Patch string TODO
}

func configFromFlags() config {
	var cfg config
	flag.StringVar(&cfg.Champ, "champ", "aatrox", "The champ to get the winrates graph for")
	flag.StringVar(&cfg.Role, "role", "default", "The role for the champ, empty if you want their most used role. Valid values: top, jungle, middle, bot, support.")
	flag.Parse()

	log.Println("Champ:", cfg.Champ)
	log.Println("Role", cfg.Role)

	return cfg
}

func main() {
	cfg := configFromFlags()
	winrates := getWRsByRank(cfg)
	var data []chart.Value
	for _, rank := range Ranks {
		data = append(data, chart.Value{Label: rank, Value: winrates[rank]})
	}
	renderGraph(cfg, data)
}

func renderGraph(cfg config, data []chart.Value) {
	graph := chart.BarChart{
		Title:    cfg.Champ + " winrates by rank, role: " + cfg.Role,
		Height:   512,
		BarWidth: 60,
		Bars:     data,
		Background: chart.Style{
			Padding: chart.Box{
				Top: 60,
			},
		},
	}

	f, err := os.Create(cfg.Champ + "-" + cfg.Role + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = graph.Render(chart.PNG, f)
	if err != nil {
		panic(err)
	}
}

func getWRsByRank(cfg config) map[string]float64 {
	res := make(map[string]float64, len(Ranks))
	for _, rank := range Ranks {
		wr := getWR(cfg, rank)
		res[rank] = wr
	}
	return res
}

func getWR(cfg config, rank string) float64 {
	var wr float64
	c := colly.NewCollector(
		colly.AllowedDomains("u.gg"),
	)

	c.OnHTML("div .win-rate > .value", func(e *colly.HTMLElement) {
		var err error
		strWR := strings.Replace(e.Text, "%", "", 1)
		if strWR == "" {
			panic("Could not get String WR for: " + cfg.Champ + " - " + cfg.Role)
		}
		wr, err = strconv.ParseFloat(strWR, 64)
		if err != nil {
			panic("GetWR:" + err.Error())
		}
	})

	c.Visit("https://u.gg/lol/champions/" + cfg.Champ + "/build?rank=" + rank + "&role=" + cfg.Role)
	time.Sleep(sleepTimeBetweenRequests)

	log.Println(rank, wr)

	return wr
}
