package ayaml

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAyamlNew(t *testing.T) {
	_, err := New("fileNotExists")
	require.Error(t, err)

	y, err := New("./test_fixture/invalid.yml")
	require.Error(t, err)

	y, err = New("./test_fixture/a.yml")
	require.NoError(t, err)
	m := *y.Data
	assert.Equal(t, 100000000, m["valid_user"]["id"])
	assert.Equal(t, "Taro", m["valid_user"]["name"])
	assert.Equal(t, "2014-01-01 00:00:00", m["valid_user"]["created"])
}
