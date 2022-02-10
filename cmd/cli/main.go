package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"weekend-md/pkg/parse"
)

func main() {
	input := os.Stdin
	flag.Func("filepath", "`path` to file input (default STDIN)", func(s string) error {
		if s == " " {
			flag.PrintDefaults()
			return nil
		}
		f, err := os.Open(s)
		if err != nil {
			return err
		}

		input = f
		return nil
	})
	flag.Parse()

	prsr := parse.NewMdParser()
	buf := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(input)
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
	fmt.Println(buf.String())
}
