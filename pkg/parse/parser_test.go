package parse

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

const EXPECTED_FULL = `<h1>Sample Document</h1>

<p>Hello!</p>

<p>This is sample markdown for the <a href="https://www.mailchimp.com">Mailchimp</a> homework assignment.</p>

<p><a href="https://www.mailchimp.com">Mailchimp</a></p>

<h1>Header one</h1>

<p>Hello there</p>

<p>How are you? What&#39;s going on?</p>

<h2>Another Header</h2>

<p>This is a paragraph <a href="http://google.com">with an inline link</a>. Neat, eh?</p>

<h2>This is a header <a href="http://yahoo.com">with a link</a></h2>
`

func Test_Single_Header(t *testing.T) {
	md := []byte(`# Sample Document`)
	expected := `<h1>Sample Document</h1>`
	prsr := NewMdParser()
	actual := fmt.Sprintf("%s", prsr.Parse(md))
	if actual != expected {
		t.Error(fmt.Sprintf("expected %v to be %v", expected, actual))
	}
}

func Test_Inline_Link(t *testing.T) {
	md := []byte(`[Mailchimp](https://www.mailchimp.com)`)
	expected := `<p><a href="https://www.mailchimp.com">Mailchimp</a></p>`
	prsr := NewMdParser()
	actual := fmt.Sprintf("%s", prsr.Parse(md))
	if actual != expected {
		t.Error(fmt.Sprintf("expected %v to be %v", expected, actual))
	}
}

func Test_Headers(t *testing.T) {
	md := []byte(`## This is a header`)
	expected := `<h2>This is a header</h2>`
	prsr := NewMdParser()
	actual := fmt.Sprintf("%s", prsr.Parse(md))
	if actual != expected {
		t.Error(fmt.Sprintf("expected %v to be %v", expected, actual))
	}
}

func TestMd_Full(t *testing.T) {
	// Read input from the user
	file, err := os.Open("test.md")
	if err != nil {
		log.Fatalf("error openning file: %v", err)
	}
	expected := EXPECTED_FULL
	prsr := NewMdParser()
	buf := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			buf.WriteByte('\n')
			continue

		}

		line = prsr.Parse(line)
		buf.Write(line)
		buf.WriteByte('\n')

	}
	actual := fmt.Sprintf("%s", buf.String())
	if actual != expected {
		t.Error(fmt.Sprintf("expected %v to be %v", expected, actual))
	}
}
