package ayaml

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type Ayaml struct {
	Data map[string]interface{}
}

func (a *Ayaml) RegisterBasePath(path string) {
}

func (a Ayaml) Scheme(s string) Ayaml {
	return a
}

func (a *Ayaml) With(data map[string]interface{}) *Ayaml {
	return a
}

func (*Ayaml) SetData(marshaler Marshaler) {
}
