package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/iancoleman/strcase"
)

// Reads all .json files in the current folder
// and encodes them as strings literals in proto/swagger.pb.go
func main() {
	var (
		dir     = "proto"
		swagger = path.Join(dir, "swagger.pb.go")
	)
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	out, err := os.Create(swagger)
	if err != nil {
		panic(err)
	}

	out.Write([]byte("package proto \n\nconst (\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".json") {
			name := strings.Trim(f.Name(), ".json")
			name = strings.Replace(name, ".", "", -1)
			name = strings.Replace(name, "swagger", "Swagger", -1)
			name = strcase.ToCamel(name)
			out.Write([]byte(name + " = `"))
			f, err := os.Open(path.Join(dir, f.Name()))
			if err != nil {
				panic(err)
			}
			data, err := ioutil.ReadAll(f)
			if err != nil {
				panic(err)
			}
			content := strings.Replace(string(data), "`", `"`, -1)
			out.WriteString(content)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}
