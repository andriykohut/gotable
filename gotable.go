// Package tablefmt provides tools to convert slices of maps to ASCII tables
package gotable

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

// ANSI escape sequences for displaying bold text
const (
	boldStart = "\033[1m"
	boldEnd   = "\033[0m"
)

type Table struct {
	Headers    []string
	Rows       []map[string]string
	BoldHeader bool
	Separators []rune
	widths     map[string]int
}

/*
 * Calculates maximum column widths for each map in t.Rows. Returns
 * map[string]int where keys are column names and values are maximum
 * length row item.
 */
func (t *Table) widthForCols() map[string]int {
	if t.widths == nil {
		widths := make(map[string]int)
		for _, item := range t.Rows {
			for key, val := range item {
				_, ok := widths[key]
				if !ok {
					widths[key] = utf8.RuneCountInString(key)
				}
				valLen := utf8.RuneCountInString(val)
				if valLen > widths[key] {
					widths[key] = valLen
				}
			}
		}
		t.widths = widths
	}
	return t.widths
}

/*
 * Create line string from widths map and list of headers to separate each
 * row in a table.
 */
func (t *Table) line() string {
	var line bytes.Buffer
	line.WriteRune(t.Separators[0])
	for _, header := range t.Headers {
		line.WriteString(strings.Repeat(string(t.Separators[1]), t.widthForCols()[header]+2))
		line.WriteRune(t.Separators[0])
	}
	return line.String()
}

// Make text bold :)
func bold(text string) string {
	return boldStart + text + boldEnd
}

// Generate table string from slice of maps and list of headers.
func (t *Table) GetTable() string {
	widths := t.widthForCols()
	var table bytes.Buffer
	line := t.line()
	table.WriteString(line)
	table.WriteRune('\n')
	for _, header := range t.Headers {
		table.WriteRune(t.Separators[2])
		table.WriteRune(' ')
		if t.BoldHeader {
			table.WriteString(bold(header))
		} else {
			table.WriteString(header)
		}
		gap := widths[header] - utf8.RuneCountInString(header)
		if gap > 0 {
			table.WriteString(strings.Repeat(" ", gap))
		}
		table.WriteRune(' ')
	}
	table.WriteRune(t.Separators[2])
	table.WriteRune('\n')
	table.WriteString(line)
	for _, item := range t.Rows {
		table.WriteRune('\n')
		table.WriteRune(t.Separators[2])
		for _, header := range t.Headers {
			table.WriteRune(' ')
			table.WriteString(item[header])
			gap := widths[header] - utf8.RuneCountInString(item[header])
			if gap > 0 {
				table.WriteString(strings.Repeat(" ", gap))
			}
			table.WriteRune(' ')
			table.WriteRune(t.Separators[2])
		}
		table.WriteRune('\n')
		table.WriteString(line)
	}
	return table.String()
}

/*
 * Create new table struct.
 * maps - table rows, can take additional arguments: []string for table headers,
 * []rune for separators, bool to make table headers bold.
 */
func NewTable(maps []map[string]string, args ...interface{}) *Table {
	table := &Table{}
	table.Rows = maps
	for _, arg := range args {
		switch arg.(type) {
		case []string:
			table.Headers = arg.([]string)
		case []rune:
			table.Separators = arg.([]rune)
		case bool:
			table.BoldHeader = arg.(bool)
		}
	}
	if table.Headers == nil {
		var headers []string
		for k, _ := range maps[0] {
			headers = append(headers, k)
		}
		table.Headers = headers
	}
	if table.Separators == nil {
		table.Separators = []rune{'+', '-', '|'}
	}
	return table
}
