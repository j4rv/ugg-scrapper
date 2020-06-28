package uggscrapper

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// Config stores info about what we are interested in when calling a function (champ, role, patch...)
type Config struct {
	Champ string
	Role  string
	Patch string
}

// QueryParamChamp returns the config's Champ property with the valid's u.gg format (lowercase and no spaces)
//   This allows the users to pass "Aurelion Sol" in the config instead of "aurelionsol"
func (c Config) QueryParamChamp() string {
	res := c.Champ
	res = strings.ReplaceAll(res, " ", "")
	res = strings.ToLower(res)
	return res
}

// QueryParamRole returns the config's Role property with the valid's u.gg format (lowercase and no spaces)
//   This allows the users to pass "Top" in the config instead of "top"
func (c Config) QueryParamRole() string {
	return strings.ToLower(c.Role)
}

// QueryParamPatch returns the config's Patch property with the valid's u.gg format (with underscores)
//   This allows the users to pass "10.12" in the config instead of "10_12"
func (c Config) QueryParamPatch() string {
	return strings.Replace(c.Patch, ".", "_", 1)
}

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

// GetWRsByRank returns a map of Rank - Winrate, with the Ranks constant keys
func GetWRsByRank(cfg Config) map[string]float64 {
	res := make(map[string]float64, len(Ranks))
	for _, rank := range Ranks {
		wr := GetWR(cfg, rank)
		res[rank] = wr
	}
	return res
}

// GetWR makes a request to u.gg and returns the winrate of the config and rank passed
func GetWR(cfg Config, rank string) float64 {
	var wr float64

	c := colly.NewCollector(colly.AllowedDomains("u.gg"))

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

	url := fmt.Sprintf("https://u.gg/lol/champions/%s/build?rank=%s&role=%s&patch=%s",
		cfg.QueryParamChamp(), rank, cfg.QueryParamRole(), cfg.QueryParamPatch())
	c.Visit(url)
	log.Println(rank, wr, url)
	time.Sleep(sleepTimeBetweenRequests)

	return wr
}
