package syllparser

import (
	//"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/syllab-team/syll-api/core/model"
)

const (
	_idxDateG = 8
	_idxDateT = 10
)

const (
	_sepBlank = " "
	_sepBrack = "("
)

const _graphStartDailyM = 159
const _evenSize = 27

type SyllSyntaxer struct {
	controller *syntaxController

	/* 	mod modController */
}

// Late: func opt
func NewSyllSyntaxer() *SyllSyntaxer {
	return &SyllSyntaxer{
		controller: &syntaxController{
			moderator: new(modController),
		},
	}
}

type syntaxMetric struct {
	succes int
	fail   int

	over int
}

// Controller with the custom Syntaxer behavior

type syntaxController struct {
	//Mod syllabMod
	syntaxMetric *syntaxMetric

	moderator *modController
}

const (
	_modGroup = "g"
	_modTeach = "t"
)

var _modMap = map[string]int{
	_modGroup: _idxDateG,
	_modTeach: _idxDateT,
}

type modController struct {
	mod string

	dayIdx int

	startDailerIdx int

	mondaySwitcher bool
	emptySwitcher  bool
}

func (sc *syntaxController) GetMod() string {
	return sc.moderator.mod
}

func (ss *SyllSyntaxer) GetEmptySwitchStatus() bool {
	return ss.controller.moderator.emptySwitcher
}

func (ss *SyllSyntaxer) SetMod(m string) {
	md := ss.controller.moderator

	md.mod = m

	md.dayIdx = _modMap[m]
	md.startDailerIdx = md.dayIdx + 3

	md.mondaySwitcher = m == _modGroup
	md.emptySwitcher = !md.mondaySwitcher
	//println("---PT----", md.mod, md.dayIdx, md.startDailerIdx, md.mondaySwitcher)
}

func (sc *syntaxController) Count(c int) {
	sc.syntaxMetric.over++
	switch c {
	case -1:
		sc.syntaxMetric.fail++
	case 1:
		sc.syntaxMetric.succes++
	}
}

func (ss *SyllSyntaxer) DaySyntaxer(txt string) model.Day {
	var (
		//Late: getters
		dayIdx       = ss.controller.moderator.dayIdx
		dailerIdx    = ss.controller.moderator.startDailerIdx
		switchStatus = ss.controller.moderator.mondaySwitcher
	)

	dt := txt[:dayIdx]
	s := txt[dailerIdx:]

	return model.Day{
		Dailer: func(t string) string {
			if switchStatus && t[1] == _graphStartDailyM {
				return s[:strings.Index(s, _sepBrack)]
			}

			return s[:strings.Index(s, _sepBlank)]
		}(s),
		Date: dt,
		EvenStatus: func(t string) bool {
			return len(txt[strings.Index(txt, _sepBrack):]) == _evenSize
		}(txt),
	}
}

const (
	_timeSelector    = ".time"
	_discSelector    = ".disc"
	_audSelector     = ".aud"
	_teacSelector    = ".teac"
	_teacDivSelector = "div .teac"
)

const (
	_emptyStatusDisc = "empty"
	_emptyTrLen      = 23
)

var (
	_dirtyRune9    rune = 9
	_dirtyRune10   rune = 10
)
var _isDirty = func(r rune) bool {
	return r == _dirtyRune10 || r == _dirtyRune9
}

func (ss *SyllSyntaxer) SchedSyntaxer(e *colly.HTMLElement) model.Schedule {
	d := e.ChildText(_discSelector)
	d = d[:len(d)-1]
	var st int

	return model.Schedule{
		Time: e.ChildText(_timeSelector),
		Auditorium: func() string { 
			_t := e.ChildText(_audSelector)

			var temp []rune
			for _, r := range _t {
				if !_isDirty(r) {
					temp = append(temp, r)
					continue
				}
				break
			}

			return string(temp)
		}(),
		Subject: func(t string) string {
			st = strings.Index(t, _sepBrack)
			if st != -1 {
				return t[:st-1]
			}
			return t
		}(d),
		ScheduleStatus: func(t string) string {
			if st != -1 {
				return t[st+1:]
			}
			// May change the description
			return _emptyStatusDisc
		}(d),
		Entity: func() string {
			t := e.ChildText(_teacDivSelector)
			var _endGroupsMarker = "-"
			t += _endGroupsMarker

			var _sepGroupsRune rune = '/'
			var _limitSize = len(t) - 1
			var temp []rune
			for i, r := range t {
				if i == _limitSize {
					return string(temp)
				}
				if !_isDirty(r) {
					temp = append(temp, r)
					if nxt := rune(byte(t[i+1])); _isDirty(nxt) {
						temp = append(temp, _sepGroupsRune)
					}
				}
			}

			return string(temp)
		}(),
	}
}
