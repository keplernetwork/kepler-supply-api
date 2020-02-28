package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	flagNodeAPI  = kingpin.Flag("api", "kepler node API").Default("http://127.0.0.1:7413").String()
	flagBindAddr = kingpin.Flag("bind", "API bind address").Default(":10010").String()
)

// {
//   "height": 464343,
//   "last_block_pushed": "006c106b75a18e5fce280d3b39ba3a917cfb93645c9cfe2b784623ffded69387",
//   "prev_block_to_last": "004c49bacc30ae286453cdac6a7ecfff9b8f292be4c799c1a7d28a0fc698cb58",
//   "total_difficulty": 3578050871100
// }
type ChainTip struct {
	Height          int    `json:"height"`
	LastBlockPushed string `json:"last_block_pushed"`
	PrevBlockToLast string `json:"prev_block_to_last"`
	Total           int    `json:"total_difficulty"`
}

const (
	keplerBase      = 10e9
	blockTime       = 60
	hourHeight      = 3600 / blockTime
	yearHeight      = 24 * 7 * 52 * hourHeight
	halvingInterval = 2 * yearHeight
	initialReward   = 1000 * keplerBase
)

// emission returns total supply at height, in nano kepler
func emission(height int) int {
	if height >= halvingInterval {
		panic("please update code to support halving")
	}

	if height == 0 {
		return 42_000_000
	}

	// FIXME: consider halving
	return 42_000_000 + height*1000
}

func getTotalSupply(c echo.Context) error {
	nodeAPI := *flagNodeAPI

	res, err := http.Get(fmt.Sprintf("%s/v1/chain", nodeAPI))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Kepler API error: %s", res.Status)
	}

	dec := json.NewDecoder(res.Body)

	var tip ChainTip
	err = dec.Decode(&tip)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("%d\n", emission(tip.Height)))
}

func run() error {
	kingpin.Parse()

	e := echo.New()
	e.HideBanner = true

	e.GET("supply", getTotalSupply)

	return e.Start(*flagBindAddr)
}

func main() {
	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}
