package ayaml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeIncrementByOne(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"id":   1,
					"name": "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	r := rangeIncrement{
		ayamlSeq: as,
		key:      "id",
		min:      10,
		max:      13,
	}

	as = r.ByOne()
	assert.Len(t, as.Results, 4)
	for i, r := range as.Results {
		assert.Equal(t, 10+i, r.fileData["valid_user"]["id"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}

	as = &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"id":   1,
					"name": "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	r = rangeIncrement{
		ayamlSeq: as,
		key:      "id",
		min:      10,
		max:      10,
	}
	as = r.ByOne()
	assert.Len(t, as.Results, 1)
	assert.Equal(t, 10, as.Results[0].fileData["valid_user"]["id"])
	assert.Equal(t, "UserName", as.Results[0].fileData["valid_user"]["name"])
}

func TestRangeDecrementByOne(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"id":   1,
					"name": "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	r := rangeDecrement{
		ayamlSeq: as,
		key:      "id",
		min:      10,
		max:      13,
	}

	as = r.ByOne()
	assert.Len(t, as.Results, 4)
	for i, r := range as.Results {
		assert.Equal(t, 13-i, r.fileData["valid_user"]["id"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}
