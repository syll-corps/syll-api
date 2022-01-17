package syllparser

import (
	"log"
	"time"

	"github.com/gocolly/colly"
	"github.com/syllab-team/syll-api/core/xpool"
)

const syllElementSelector = ".tt tr"

func (spr *SyllParser) CollectSyllabByGroup(group string) (*xpool.XerPool, error) {
	g := spr.Engine
	hc := spr.HtmlController
	sc := spr.Syntaxer

	//ctr := 0
	p := xpool.NewXerPool()
	g.OnHTML(hc.SyllSelector, func(e *colly.HTMLElement) {
		//println("----HTML-----", e.Text)
		hc.Control()
		if hc.GetLocker() {
			return
		}

		// Day string getting by u tag
		if dayTxt := e.ChildText("u"); dayTxt != "" {
			if p.CurStatus() {
				p.PushModel()
			}

			log.Println("-------DAY-----------", dayTxt)
			d := sc.DaySyntaxer(dayTxt)
			log.Println("-------DAY-SER-----------", d)

			p.SetCur(d)
			return
		}

		sch := sc.SchedSyntaxer(e)
		log.Println("-------SCH-SER-----------", sch)

		p.PushSched(sch)
	})

	var err error
	g.OnError(func(r *colly.Response, e error) {
		sc.controller.Count(-1)
		// log moment
		log.Printf("parse error - [%v]", e)
		err = e
	})

	go func() {
		if er := g.Visit("http://schedule.tsu.tula.ru/?group=221201"); er != nil {
			log.Println("---VISIT---")
		}

	}()

	g.Wait()
	time.Sleep(time.Second * 2)
	return p, err
}

func (spr *SyllParser) CollectSyllabByName(name string) error {

	return nil
}
