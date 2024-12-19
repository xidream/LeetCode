package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	xflag "github.com/xidream/go/flag"
)

const (
	example1    = "**Example 1:**"
	constraints = "**Constraints:**"
	followUp    = `**Follow\-up:**`

	leetCodeMarkdownBadge = `<a href="https://github.com/whateverzpy/LeetCode-Markdown"><img src="https://img.shields.io/badge/Markdown-FFA116?logo=leetcode&labelColor=555"/></a>`
	br                    = "<br>"

	easy   = "easy"
	medium = "medium"
	hard   = "hard"
)

var (
	difficulties = []string{
		easy,
		medium,
		hard,
	}
)

func main() {
	var (
		path       = flag.String("path", "", "the source markdown file to prettify")
		number     = flag.Int("number", 0, "number of the problem")
		difficulty = flag.String("difficulty", "", fmt.Sprintf("difficulty of the problem, one of %v", difficulties))
		topics     = flag.String("topics", "", "topics of the problem")
		write      = flag.Bool("write", false, "write prettified markdown to the source file")
	)
	flag.Parse()

	b, err := os.ReadFile(*path)
	if err != nil {
		xflag.Fatal(err)
	}

	if *number == 0 {
		xflag.Fatal("missing -number")
	}

	if !slices.Contains(difficulties, *difficulty) {
		xflag.Fatal("invalid -difficulty: ", *difficulty)
	}

	var (
		lines      = strings.Split(string(b), "\n")
		prettified = []string{
			numberTitle(lines[0], *number),
			"",
			leetCodeMarkdownBadge,
			newDifficultyBadge(*difficulty),
			newTopicsBadge(*topics),
			"",
		}
	)

	for _, line := range lines[2:] {
		switch {
		case isQuote(line):
			prettified = append(prettified, appendBRTag(line))
		case isEmptyQuote(line):
			continue
		case isNewSection(line):
			prettified = append(
				prettified,
				br,
				"",
				line,
			)
		default:
			prettified = append(prettified, line)
		}
	}

	output := strings.Join(prettified, "\n")
	if *write {
		os.WriteFile(*path, []byte(output), 0644)
		return
	}

	logrus.Info(output)
}

func numberTitle(line string, number int) string {
	return "# " + strconv.Itoa(number) + ". " + line[2:]
}

func newDifficultyBadge(difficulty string) string {
	switch difficulty {
	case easy:
		return `![](https://img.shields.io/badge/Difficulty-Easy-green)`
	case medium:
		return `![](https://img.shields.io/badge/Difficulty-Medium-orange)`
	case hard:
		return `![](https://img.shields.io/badge/Difficulty-Hard-red)`
	}
	return ""
}

func newTopicsBadge(topics string) string {
	topics = strings.ReplaceAll(topics, " ", "_")

	return `![](https://img.shields.io/badge/Topics-` + topics + `-blue)`
}

func isQuote(line string) bool {
	return strings.HasPrefix(line, "> ")
}

func appendBRTag(line string) string {
	return line + " " + br
}

func isEmptyQuote(line string) bool {
	return line == ">"
}

func isNewSection(line string) bool {
	return line == example1 ||
		line == constraints ||
		strings.HasPrefix(line, followUp)
}
