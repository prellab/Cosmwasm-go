// Code generated by tinyjson for marshaling/unmarshaling. DO NOT EDIT.

package src

import (
	tinyjson "github.com/CosmWasm/tinyjson"
	jlexer "github.com/CosmWasm/tinyjson/jlexer"
	jwriter "github.com/CosmWasm/tinyjson/jwriter"
)

// suppress unused package warning
var (
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ tinyjson.Marshaler
)

func tinyjson4e1b3cc1DecodeGithubComCosmwasmCosmwasmGoExampleHackatomSrc(in *jlexer.Lexer, out *State) {
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
		case "VERIFIER":
			out.Verifier = string(in.String())
		case "BENEFICIARY":
			out.Beneficiary = string(in.String())
		case "FUNDER":
			out.Funder = string(in.String())
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
func tinyjson4e1b3cc1EncodeGithubComCosmwasmCosmwasmGoExampleHackatomSrc(out *jwriter.Writer, in State) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"VERIFIER\":"
		out.RawString(prefix[1:])
		out.String(string(in.Verifier))
	}
	{
		const prefix string = ",\"BENEFICIARY\":"
		out.RawString(prefix)
		out.String(string(in.Beneficiary))
	}
	{
		const prefix string = ",\"FUNDER\":"
		out.RawString(prefix)
		out.String(string(in.Funder))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v State) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	tinyjson4e1b3cc1EncodeGithubComCosmwasmCosmwasmGoExampleHackatomSrc(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalTinyJSON supports tinyjson.Marshaler interface
func (v State) MarshalTinyJSON(w *jwriter.Writer) {
	tinyjson4e1b3cc1EncodeGithubComCosmwasmCosmwasmGoExampleHackatomSrc(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *State) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	tinyjson4e1b3cc1DecodeGithubComCosmwasmCosmwasmGoExampleHackatomSrc(&r, v)
	return r.Error()
}

// UnmarshalTinyJSON supports tinyjson.Unmarshaler interface
func (v *State) UnmarshalTinyJSON(l *jlexer.Lexer) {
	tinyjson4e1b3cc1DecodeGithubComCosmwasmCosmwasmGoExampleHackatomSrc(l, v)
}