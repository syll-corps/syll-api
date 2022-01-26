package syllparser

import (
	"sync"

	"github.com/gocolly/colly"
	"github.com/syllab-team/syll-api/configs"
	"github.com/syllab-team/syll-api/core/xbus"
	"go.uber.org/zap"
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

type registrBlocker struct {
	blockChannel chan bool
} 

// For count the re-element in the parsed html and escape other
type syllController struct {
	selector selectorSwitcher

	// Late: delete
	limiter limitController
	escaper *htmlEscaper 

	// Blocker of the registr goroutine for the correct collect and push data to the pool 
	blocker *registrBlocker

	// Late: delete
	Counter    int
	LockStatus bool
}

func (sc *syllController) writeBlockStatus() {
	sc.blocker.blockChannel <- true
}
 
func (sc *syllController) waitBlockStatus() {
	<- sc.blocker.blockChannel 
}
type selectorSwitcher struct {
	syllSelector string
	teacSelector string
}

func (sc *syllController) GetSelectorSyll() string {
	return sc.selector.syllSelector
}

func (sc *syllController) GetSelectorTeac() string {
	return sc.selector.teacSelector
}

// Test
type htmlEscaper struct {
	// Late: custom
	mod string

	escaper func(e *colly.HTMLElement) bool

	escapeLimit int
}

func (esc *htmlEscaper) SetLimit(l int) {
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

type parserModerator struct {
	mod string
}
type SyllParser struct {
	// Core of the parser
	Engine *colly.Collector

	// Bus between parser methods and repository.
	BusManager *xbus.CrossPoolManager

	// Storage instance
	Repository interface{}

	// Data race escaper
	Locker *sync.Mutex

	// Logger instance
	SyllLogger *zap.Logger

	Moderator *parserModerator

	// Work struct for the work with the html-element
	HtmlController *syllController

	// Custom struct for the set the behavior of the syntaxer and the count-metrics analyser
	Syntaxer *SyllSyntaxer

	// Config manager for the setup local-controllers and util-functionality.
	Config *configs.SyllConfigManager

	// Controller for link-managment
	Linker *SyllabLinker
}  


func (sp *SyllParser) getController() *syllController {
	return sp.HtmlController
}
func (sp *SyllParser) getSyntaxer() *SyllSyntaxer {
	return sp.Syntaxer
}
func (sp *SyllParser) getCollector() *colly.Collector {
	return sp.Engine
}
func (sp *SyllParser) getBusConstructor() *xbus.ConstructManager {
	return sp.BusManager.Constructor
}


// Bus options.
func _defaultBusManager() *xbus.CrossPoolManager {
	return xbus.NewCrossPoolManager()
}

// Funtional options part
type ParserOption func(*SyllParser)

func WithLogger(l *zap.Logger) ParserOption {
	return func(sp *SyllParser) {
		sp.SyllLogger = l
	}
}

func _defaultLogger() *zap.Logger {
	// Late: add new logger settings and logic
	l, err := zap.NewDevelopment()
	if err != nil {
		return &zap.Logger{}
	}

	return l
}

func WithRepository(r interface{}) ParserOption {
	return func(sp *SyllParser) {
		sp.Repository = r
	}
}

// Late: returns the repository instance
func _deafaultRepository() interface{} {
	const _tempDefaultRepoValue = '1'
	return _tempDefaultRepoValue
}

// ConfigManager and configs options.
func WithConfigManager(cfg *configs.SyllConfigManager) ParserOption {
	return func(sp *SyllParser) {
		sp.Config = cfg
	}
}

func _defaultConfigManager() *configs.SyllConfigManager {
	return configs.NewSyllConfigManager()
}

// Linker options and setuping the links.
func WithLinkerOtherCore(c string) ParserOption {
	return func(sp *SyllParser) {
		sp.Linker.core = c
	}
}

func WithLinkerByOptions() ParserOption {
	return func(sp *SyllParser) {
		l := &SyllabLinker{}
		l.core = sp.Config.GetLinkCore()
		sp.Linker = l
	}
}

// Unrecommended
func _defaultLinker() *SyllabLinker {
	return &SyllabLinker{}
}

// Mod functional options settings
func WithGroupsMod() ParserOption {
	return func(sp *SyllParser) {
		sp.SetParserMod(_modGroup)
	}
}

func WithTeachsMod() ParserOption {
	return func(sp *SyllParser) {
		sp.SetParserMod(_modTeach)
	}
}

func _defaultModerator() *parserModerator {
	return &parserModerator{}
}

func (sp *SyllParser) SetParserMod(m string) {
	sp.Moderator.mod = m

	sp.HtmlController.SetMod(m)
	sp.Syntaxer.SetMod(m)
}

// Syntaxer limiter options
func _defaultSyntaxer() *SyllSyntaxer {
	return NewSyllSyntaxer()
}

// Controllers options.
func _defaultSyllController() *syllController {
	const _syllElementSelector = ".tt tr"

	// Late: deafult-value constructors 
	return &syllController{
		selector: selectorSwitcher{
			syllSelector: _syllElementSelector,
			teacSelector: _teacSelector,
		},
		blocker: &registrBlocker{
			blockChannel: make(chan bool),
		},
		escaper: func() *htmlEscaper {
			e := new(htmlEscaper)
			//Late: change const name
			e.SetLimit(_defEmptyTrLen)
			return e
		}(),
	}
}

//Lock options.
func _defaultLock() *sync.Mutex {
	return &sync.Mutex{}
}

//Engine and collectors logic options.
func WithAsyncCollector() ParserOption {
	const _asyncModMarker = true
	return func(sp *SyllParser) {
		sp.Engine = colly.NewCollector(
			colly.Async(_asyncModMarker),
		)
	}
}

func _defaultEngingeCollector() *colly.Collector {
	return colly.NewCollector()
}

func NewSyllParser(opts ...ParserOption) *SyllParser {
	// Default value settings
	s := &SyllParser{
		Engine:         _defaultEngingeCollector(),
		BusManager:     _defaultBusManager(),
		Config:         _defaultConfigManager(),
		Repository:     _deafaultRepository(),
		SyllLogger:     _defaultLogger(),
		Moderator:      _defaultModerator(),
		HtmlController: _defaultSyllController(),
		Syntaxer:       _defaultSyntaxer(),
		Linker:         _defaultLinker(),
		Locker:         _defaultLock(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
