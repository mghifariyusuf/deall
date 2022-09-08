package errorCustomStatus

import (
	"fmt"
	"strings"
)

type Error struct {
	Status string       `json:"status"`
	Code   string       `json:"code"`
	Title  string       `json:"title"`
	Detail string       `json:"detail"`
	Source *ErrorSource `json:"source,omitempty"`
	Op     string       `json:"-"`
	Err    error        `json:"-"`
}

type ErrorSource struct {
	Parameter string `json:"parameter,omitempty"`
	Header    string `json:"header,omitempty"`
}

func New(status, code, title, detail string) *Error {
	return &Error{
		Status: status,
		Code:   code,
		Title:  title,
		Detail: detail,
	}
}

func NewOpError(op string, err error) *Error {
	return &Error{
		Op:  op,
		Err: err,
	}
}

func (e *Error) WithSource(parameter, header string) *Error {
	clone := e.clone()
	clone.Source = &ErrorSource{
		Parameter: parameter,
		Header:    header,
	}
	return clone
}

func (e *Error) clone() *Error {
	clone := *e
	return &clone
}

func (e *Error) Error() string {
	var buf strings.Builder

	if e.Err != nil {

		if e.Op != "" {
			fmt.Fprintf(&buf, "%s: ", e.Op)
		}
		buf.WriteString(e.Err.Error())
	} else {
		fmt.Fprintf(
			&buf,
			`{"status":"%s","code":"%s","title":"%s","detail":"%s"`,
			e.Status, e.Code, e.Title, e.Detail,
		)
		if e.Source != nil {
			fmt.Fprintf(
				&buf,
				`,"source":{"parameter":"%s","header":"%s"}`,
				e.Source.Parameter, e.Source.Header,
			)
		}
		buf.WriteRune('}')
	}

	return buf.String()
}

func (e *Error) ErrorDetail() string {
	return errorDetail(e)
}

func errorDetail(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Detail != "" {
		return e.Detail
	} else if ok && e.Err != nil {
		return errorDetail(e.Err)
	} else if !ok {
		return err.Error()
	}
	return "An internal error has occurred."
}
