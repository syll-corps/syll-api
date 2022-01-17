package model

//easyjson:json
type Day struct {
	// Daily marker for the ui
	Dailer string `json:"dailer"`

	// Date of the schedule
	Date string `json:"date"`

	// Is even the week status
	EvenStatus bool `json:"evenStatus"`
}

//easyjson:json
type Schedule struct {
	// TIme of the class with the subj
	Time string `json:"time"`

	Auditorium string `json:"auditorium"`
	Teacher    string `json:"teacher"`
	Subject    string `json:"subject"`

	// Status of the class - prob is the bool
	ScheduleStatus string `json:"scheduleStatus"`
}

//easyjson:json
type SyllabModel struct {
	DayInfo   Day
	Schedules []Schedule
}
