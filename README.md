Warmup
======

[![Build Status](https://travis-ci.org/peterhellberg/warmup.svg?branch=master)](https://travis-ci.org/peterhellberg/warmup)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/warmup)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/warmup#license-mit)

HTTP cache warming.

Go 1.1+ is required since `bufio.NewScanner` is used.

![Example output](http://assets.c7.se/skitch/warmup-20131219-032813.png)

### Installation

```bash
go get -u github.com/peterhellberg/warmup
```

### Usage of warmup

```
  -base-url="http://0.0.0.0:7000": The base URL used for paths
  -color=true: If you want the output in color
  -delay=100: Delay (in ms) between requests
  -fatal-errors=false: Useful for automated testing
  -filename="urls.txt": List of URLs
  -limit=100: Limit of concurrent requests
```

You probably want to warm the (soon to be) hottest routes on your site.

## License (MIT)

Copyright (c) 2013-2015 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
