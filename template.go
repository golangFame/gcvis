package main

import _ "embed"

//go:embed template.html
var GCVIS_TMPL string

//func init() {
//	GCVIS_TMPL = readTemplateFile()
//}
//
//func readTemplateFile() (htmlString string) {
//	filePath := "template.html"
//	htmlBytes, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		log.Fatal(fmt.Errorf("Failed to read HTML file: %v", err))
//		return
//	}
//
//	htmlString = string(htmlBytes)
//	return
//}
