package main

import (
	"log"
	"time"

	"github.com/syllab-team/syll-api/configs"
	"github.com/syllab-team/syll-api/core/syllparser"
)

func main() {
	cfg := configs.NewSyllConfigManagerTest()
	if err := cfg.RiseSyllConfigs(); err != nil {
		log.Printf("with config - [%e]", err)
	}
	parser := syllparser.NewSyllParser(
		syllparser.WithConfigManager(cfg),
		syllparser.WithLinkerByOptions(),
		syllparser.WithAsyncCollector(),
		syllparser.WithGroupsMod(),
	)

	er := parser.CollectSyllabGroup("221211")
	if er != nil {
		log.Printf("\n\nCRUSH - [%e]", er)
	}
	time.Sleep(time.Second * 3)
}
