package ayaml

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
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
		assert.Equal(t, "2014-01-01T00:00:00Z", d["created"])
	}

	data, err = Seq(a).Range("id", 103, 100).ByOne().Dump()
	require.NoError(t, err)
	assert.Len(t, data, 4)
	for i, d := range data {
		assert.Equal(t, 103-i, d["id"])
		assert.Equal(t, "Taro", d["name"])
		assert.Equal(t, "2014-01-01T00:00:00Z", d["created"])
	}

	min := "2014-01-01T00:00:00Z"
	max := "2014-01-01T00:00:03Z"

	data, err = Seq(a).Between("created", time.RFC3339, min, max).BySecond().Dump()
	require.NoError(t, err)
	assert.Len(t, data, 4)
	for i, d := range data {
		assert.Equal(t, 100000000, d["id"])
		assert.Equal(t, "Taro", d["name"])
		assert.Equal(t, "2014-01-01T00:00:0"+strconv.Itoa(i)+"Z", d["created"])
	}

	data, err = Seq(a).Between("created", time.RFC3339, max, min).BySecond().Dump()
	require.NoError(t, err)
	assert.Len(t, data, 4)
	for i, d := range data {
		assert.Equal(t, 100000000, d["id"])
		assert.Equal(t, "Taro", d["name"])
		assert.Equal(t, "2014-01-01T00:00:0"+strconv.Itoa(3-i)+"Z", d["created"])
	}
}
