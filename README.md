# gopher-get

A simple Gopher client to download a page/resource over the Gopher protocol, written in Go.

This client is adapted from the example client from the https://github.com/prologic/go-gopher package.

# Features

* Download a gopher page to a specified location
* Prints a description of the item type 
* Erm thats it

# Usage

```

gopher-get <url> <outputfile>

```

## Flags

```
-json    if content is a gophermap, output json
```

## Todo

Possible features to come:

* [ ] allow stream content to stdout
* [ ] add -o for output file
* [ ] maybe remove json option
* [ ] move display of item type to stderr
* [ ] add download abandoning if content exceeds a size/timeout threshold 