package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"github.com/guitarpawat/the-internship-2019/weather/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var xmlFile, jsonFile string

func init() {
	flag.StringVar(&xmlFile, "xml", "weather.xml", "XML file location")
	flag.StringVar(&jsonFile, "json", "", "JSON file location, if not specified, output will use the same name as xml file")
	flag.Parse()
}

func main() {
	if jsonFile == "" {
		jsonFile = strings.TrimSuffix(xmlFile, filepath.Ext(xmlFile)) + ".json"
	}

	input, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		log.Fatalln("error while reading xml file:", err)
	}

	xmlStruct := model.CurrentXML{}
	err = xml.Unmarshal(input, &xmlStruct)
	if err != nil {
		log.Fatalln("error while unmarshalling xml:", err)
	}

	jsonStruct := xmlStruct.ToJSON()

	output, err := json.MarshalIndent(jsonStruct, "", "  ")
	if err != nil {
		log.Fatalln("error while marshalling json:", err)
	}

	err = ioutil.WriteFile(jsonFile, output, os.FileMode(os.O_CREATE|os.O_WRONLY))
	if err != nil {
		log.Fatalln("error while writing json file:", err)
	}
}
