package main

import (
	"sync"
	"time"

	"github.com/syllab-team/syll-api/configs"
	"github.com/syllab-team/syll-api/core/syllparser"
)

func main() {
	cfg := configs.NewSyllConfigManagerTest()
	if err := cfg.RiseSyllConfigs(); err != nil {
		println("----CFG-----", cfg.RiseSyllConfigs().Error())
	}
	parser := syllparser.NewSyllParser(
		cfg,
		12,
		&sync.Mutex{},
		12,
	)

	parser.CollectSyllabTeach("622401")
	time.Sleep(time.Second * 3)
	//println(viper.GetString("group"))
}