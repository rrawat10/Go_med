package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A blank (or whitespace-only) line.
var ErrEmpty = errors.New("empty line")

// A non-blank line that does not have exactly three fields.
var ErrFields = errors.New("wrong field count")

// An age or score outside its allowed range.
type RangeError struct {
	Field    string
	Min, Max int
	Got      int
}

func (e RangeError) Error() string {
	return fmt.Sprintf("%s %d not in %d..%d", e.Field, e.Got, e.Min, e.Max)
}

// check validates one line "name age score" and returns the name, or the FIRST
// problem found: ErrEmpty for a blank line, ErrFields for a line without exactly
// three fields, then age (parse, then range 0..120), then score (parse, then
// range 0..100).
func check(line string) (string, error) {
	// TODO: use strings.Fields and strconv.Atoi; wrap parse failures with %w
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return "", ErrEmpty
	}
	if len(fields) != 3 {
		return "", ErrFields
	}

	age, err := strconv.Atoi(fields[1])
	if err != nil {
		return "", fmt.Errorf("age : %w", err)
	}
	if age < 0 || age > 120 {
		return "", RangeError{
			Field: "age",
			Min:   0,
			Max:   120,
			Got:   age,
		}
	}
	score, err := strconv.Atoi(fields[2])
	if err != nil {
		return "", fmt.Errorf("score : %w", err)
	}
	if score < 0 || score > 150 {
		return "", RangeError{
			Field: "score",
			Min:   0,
			Max:   150,
			Got:   score,
		}
	}
	return fields[0], nil

}

// validate adds the line number. Do not change it.
func validate(n int, line string) (string, error) {
	name, err := check(line)
	if err != nil {
		return "", fmt.Errorf("line %d: %w", n, err)
	}
	return name, nil
}

// classify turns an error from validate into the line to print:
// "empty", "fields", "parse: TEXT", or "range: ...".
func classify(err error) string {
	// TODO
	if errors.Is(err, ErrEmpty) {
		return "empty"
	}
	if errors.Is(err, ErrFields) {
		return "fields"
	}
	var numErr *strconv.NumError

	if errors.As(err, &numErr) {
		return "parse: " + numErr.Num
	}
	var rangeErr RangeError

	if errors.As(err, &rangeErr) {
		return "range: " + rangeErr.Error()
	}

	return "unknown"
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	n := 0
	for sc.Scan() {
		n++
		name, err := validate(n, sc.Text())
		if err != nil {
			fmt.Println(classify(err))
		} else {
			fmt.Println("ok " + name)
		}
	}
}
