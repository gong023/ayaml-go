package ayaml

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SchemaData map[string]interface{}
type fileData map[string]SchemaData

type Ayaml struct {
	fileData fileData
	schema   string
}

func New(filename string) (ayaml Ayaml, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	var f fileData
	err = yaml.Unmarshal(b, &f)
	if err != nil {
		return
	}
	ayaml.fileData = f
	return
}

func (a *Ayaml) Schema(key string) *Ayaml {
	a.schema = key
	return a
}

func (a *Ayaml) With(newData SchemaData) *Ayaml {
	for k, v := range newData {
		if _, ok := a.fileData[a.schema][k]; ok {
			a.fileData[a.schema][k] = v
		}
	}
	return a
}

func (a *Ayaml) Dump() (SchemaData, error) {
	if a.schema == "" {
		return nil, errors.New("schema should be set before dump")
	}
	if _, ok := a.fileData[a.schema]; ok {
		return a.fileData[a.schema], nil
	} else {
		return nil, errors.New(a.schema + " is not found in filedata")
	}
}

func (a *Ayaml) WithDump(newData SchemaData) (SchemaData, error) {
	return a.With(newData).Dump()
}
