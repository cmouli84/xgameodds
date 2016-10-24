package main

import (
	"fmt"

	"github.com/cmouli84/xgameodds/infrastructure"
)

func main() {
	fmt.Println(string(infrastructure.GetHTTPResponse("http://api.thescore.com/nfl/schedule")))
}
