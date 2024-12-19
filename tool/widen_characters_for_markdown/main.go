package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	xflag "github.com/xidream/go/flag"
	"golang.org/x/text/width"
)

const (
	TypeBlockquote = "blockquote"
)

type Range struct {
	Start int
	End   int
}

func NewRange(s string) (*Range, error) {
	rv := &Range{}
	if _, err := fmt.Sscanf(s, "%d-%d", &rv.Start, &rv.End); err != nil {
		return nil, err
	}
	if rv.Start < 1 || rv.End < 1 || rv.Start > rv.End {
		return nil, fmt.Errorf("invalid range: %v", rv)
	}
	return rv, nil
}

func (r *Range) String() string {
	return fmt.Sprintf("%+v", *r)
}

func main() {
	var (
		path  = flag.String("path", "", "the source markdown file to widen characters")
		typ   = flag.String("type", "", "the markdown element type to widen, one of [blockquote]")
		rang  = flag.String("range", "", "the range of the line numbers to widen characters, to specify line 3 to line 6 and line 9 to line 12: 3-6,9-12")
		write = flag.Bool("write", false, "write the result to the source file")
	)
	flag.Parse()

	b, err := os.ReadFile(*path)
	if err != nil {
		xflag.Fatal(err)
	}

	var widen func(string) (string, error)
	switch *typ {
	case TypeBlockquote:
		widen = widenBlockquote
	default:
		xflag.Fatal("invalid -type: ", *typ)
	}

	var (
		ss     = strings.Split(*rang, ",")
		ranges = make([]*Range, len(ss))
	)
	for i, s := range ss {
		ranges[i], err = NewRange(s)
		if err != nil {
			xflag.Fatal(err)
		}
	}

	var (
		lines   = strings.Split(string(b), "\n")
		widened = make([]string, len(lines))
	)
	for _, r := range ranges {
		for i := r.Start - 1; i < r.End; i++ {
			if widened[i], err = widen(lines[i]); err != nil {
				xflag.Fatal(err)
			}
		}
	}
	for i, line := range widened {
		if line == "" {
			widened[i] = lines[i]
		}
	}

	output := strings.Join(widened, "\n")
	if *write {
		os.WriteFile(*path, []byte(output), 0644)
		return
	}

	logrus.Info(output)
}

func widenBlockquote(line string) (string, error) {
	var (
		prefix = "> "
		suffix = " <br>"
	)

	if !strings.HasPrefix(line, prefix) || !strings.HasSuffix(line, suffix) {
		return "", fmt.Errorf("not a blockquote line: %v", line)
	}

	trimmed := line[len(prefix) : len(line)-len(suffix)]
	return prefix + width.Widen.String(trimmed) + suffix, nil
}
