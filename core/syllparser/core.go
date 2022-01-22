package syllparser

import (
	"log"
	"time"

	"github.com/gocolly/colly"
	"github.com/syllab-team/syll-api/core/xpool"
)

// Core parsing function. Collect the models and set this in LOCAL the xpool.
func (spr *SyllParser) CollectSyllab(link string) (*xpool.XerPool, error) {
	g := spr.Engine
	hr := spr.HtmlController
	sr := spr.Syntaxer

	//log.Println("-------LIMIT>>>-----------", hr.limiter.cursor.curLim)

	p := xpool.NewXerPool()
	g.OnHTML(hr.GetSelectorSyll(), func(e *colly.HTMLElement) {
		// Day string getting by u tag
		if dayTxt := e.ChildText("u"); dayTxt != "" {
			if p.CurStatus() {
				p.PushModel()
			}

			log.Println("-------DAY-----------", dayTxt)
			d := sr.DaySyntaxer(dayTxt)
			log.Println("-------DAY-SER-----------", d)

			p.SetCur(d)
			return
		}

		if hr.GetEscapeStatus(e) {
			println("---ESCAPE STATUS IS OK----")
			return
		}

		sch := sr.SchedSyntaxer(e)
		log.Println("-------SCH-SER-----------", sch)
		p.PushSched(sch)
	})

	g.OnRequest(func(r *colly.Request) {
		println("                            [URL] - ", r.URL.String())
	})

	var err error
	go func() {
		if err = g.Visit(link); err != nil {
			log.Println("---VISIT---", link)
		}
	}()

	g.Wait()

	// Late: gracefull connection logic
	time.Sleep(time.Second * 3)
	return p, err
}

func (spr *SyllParser) CollectSyllabGroup(group string) (*xpool.XerPool, error) {
	l := spr.Linker.MakeGroupUri(group)
	p, err := spr.CollectSyllab(l)

	//c.Rebase()
	return p, err
}

func (spr *SyllParser) CollectSyllabTeach(group string) {
	g := spr.Engine
	c := g.Clone()
	hr := spr.HtmlController
	sr := spr.Syntaxer

	c.OnHTML(hr.GetSelectorSyll(), func(h *colly.HTMLElement) {
		// Day string getting by u tag
		if dayTxt := h.ChildText("u"); dayTxt != "" {
			//Late: pool-logic
			log.Println("\n-------DAY-----------", dayTxt)
			d := sr.DaySyntaxer(dayTxt)
			log.Println("-------DAY-SER-----------", d)
			return
		}

		if hr.GetEscapeStatus(h) {
			//Temp: log
			println("---ESCAPE STATUS IS OK----")
			return
		}

		sch := sr.SchedSyntaxer(h)
		log.Println("-------SCH-SER-----------", sch)
	})

	g.OnHTML(hr.GetSelectorTeac(), func(e *colly.HTMLElement) {
		u := e.ChildAttr("a", "href")
		c.Visit(u)
		c.Wait()
	})

	g.OnRequest(func(r *colly.Request) {
		//Late: add log-func and err-handling
		println("\n                            [TEACHER] - ", r.URL.String())
		println()
	})

	g.Visit(spr.Linker.MakeGroupUri(group))
	g.Wait()
	time.Sleep(time.Second * 3)
}
