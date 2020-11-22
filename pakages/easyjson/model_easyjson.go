// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package easyjson

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

func easyjsonC80ae7adDecodeEasyjsonStudy(in *jlexer.Lexer, out *Foo) {
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
		case "uuid":
			out.UUID = string(in.String())
		case "state":
			out.State = string(in.StringIntern())
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
func easyjsonC80ae7adEncodeEasyjsonStudy(out *jwriter.Writer, in Foo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"uuid\":"
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
	}
	{
		const prefix string = ",\"state\":"
		out.RawString(prefix)
		out.String(string(in.State))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Foo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeEasyjsonStudy(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Foo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeEasyjsonStudy(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Foo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeEasyjsonStudy(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Foo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeEasyjsonStudy(l, v)
}
