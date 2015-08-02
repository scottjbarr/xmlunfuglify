# XML Unfugly

A small Go program that unfuglies (or prettifies) XML form stdin.

The solution was found on Stackoverflow, and I've just packaged it up.

## Build

    go install github.com/scottjbarr/xmlunfugly

## Usage

Command line program example.

    cat ugly.xml | xmlunfugly > unfugly.xml

## References

- [Go: How would you “Pretty Print”/“Prettify” HTML?](https://stackoverflow.com/questions/21117161/go-how-would-you-pretty-print-prettify-html)

## Licence

The MIT License (MIT)

Copyright (c) 2015 Scott Barr

See [LICENSE.md](LICENSE.md)
