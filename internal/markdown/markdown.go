// Package markdown implements the functions, types, and interfaces for the module.
package markdown

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type Markdowns []Markdown

type Markdown struct {
	path    string
	relPath string
	level   int
	title   string
}

func (m *Markdown) Path() string {
	return m.path
}

func (m *Markdown) RelPath() string {
	return m.relPath
}

func (m *Markdown) Title() string {
	if m.title == "" {
		m.ReadTitle()
	}
	return m.title
}

func (m *Markdown) Level() int {
	m.level = len(strings.Split(m.path, "\\"))
	return m.level
}

func (m *Markdown) Rel(rel string) {
	m.relPath, _ = filepath.Rel(rel, m.path)
}

func (m *Markdown) ReadTitle() {
	m.title = extractTitle(m.path)
}

func extractTitle(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		return filepath.Base(filename)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}

	return filepath.Base(filename)
}

func New(path string) Markdown {
	return Markdown{
		path:  path,
		level: 0,
	}
}
