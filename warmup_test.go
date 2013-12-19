package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWarmup(t *testing.T) {
	Convey("Default flags", t, func() {
		So(*filename, ShouldEqual, "urls.txt")
		So(*showColor, ShouldBeTrue)
		So(*limit, ShouldEqual, 100)
	})

	Convey("fatalMessage()", t, func() {
		Convey("returns a string prepended with (red) ERROR:", func() {
			So(fatalMessage("foo"), ShouldEqual, "\033[0;31mERROR:\033[0m foo")
		})
	})

	Convey("readURLs()", t, func() {
		Convey("it reads the example urls.txt", func() {
			file := "urls.txt"

			urls, _ := readURLs(&file)

			expected := []string{
				"http://c7.se",
				"http://c7.se/from-ruby-to-lua/",
				"http://example.org/",
				"http://example.com/error",
				"https://humans.herokuapp.com/",
				"https://github.com/peterhellberg/warmup/",
			}

			So(urls, ShouldResemble, expected)
		})
	})

	Convey("color()", t, func() {
		Convey("makes text colorful", func() {
			*showColor = true
			So(color("1;36", "foo"), ShouldEqual, "\033[1;36mfoo\033[0m")
		})

		Convey("that is unless color has been disabled", func() {
			*showColor = false
			So(color("1;36", "foo"), ShouldEqual, "foo")
		})

		*showColor = true

		Convey("red()   turns text… red", func() {
			So(red("foo"), ShouldEqual, "\033[0;31mfoo\033[0m")
		})

		Convey("green() turns text… green", func() {
			So(green("bar"), ShouldEqual, "\033[0;32mbar\033[0m")
		})

		Convey("blue()  turns text… blue", func() {
			So(blue("baz"), ShouldEqual, "\033[0;34mbaz\033[0m")
		})

		Convey("gray()  turns text… gray", func() {
			So(gray("baz"), ShouldEqual, "\033[1;30mbaz\033[0m")
		})
	})
}
