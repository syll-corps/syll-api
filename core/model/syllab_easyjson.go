// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonE9148b71DecodeGithubComSyllabTeamSyllApiCoreModel(in *jlexer.Lexer, out *SyllabModel) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "DayInfo":
			(out.DayInfo).UnmarshalEasyJSON(in)
		case "Schedules":
			if in.IsNull() {
				in.Skip()
				out.Schedules = nil
			} else {
				in.Delim('[')
				if out.Schedules == nil {
					if !in.IsDelim(']') {
						out.Schedules = make([]Schedule, 0, 0)
					} else {
						out.Schedules = []Schedule{}
					}
				} else {
					out.Schedules = (out.Schedules)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Schedule
					(v1).UnmarshalEasyJSON(in)
					out.Schedules = append(out.Schedules, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9148b71EncodeGithubComSyllabTeamSyllApiCoreModel(out *jwriter.Writer, in SyllabModel) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"DayInfo\":"
		out.RawString(prefix[1:])
		(in.DayInfo).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Schedules\":"
		out.RawString(prefix)
		if in.Schedules == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Schedules {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SyllabModel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9148b71EncodeGithubComSyllabTeamSyllApiCoreModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SyllabModel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9148b71EncodeGithubComSyllabTeamSyllApiCoreModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SyllabModel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9148b71DecodeGithubComSyllabTeamSyllApiCoreModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SyllabModel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9148b71DecodeGithubComSyllabTeamSyllApiCoreModel(l, v)
}
func easyjsonE9148b71DecodeGithubComSyllabTeamSyllApiCoreModel1(in *jlexer.Lexer, out *Schedule) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "time":
			out.Time = string(in.String())
		case "auditorium":
			out.Auditorium = string(in.String())
		case "entity":
			out.Entity = string(in.String())
		case "subject":
			out.Subject = string(in.String())
		case "scheduleStatus":
			out.ScheduleStatus = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9148b71EncodeGithubComSyllabTeamSyllApiCoreModel1(out *jwriter.Writer, in Schedule) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix[1:])
		out.String(string(in.Time))
	}
	{
		const prefix string = ",\"auditorium\":"
		out.RawString(prefix)
		out.String(string(in.Auditorium))
	}
	{
		const prefix string = ",\"entity\":"
		out.RawString(prefix)
		out.String(string(in.Entity))
	}
	{
		const prefix string = ",\"subject\":"
		out.RawString(prefix)
		out.String(string(in.Subject))
	}
	{
		const prefix string = ",\"scheduleStatus\":"
		out.RawString(prefix)
		out.String(string(in.ScheduleStatus))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Schedule) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9148b71EncodeGithubComSyllabTeamSyllApiCoreModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Schedule) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9148b71EncodeGithubComSyllabTeamSyllApiCoreModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Schedule) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9148b71DecodeGithubComSyllabTeamSyllApiCoreModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Schedule) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9148b71DecodeGithubComSyllabTeamSyllApiCoreModel1(l, v)
}
func easyjsonE9148b71DecodeGithubComSyllabTeamSyllApiCoreModel2(in *jlexer.Lexer, out *Day) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "dailer":
			out.Dailer = string(in.String())
		case "date":
			out.Date = string(in.String())
		case "evenStatus":
			out.EvenStatus = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonE9148b71EncodeGithubComSyllabTeamSyllApiCoreModel2(out *jwriter.Writer, in Day) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dailer\":"
		out.RawString(prefix[1:])
		out.String(string(in.Dailer))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.String(string(in.Date))
	}
	{
		const prefix string = ",\"evenStatus\":"
		out.RawString(prefix)
		out.Bool(bool(in.EvenStatus))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Day) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE9148b71EncodeGithubComSyllabTeamSyllApiCoreModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Day) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE9148b71EncodeGithubComSyllabTeamSyllApiCoreModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Day) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE9148b71DecodeGithubComSyllabTeamSyllApiCoreModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Day) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE9148b71DecodeGithubComSyllabTeamSyllApiCoreModel2(l, v)
}
