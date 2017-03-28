package ayaml

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestDateIncrementBySecond(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"created": "2014-01-01T00:00:00Z",
					"name":    "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	d := dateIncrement{
		ayamlSeq: as,
		key:      "created",
		min:      "2014-01-01T00:00:00Z",
		max:      "2014-01-01T00:00:03Z",
	}

	d.BySecond(time.RFC3339)
	assert.Len(t, d.ayamlSeq.Results, 4)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "2014-01-01T00:00:0"+strconv.Itoa(i)+"Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}

func TestDateDecrementBySecond(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"created": "2014-01-01T00:00:00Z",
					"name":    "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		min:      "2014-01-01T00:00:00Z",
		max:      "2014-01-01T00:00:03Z",
	}

	d.BySecond(time.RFC3339)
	assert.Len(t, d.ayamlSeq.Results, 4)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "2014-01-01T00:00:0"+strconv.Itoa(3-i)+"Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}

func TestDateIncrementByMinute(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"created": "2014-01-01T00:00:00Z",
					"name":    "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	d := dateIncrement{
		ayamlSeq: as,
		key:      "created",
		min:      "2014-01-01T00:00:00Z",
		max:      "2014-01-01T00:03:00Z",
	}

	d.ByMinute(time.RFC3339)
	assert.Len(t, d.ayamlSeq.Results, 4)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "2014-01-01T00:0"+strconv.Itoa(i)+":00Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}

func TestDateDecrementByMinute(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"created": "2014-01-01T00:00:00Z",
					"name":    "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		min:      "2014-01-01T00:00:00Z",
		max:      "2014-01-01T00:03:00Z",
	}

	d.ByMinute(time.RFC3339)
	assert.Len(t, d.ayamlSeq.Results, 4)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "2014-01-01T00:0"+strconv.Itoa(3-i)+":00Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}

func TestDateIncrementByHour(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"created": "2014-01-01T00:00:00Z",
					"name":    "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	d := dateIncrement{
		ayamlSeq: as,
		key:      "created",
		min:      "2014-01-01T00:00:00Z",
		max:      "2014-01-01T03:00:00Z",
	}

	d.ByHour(time.RFC3339)
	assert.Len(t, d.ayamlSeq.Results, 4)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "2014-01-01T0"+strconv.Itoa(i)+":00:00Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}

func TestDateDecrementByHour(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"created": "2014-01-01T00:00:00Z",
					"name":    "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		min:      "2014-01-01T00:00:00Z",
		max:      "2014-01-01T03:00:00Z",
	}

	d.ByHour(time.RFC3339)
	assert.Len(t, d.ayamlSeq.Results, 4)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "2014-01-01T0"+strconv.Itoa(3-i)+":00:00Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}

func TestDateIncrementByDay(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"created": "2014-01-01T00:00:00Z",
					"name":    "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	d := dateIncrement{
		ayamlSeq: as,
		key:      "created",
		min:      "2014-01-01T00:00:00Z",
		max:      "2014-01-03T00:00:00Z",
	}

	d.ByDay(time.RFC3339)
	assert.Len(t, d.ayamlSeq.Results, 3)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "2014-01-0"+strconv.Itoa(1+i)+"T00:00:00Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}

func TestDateDecrementByDay(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"created": "2014-01-01T00:00:00Z",
					"name":    "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		min:      "2014-01-01T00:00:00Z",
		max:      "2014-01-03T00:00:00Z",
	}

	d.ByDay(time.RFC3339)
	assert.Len(t, d.ayamlSeq.Results, 3)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "2014-01-0"+strconv.Itoa(3-i)+"T00:00:00Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}

func TestDateIncrementByYear(t *testing.T) {
	as := &AyamlSeq{
		Base: &Ayaml{
			fileData: fileData{
				"valid_user": SchemaData{
					"created": "2014-01-01T00:00:00Z",
					"name":    "UserName",
				},
			},
			schema: "valid_user",
		},
	}
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		min:      "2010-01-01T00:00:00Z",
		max:      "2013-01-01T00:00:00Z",
	}

	d.ByYear(time.RFC3339)
	assert.Len(t, d.ayamlSeq.Results, 4)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "201"+strconv.Itoa(3-i)+"-01-01T00:00:00Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}
