package xbus

import (
	"errors"

	"github.com/google/go-cmp/cmp"
	"github.com/syllab-team/syll-api/core/model"
)

// The mod manager.
type xpoolModerator struct {
	mod string
}

type constructCursor struct {
	cursor string

	initStatus bool
}

type constructTrap struct {
	model model.SyllabModel

	// Serialized models data.
	Pool [][]byte
}

type ConstructManager struct {
	Crosser *model.SyllabX

	Cursor *constructCursor

	Trap *constructTrap

	mod string
}

func (cm *ConstructManager) setCursor(c string) {
	cm.Cursor.cursor = c
}

func (cm *ConstructManager) getInitStatus() bool {
	return cm.Cursor.initStatus
}

var _emptyModelTemplate = &model.SyllabModel{}
func (cm *ConstructManager) getModelPushStatus() bool {
	m := cm.getModel()
	return !cmp.Equal(m, _emptyModelTemplate)
}

func (cm *ConstructManager) setTrap() {
	cm.Cursor.initStatus = false
}

const _defaultCursorInitStatus = true
func setDefaultCursor() *constructCursor {
	c := &constructCursor{}
	c.initStatus = _defaultCursorInitStatus
	return c
}

func (cm *ConstructManager) setDay(d model.Day) {
	cm.Trap.model.DayInfo = d
}

func (cm *ConstructManager) setStep(d model.Day) {
	cm.setDay(d)
	c := cm.buildDayCursor(d)
	cm.setCursor(c)
}

func (cm *ConstructManager) PushSched(s model.Schedule) {
	cm.Trap.model.Schedules = append(cm.Trap.model.Schedules, s)
}

func (cm *ConstructManager) GetCursor() string {
	return cm.Cursor.cursor
}

func (cm *ConstructManager) buildDayCursor(d model.Day) string {
	return d.Dailer
}

func (cm *ConstructManager) SwitchCursor(d model.Day) {
	if cm.getInitStatus() {
		cm.setTrap()
		cm.setStep(d)
		return
	}

	// serialization and pushing to the pool
	cm.PushCrossedModelToPool()

	cm.vanishModel()
	cm.setStep(d)
}

func (cm *ConstructManager) vanishModel() {
	cm.Trap.model.Schedules = make([]model.Schedule, 0)
}

func (cm *ConstructManager) VanishPool() {
	cm.Trap.Pool = make([][]byte, 0)
}

func (cm *ConstructManager) VanishTrap() {
	cm.Trap = &constructTrap{}
}

func newConstructManager() *ConstructManager {
	const _defaultLoggerSettings = "_l"
	return &ConstructManager{
		Crosser: model.NewSyllabX(_defaultLoggerSettings),
		Cursor:  setDefaultCursor(),
		Trap:    &constructTrap{},
	}
}

func (cm *ConstructManager) getModel() *model.SyllabModel {
	return &cm.Trap.model
}

var errNilModelPushTaking = errors.New("nil model - crossed is blocked")
func (cm *ConstructManager) PushCrossedModelToPool() error {
	if !cm.getModelPushStatus() {
		return errNilModelPushTaking
	}

	b, err := cm.Crosser.Cross(cm.getModel())
	if err != nil {
		return err
	}
	cm.Trap.Pool = append(cm.Trap.Pool, b)
	return nil
}

// Main struct responsibiled of the pool staging and pulling the serialized data to Repo.
type CrossPoolManager struct {
	// Construct the model by parts and serialize it.
	Constructor *ConstructManager

	// Connect with the redis and send data to.
	Connector *ConnectManager

	// Switch the mod and set the mod-set logic.
	moderator *xpoolModerator
}

func NewCrossPoolManager() *CrossPoolManager {
	return &CrossPoolManager{
		Constructor: newConstructManager(),
	}
}
