package parse

import (
	"bytes"
	"fmt"
	"html"
	"regexp"
)

var (
	h1Reg      = regexp.MustCompile(`^#(\s|)(.*?)$`)
	h2Reg      = regexp.MustCompile(`^##(\s|)(.*?)$`)
	h3Reg      = regexp.MustCompile(`^###(\s|)(.*?)$`)
	h4Reg      = regexp.MustCompile(`^####(\s|)(.*?)$`)
	h5Reg      = regexp.MustCompile(`^#####(\s|)(.*?)$`)
	h6Reg      = regexp.MustCompile(`^######(\s|)(.*?)$`)
	linkReg    = regexp.MustCompile(`\[(.+?)\]\(((?:https?:)?(\/|\\)(?:www\.)?.+?\.[a-z]+)\)`)
	lineBrkReg = regexp.MustCompile(`/\n/gm`)
)

// Option - TODO: an optimization would be to add functional options to the NewMdParser
type Option func(*md)

type md struct {
}

func (m md) Parse(input []byte) []byte {
	output := input
	output = m.escapeSpecialChars(output)
	output = m.headerParser(output)
	output = m.linkParser(output)
	output = m.inlineParser(output)
	return output
}

type MdParser interface {
	Parse(input []byte) []byte
}

// NewMdParser creates a &md which implements MdParser
func NewMdParser(options ...Option) MdParser {
	md := &md{}
	for _, opt := range options {
		opt(md)
	}
	return md
}

func (m *md) escapeSpecialChars(input []byte) []byte {
	res := html.EscapeString(string(input))
	return []byte(res)
}

func (m *md) linkParser(input []byte) []byte {
	res := linkReg.ReplaceAll(input, []byte(`<a href="$2">$1</a>`))
	return res
}

func (m *md) newLineToRrTransformer(input []byte) []byte {
	return lineBrkReg.ReplaceAll(input, []byte(`<br>`))
}

func (m *md) swapInline(input []byte) []byte {
	res := bytes.TrimSpace(input)
	if len(res) == 0 {
		return res
	}

	//check and ensure h tags are not added
	if bytes.Index(res, []byte(`<h`)) != -1 {
		return res
	}
	foo := m.newLineToRrTransformer(res)
	foo = []byte(fmt.Sprintf(`<p>%s</p>`, foo))
	return foo
}

func (m *md) inlineParser(input []byte) []byte {
	res := m.swapInline(input)
	return res
}

func (m *md) headerParser(input []byte) []byte {
	// convert headings based on # count
	if input[0] == '#' {

		count := bytes.Count(input, []byte(`#`))
		switch count {
		case 1:
			input = h1Reg.ReplaceAll(input, []byte(`<h1>$2</h1>`))
		case 2:
			input = h2Reg.ReplaceAll(input, []byte(`<h2>$2</h2>`))
		case 3:
			input = h3Reg.ReplaceAll(input, []byte(`<h3>$2</h3>`))
		case 4:
			input = h4Reg.ReplaceAll(input, []byte(`<h4>$2</h4>`))
		case 5:
			input = h5Reg.ReplaceAll(input, []byte(`<h5>$2</h5>`))
		case 6:
			input = h6Reg.ReplaceAll(input, []byte(`<h6>$2</h6>`))
		}
	}
	return input
}
