package main

import (
	"reflect"
	"testing"
)

func TestDefaultFlags(t *testing.T) {
	if got, want := *filename, "urls.txt"; got != want {
		t.Errorf(`*filename = %q, want %q`, got, want)
	}

	if got, want := *showColor, true; got != want {
		t.Errorf(`*showColor = %v, want %v`, got, want)
	}

	if got, want := *fatalErrors, false; got != want {
		t.Errorf(`*fatalErrors = %v, want %v`, got, want)
	}

	if got, want := *limit, 100; got != want {
		t.Errorf(`*limit = %d, want %d`, got, want)
	}

	if got, want := *delay, 100; got != want {
		t.Errorf(`*delay = %d, want %d`, got, want)
	}

	if got, want := *baseURL, "http://0.0.0.0:7000"; got != want {
		t.Errorf(`*baseURL = %q, want %q`, got, want)
	}
}

func TestFatalMessage(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"red", "\x1b[0;31mERR\x1b[0m red"},
	}

	for _, test := range tests {
		if got := fatalMessage(test.in); got != test.want {
			t.Fatalf(`fatalMessage(%q) = %q, want %q`, test.in, got, test.want)
		}
	}
}

func TestReadURLs(t *testing.T) {
	file := "urls.txt"

	urls, _ := readURLs(&file)

	expected := []string{
		"http://c7.se",
		"http://c7.se/from-ruby-to-lua/",
		"http://example.org/",
		"https://humans.herokuapp.com/",
		"https://github.com/peterhellberg/warmup/",
		"http://example.com/error",
		"http://0.0.0.0:9912/no_server/",
		"http://0.0.0.0:7000/path/to/something",
	}

	if !reflect.DeepEqual(urls, expected) {
		t.Fatalf(`unexpected slice of URLs`)
	}
}

func TestColor(t *testing.T) {
	*showColor = false

	if got, want := color("1;36", "foo"), "foo"; got != want {
		t.Fatalf(`color("1;36", "foo") = %q, want %q`, got, want)
	}

	*showColor = true

	if got, want := color("1;36", "foo"), "\033[1;36mfoo\033[0m"; got != want {
		t.Fatalf(`color("1;36", "foo") = %q, want %q`, got, want)
	}
}

func TestRed(t *testing.T) {
	*showColor = false

	if got, want := red("foo"), "foo"; got != want {
		t.Fatalf(`red("foo") = %q, want %q`, got, want)
	}

	*showColor = true

	if got, want := red("foo"), "\033[0;31mfoo\033[0m"; got != want {
		t.Fatalf(`red("foo") = %q, want %q`, got, want)
	}
}

func TestGreen(t *testing.T) {
	*showColor = false

	if got, want := green("bar"), "bar"; got != want {
		t.Fatalf(`green("bar") = %q, want %q`, got, want)
	}

	*showColor = true

	if got, want := green("bar"), "\033[0;32mbar\033[0m"; got != want {
		t.Fatalf(`green("bar") = %q, want %q`, got, want)
	}
}

func TestBlue(t *testing.T) {
	*showColor = false

	if got, want := blue("baz"), "baz"; got != want {
		t.Fatalf(`blue("baz") = %q, want %q`, got, want)
	}

	*showColor = true

	if got, want := blue("baz"), "\033[0;34mbaz\033[0m"; got != want {
		t.Fatalf(`blue("baz") = %q, want %q`, got, want)
	}
}

func TestGray(t *testing.T) {
	*showColor = false

	if got, want := gray("qux"), "qux"; got != want {
		t.Fatalf(`gray("qux") = %q, want %q`, got, want)
	}

	*showColor = true

	if got, want := gray("qux"), "\033[1;30mqux\033[0m"; got != want {
		t.Fatalf(`gray("qux") = %q, want %q`, got, want)
	}
}
