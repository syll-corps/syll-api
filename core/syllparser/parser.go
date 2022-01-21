package syllparser

import (
	"sync"

	"github.com/gocolly/colly"
	"github.com/syllab-team/syll-api/configs"
)

type limitCursor struct {
	// ??
	curGroup string

	curLim int
}

type limitController struct {
	// Cursor of the group in the groupPool provides the easier escape logic.
	cursor limitCursor

	// Late
	limitPool map[string]int
}

// Marker of the teach-mod of parsing. Another marker is the number of the group (221201)
const _markTeach = "T"

// For count the re-element in the parsed html and escape other
type syllController struct {
	SyllSelector string

	limiter limitController 
	escaper *htmlEscaper

	Counter    int
	LockStatus bool
} 

// Test
type htmlEscaper struct {
	// Late: custom
	mod string 

	escaper func(e *colly.HTMLElement) bool 

	escapeLimit int
} 

func(esc *htmlEscaper) SetLimit(l int) {
	esc.escapeLimit = l
}

func (sc *syllController) SetMod(m string) { 
	esc := sc.escaper
	esc.mod = m  

	switch m {
	case _modTeach:
		esc.escaper = func(e *colly.HTMLElement) bool {
			return len(e.Text) == esc.escapeLimit
		}
	case _modGroup: 
		esc.escaper = func(e *colly.HTMLElement) bool {
			return e.ChildText(".aud") == ""
		}
	}
} 

func (sc *syllController) GetEscapeStatus(e *colly.HTMLElement) bool {
	return sc.escaper.escaper(e)
}

// Test LRU-cash value. Late:  -> modController.pool + logic auto push into the map.
var _testLimitsMap = map[string]int{
	_markTeach: 0,
	"221201":   8,
	"221211":   7,
}

func (sc *syllController) SetCur(marker string) {
	sc.limiter.cursor.curGroup = marker
	sc.limiter.cursor.curLim = _testLimitsMap[marker]
	println("---LIMIT----", sc.limiter.cursor.curLim)
}

func (sc *syllController) Control() {
	sc.Counter++
	if sc.Counter < sc.limiter.cursor.curLim {
		sc.LockStatus = true
		println("----LOCK-STATUS-----", sc.LockStatus)
		return
	}

	sc.LockStatus = false
}

// Finish operaiton in the sched-parsing
func (sc *syllController) Rebase() {
	sc.Counter = 0
	sc.LockStatus = false
}

// Status getter
func (sc *syllController) GetLocker() bool {
	return sc.LockStatus
}

type SyllParser struct {
	// Core of the parser
	Engine *colly.Collector

	// Storage instance
	Repository interface{}

	// Data race escaper
	Locker *sync.Mutex

	// Logger instance
	SyllLogger interface{}

	//	Work struct for the work with the html-element
	HtmlController *syllController

	// Custom struct for the set the behavior of the syntaxer and the count-metrics analyser
	Syntaxer *SyllSyntaxer

	// Config manager for the setup local-controllers and util-functionality.
	Config *configs.SyllConfigManager

	// Controller for link-managment
	Linker *SyllabLinker
}

func NewSyllParser(cfg *configs.SyllConfigManager, repo interface{},
	mx *sync.Mutex, lg interface{}) *SyllParser {
	s := &SyllParser{
		Engine: colly.NewCollector(
			colly.Async(true),
		),
		Repository: repo,
		Locker:     mx,
		SyllLogger: lg,

		// Late: func opt
		HtmlController: &syllController{
			SyllSelector: syllElementSelector,
			escaper: func() *htmlEscaper {
				e := new(htmlEscaper)
				//Late: change const name
				e.SetLimit(_emptyTrLen)
				return e
			}(),
		},
		Syntaxer: NewSyllSyntaxer(),
		Config:   cfg,
	}

	s.Linker = &SyllabLinker{
		core: s.Config.GetLinkCore(),
	}

	return s
}
