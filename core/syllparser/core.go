package syllparser

import (
	"log"
	"time"

	"github.com/gocolly/colly"
)

// Core parsing function. Collect the models and set this in LOCAL the xpool.
func (spr *SyllParser) CollectSyllab(link string) ([][]byte, error) {
	g := spr.Engine
	hr := spr.HtmlController
	sr := spr.Syntaxer

	// bus
	cor := spr.BusManager.Constructor
	g.OnHTML(hr.GetSelectorSyll(), func(e *colly.HTMLElement) {
		// Day string getting by u tag
		if dayTxt := e.ChildText("u"); dayTxt != "" {
			d := sr.DaySyntaxer(dayTxt)
			// Late: ->log
			log.Println("-------DAY-SER-----------", d)
			cor.SwitchCursor(d)
			return
		}

		if hr.GetEscapeStatus(e) {
			println("---ESCAPE STATUS IS OK----")
			return
		}

		sch := sr.SchedSyntaxer(e)
		log.Println("-------SCH-SER-----------", sch)

		// push sched to the trap.
		cor.PushSched(sch)
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

	// Push last model to pool
	if err := cor.PushCrossedModelToPool(); err != nil {
		log.Printf("\n\n [NIL-MODEL-PUSH-TAKING] - [%e]", err)
	}
	return cor.Trap.Pool, err
}

func (spr *SyllParser) CollectSyllabGroup(group string) error {
	l := spr.Linker.MakeGroupUri(group)
	m, err := spr.CollectSyllab(l)
	for _, r := range m {
		log.Printf("ser-model - [%s]", r)
		println()
	}
	//c.Rebase()
	return err
}

func (spr *SyllParser) CollectSyllabTeach(group string) {
	g := spr.getCollector()
	hr := spr.getController()
	sr := spr.getSyntaxer()

	cor := spr.getBusConstructor()
	c := g.Clone()

	c.OnHTML(hr.GetSelectorSyll(), func(h *colly.HTMLElement) {
		// Day string getting by u tag
		if dayTxt := h.ChildText("u"); dayTxt != "" {
			d := sr.DaySyntaxer(dayTxt)

			//Temp: log
			log.Println("-------DAY-SER-----------", d)
			cor.SwitchCursor(d)
			return
		}

		if hr.GetEscapeStatus(h) {
			//Temp: log
			println("---ESCAPE STATUS IS OK----")
			return
		}

		sch := sr.SchedSyntaxer(h)
		cor.PushSched(sch)

		//Temp: log
		log.Println("-------SCH-SER-----------          ", sch)
	})

	g.OnHTML(hr.GetSelectorTeac(), func(e *colly.HTMLElement) {
		u := e.ChildAttr("a", "href")
		c.Visit(u)

		go func() {
			c.Wait()
			if err := cor.PushCrossedModelToPool(); err != nil {

				//Err: log
				log.Printf("[PUSH MODEL PARSE ERROR] - (%v)", err)
			}

			//Log the pool content
			for _, r := range cor.Trap.Pool {
				log.Printf("\n\n---MODEL-IN-POOL----[%s]\n\n", r)
			}

			cor.VanishTrap()
			hr.writeBlockStatus()
		}()

		hr.waitBlockStatus()
	})

	c.OnRequest(func(r *colly.Request) {
		//Late: add log-func and err-handling
		println("\n                            [TEACHER] - ", r.URL.String())
		println()
	})

	g.Visit(spr.Linker.MakeGroupUri(group))
	g.Wait()
}
