
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
    "os"


	"github.com/prologic/go-gopher"
)

var (
	json = flag.Bool("json", false, "display gopher directory as JSON")
)


func check(e error) {
    if e != nil {
        panic(e)
        	os.Exit(1)

    }
}

func fatal(format string, a ...interface{}) {
	
    format = "*** " + format + " ***\n"

	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)

}

func saveFile(contents []byte, path string) {
    
    //fmt.Print("write to "+ path)

    
    d1 := []byte(contents)
    err := ioutil.WriteFile(path, d1, 0644)
    
    
    check(err)
    
}

func main() {
	var uri string
    var out string

	flag.Parse()
    
    
    uri = flag.Arg(0)
    out = flag.Arg(1)

	res, err := gopher.Get(uri)
    
    fmt.Println(res.Type)
    
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		contents, err := ioutil.ReadAll(res.Body)
		if err != nil {
            fmt.Println(err)
			log.Fatal(err)
		}
        
        //save body to file
        saveFile(contents, out)
        
	} else {
		var (
			bytes []byte
			err   error
		)

		if *json {
			bytes, err = res.Dir.ToJSON()
		} else {
			bytes, err = res.Dir.ToText()
		}
		if err != nil {
			log.Fatal(err)
		}

        //save dir contents or other bytest to file
        saveFile(bytes, out)
	}
}
