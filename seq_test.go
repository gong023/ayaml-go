package ayaml

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSeqDump(t *testing.T) {
	a, _ := New(validYaml)
	a.Schema("valid_user")

	data, err := Seq(a).Range("id", 100, 103).ByOne().Dump()
	require.NoError(t, err)
	assert.Len(t, data, 4)
	for i, d := range data {
		assert.Equal(t, 100+i, d["id"])
		assert.Equal(t, "Taro", d["name"])
		assert.Equal(t, "2014-01-01 00:00:00", d["created"])
	}

	data, err = Seq(a).Range("id", 103, 100).ByOne().Dump()
	require.NoError(t, err)
	assert.Len(t, data, 4)
	for i, d := range data {
		assert.Equal(t, 103-i, d["id"])
		assert.Equal(t, "Taro", d["name"])
		assert.Equal(t, "2014-01-01 00:00:00", d["created"])
	}
}
