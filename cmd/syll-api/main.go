package main

import (
	"log"
	"time"

	"github.com/syllab-team/syll-api/configs"
	"github.com/syllab-team/syll-api/core/syllparser"
)

func main() {
	cfg := configs.NewSyllConfigManager()
	if err := cfg.RiseSyllConfigs(); err != nil {
		log.Printf("with config - [%e]", err)
	}
	parser := syllparser.NewSyllParser(
		syllparser.WithConfigManager(cfg),
		syllparser.WithLinkerByOptions(),
		syllparser.WithAsyncCollector(),
		syllparser.WithGroupsMod(),
	)

	parser.CollectSyllabTeach("622401")
	time.Sleep(time.Second * 3)
	//println(viper.GetString("group"))
}
