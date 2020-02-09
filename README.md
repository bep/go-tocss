**Note:** You may want to use https://github.com/bep/golibsass instead of this.


[![Build Status](https://travis-ci.org/bep/go-tocss.svg?branch=master)](https://travis-ci.org/bep/go-tocss)
[![Go Report Card](https://goreportcard.com/badge/github.com/bep/go-tocss)](https://goreportcard.com/report/github.com/bep/go-tocss)

This is currently a, hopefully, simple to use [LibSass](https://sass-lang.com/libsass) Go API. It uses the C bindings in [https://github.com/wellington/go-libsass/libs](https://github.com/wellington/go-libsass/tree/master/libs) to do the heavy lifting.

The primary motivation for this project is to add `SCSS` support to [Hugo](https://gohugo.io/). It is has some generic `tocss` package names hoping that there will be a solid native Go implementation that can replace `LibSass` in the near future.
