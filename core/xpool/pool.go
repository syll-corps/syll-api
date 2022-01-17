package xpool

import (
	"github.com/syllab-team/syll-api/core/model"
)

var _initM = func() []model.SyllabModel {
	return make([]model.SyllabModel, 0)
}

var _initS = func() []model.Schedule {
	return make([]model.Schedule, 0)
}

type PoolCursor struct {
	poolSched []model.Schedule

	cursor bool
}

type XerPool struct {
	Pool []model.SyllabModel

	CurDay model.Day
	// Current section of the syllab-model
	Cursor PoolCursor
}

func (xp *XerPool) SetCur(d model.Day) {
	xp.CurDay = d
	xp.Cursor.cursor = true
}

func (xp *XerPool) PushSched(s model.Schedule) {
	xp.Cursor.poolSched = append(xp.Cursor.poolSched, s)
}

func (xp *XerPool) CurStatus() bool {
	return xp.Cursor.cursor
}

func (xp *XerPool) PushModel() {
	xp.Pool = append(xp.Pool, func() model.SyllabModel {
		return model.SyllabModel{
			DayInfo:   xp.CurDay,
			Schedules: xp.Cursor.poolSched,
		}
	}())
	xp.Cursor.poolSched = _initS()
}

func (xp *XerPool) RemAll() {
	xp.Pool = _initM()
}

func NewXerPool() *XerPool {
	return &XerPool{
		Pool:   _initM(),
		CurDay: model.Day{},
		Cursor: PoolCursor{
			poolSched: _initS(),
		},
	}
}
