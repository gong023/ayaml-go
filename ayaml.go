package ayaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type fixtureFile map[string]map[string]interface{}

type Ayaml struct {
	Data *fixtureFile
}

func New(filename string) (ayaml Ayaml, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	var f fixtureFile
	err = yaml.Unmarshal(b, &f)
	if err != nil {
		return
	}
	ayaml.Data = &f
	return
}

func (a Ayaml) Scheme(s string) *Ayaml {
	return &a
}

func (a *Ayaml) With(data map[string]interface{}) *Ayaml {
	return a
}

func (*Ayaml) SetData(marshaler Marshaler) {
}
