package main

import (
	"encoding/xml"

	"github.com/sablev/go-xsd/xsd-makepkg/tests"

	"github.com/sablev/go-util/dev/go"

	svg "github.com/sablev/go-xsd-pkg/www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd_go"
)

type SvgDoc struct {
	XMLName xml.Name `xml:"svg"`
	svg.TsvgType
}

func main() {
	var (
		dirBasePath  = udevgo.GopathSrcGithub("sablev", "go-xsd", "xsd-makepkg", "tests", "xsd-test-svg")
		makeEmptyDoc = func() interface{} { return &SvgDoc{} }
	)
	tests.TestViaRemarshal(dirBasePath, makeEmptyDoc)
}
