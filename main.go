package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Println(err)
		return
	}

	var output bytes.Buffer
	err = json.Indent(&output, input, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = os.Stdout.Write(output.Bytes())
	if err != nil {
		log.Println(err)
		return
	}
}
