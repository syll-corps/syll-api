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

	pool, er := parser.CollectSyllabGroup("622401")
	if er != nil {
		log.Printf("\n\nCRUSH - [%e]", er)
	}
	for _, s := range pool.Pool {
		log.Println("Day - ", s.DayInfo)
		for _, el := range s.Schedules {
			log.Println("Sched - ", el.Auditorium)
		}
	}
	time.Sleep(time.Second * 3)
}
