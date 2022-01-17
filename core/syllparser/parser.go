package syllparser

import (
	"sync"

	"github.com/gocolly/colly"
)

// For count the re-element in the parsed html and escape other
type syllController struct {
	SyllSelector string

	Counter int

	EscapeLimit int

	LockStatus bool
}

func (sc *syllController) Control() {
	sc.Counter++
	if sc.Counter < sc.EscapeLimit {
		sc.LockStatus = true
		return
	}

	sc.LockStatus = false
}

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
}

func NewSyllParser(repo interface{}, mx *sync.Mutex, lg interface{}) *SyllParser {
	return &SyllParser{
		Engine: colly.NewCollector(
			colly.Async(true),
		),
		Repository: repo,
		Locker:     mx,
		SyllLogger: lg,

		HtmlController: &syllController{
			SyllSelector: syllElementSelector,
			EscapeLimit:  7,
		},
		Syntaxer: NewSyllSyntaxer(),
	}
}
