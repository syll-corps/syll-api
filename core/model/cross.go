package model

import (
	"log"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type SyllabX struct {
	//	logger instance
	XLogger interface{}

	X   *SyllabCrosser
	Dex *SyllabDecrosser
}

// Custom marshaller from syllab-model to the raw-json bytes
type SyllabCrosser struct {
	writer *jwriter.Writer
}

// Custom unmarshaller from the raw-json bytes to the syllab-model
type SyllabDecrosser struct {
	lexer *jlexer.Lexer
}

// Vanish the Lexer value and remove the used-state
func (x *SyllabX) vanishLexer() {
	x.Dex.lexer = new(jlexer.Lexer)
}

// Marshaling
func (x *SyllabX) Cross(model *SyllabModel) ([]byte, error) {
	w := x.X.writer
	model.MarshalEasyJSON(w)

	//b, err := model.MarshalJSON()
	buf, err := w.BuildBytes()
	if err != nil {
		log.Println("error build bytes", err.Error())
		//logging the error with CrosserLoger
	}

	//cross.Crosser.MarshalEasyJSON()

	return buf, nil
}

// Marshaling
func (x *SyllabX) Decross(b []byte) (*SyllabModel, error) {
	l := x.Dex.lexer
	l.Data = b

	m := &SyllabModel{}
	m.UnmarshalEasyJSON(l)

	err := l.Error()
	if err != nil {
		log.Printf("error build bytes - [%e]", err)
		//logging the error with CrosserLogger
	}

	x.vanishLexer()
	return m, err
}

func NewSyllabX(logger interface{}) *SyllabX {
	return &SyllabX{
		XLogger: logger,
		X: &SyllabCrosser{
			writer: &jwriter.Writer{},
		},
		Dex: &SyllabDecrosser{
			lexer: &jlexer.Lexer{},
		},
	}
}
