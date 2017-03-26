package ayaml

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const validYaml = "./test_fixture/a.yml"

func TestAyamlNew(t *testing.T) {
	_, err := New("fileNotExists")
	require.Error(t, err)

	a, err := New("./test_fixture/invalid.yml")
	require.Error(t, err)

	a, err = New(validYaml)
	require.NoError(t, err)
	fd := a.fileData
	assert.Equal(t, 100000000, fd["valid_user"]["id"])
	assert.Equal(t, "Taro", fd["valid_user"]["name"])
	assert.Equal(t, "2014-01-01 00:00:00", fd["valid_user"]["created"])
}

func TestSchema(t *testing.T) {
	a, _ := New(validYaml)
	a.Schema("valid_user")
	assert.Equal(t, "valid_user", a.schema)
}

func TestWith(t *testing.T) {
	a, _ := New(validYaml)

	a.Schema("valid_user").With(SchemaData{
		"id":      1,
		"name":    "Jiro",
		"created": "2015-01-01 00:00:00",
	})
	fd := a.fileData
	assert.Equal(t, 1, fd["valid_user"]["id"])
	assert.Equal(t, "Jiro", fd["valid_user"]["name"])
	assert.Equal(t, "2015-01-01 00:00:00", fd["valid_user"]["created"])
}

func TestDump(t *testing.T) {
	a, _ := New(validYaml)

	_, err := a.Dump()
	require.Error(t, err)

	_, err = a.Schema("invalid schema").Dump()
	require.Error(t, err)

	d, err := a.Schema("valid_user").Dump()
	require.NoError(t, err)
	assert.Equal(t, 100000000, d["id"])
	assert.Equal(t, "Taro", d["name"])
	assert.Equal(t, "2014-01-01 00:00:00", d["created"])
}

func TestWithDump(t *testing.T) {
	a, _ := New(validYaml)

	_, err := a.WithDump(SchemaData{})
	require.Error(t, err)

	d, err := a.Schema("valid_user").WithDump(SchemaData{"id": 1})
	require.NoError(t, err)
	assert.Equal(t, 1, d["id"])
}
