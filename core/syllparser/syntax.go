package syllparser

import (
	//"fmt"

	"strings"

	"github.com/gocolly/colly"
	"github.com/syllab-team/syll-api/core/model"
)

const (
	_idxDate       = 8
	_idxStartDaily = _idxDate + 3
)

const (
	_sepBlank = " "
	_sepBrack = "("
)

const _graphStartDailyM = 159
const _evenSize = 27

type SyllSyntaxer struct {
	controller *syntaxController
}

func NewSyllSyntaxer() *SyllSyntaxer {
	return &SyllSyntaxer{
		controller: new(syntaxController),
	}
}

// Mod of the syllab. Set of value is - {G, T}
//const (
//	_gMod = 'G'
//	_tMod = 'T'
//)

//type syllabMod string

type syntaxMetric struct {
	succes int
	fail   int

	over int
}

// Controller with the custom Syntaxer behavior
type syntaxController struct {
	//Mod syllabMod

	syntaxMetric *syntaxMetric
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
	//19.01.22 - Среда (нечётная неделя)
	//date - :8
	dt := txt[:_idxDate]
	s := txt[_idxStartDaily:]
	return model.Day{
		Dailer: func(t string) string {
			if t[1] == _graphStartDailyM {
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

const _idxAud = 14
const (
	_timeSelector = ".time"
	_discSelector = ".disc"
	_audSelector  = ".aud"
	_teacSelector = ".teac"
)

const (
	_emptyStatusDisc = "empty"
)

func (ss *SyllSyntaxer) SchedSyntaxer(e *colly.HTMLElement) model.Schedule {
	d := e.ChildText(_discSelector)
	d = d[:len(d)-1]
	var st int

	//println("-PETRAKOVA-", e.ChildText(_teacSelector))
	return model.Schedule{
		Time: e.ChildText(_timeSelector),
		Auditorium: func() string {
			_t := e.ChildText(_audSelector)[:_idxAud]
			return strings.TrimSpace(_t)
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
		Teacher: e.ChildText(_teacSelector),
	}
}
