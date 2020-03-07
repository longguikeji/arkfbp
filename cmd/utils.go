package cmd

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

func writeFile(dest string, assetName string, data interface{}) {
	t, err := Asset(assetName)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	tt := template.New("test")
	tt.Parse(string(t))
	tt.Execute(&b, data)

	if err := ioutil.WriteFile(dest, []byte(b.String()), 0644); err != nil {
		panic(err)
	}
}
