package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/syllab-team/syll-api/core/model"
	"github.com/syllab-team/syll-api/core/syllparser"
)

func main() {
	parser := syllparser.NewSyllParser(
		12,
		&sync.Mutex{},
		12,
	)

	pool, er := parser.CollectSyllabByGroup("12")
	if er != nil {
		log.Printf("\n\nCRUSH - [%e]", er)
	}
	for _, s := range pool.Pool {
		log.Println("Day - ", s.DayInfo)
		for _, el := range s.Schedules {
			log.Println("Sched - ", el.Auditorium)
		}
	}
}

func visitTimer(t time.Time, c *colly.Collector) {
	c.Visit("http://schedule.tsu.tula.ru/?group=221201")
	fmt.Println(time.Since(t))
}

func MA(model *model.SyllabModel) ([]byte, error) {
	return json.Marshal(model)
}

func U(b []byte) (*model.SyllabModel, error) {
	m := &model.SyllabModel{}
	err := json.Unmarshal(b, m)
	return m, err
}
