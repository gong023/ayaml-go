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
	min, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:03Z")
	d := dateIncrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.BySecond()
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
	min, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:03Z")
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.BySecond()
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
	min, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2014-01-01T00:03:00Z")
	d := dateIncrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.ByMinute()
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
	min, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2014-01-01T00:03:00Z")
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.ByMinute()
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
	min, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2014-01-01T03:00:00Z")
	d := dateIncrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.ByHour()
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
	min, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2014-01-01T03:00:00Z")
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.ByHour()
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
	min, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2014-01-03T00:00:00Z")
	d := dateIncrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.ByDay()
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
	min, _ := time.Parse(time.RFC3339, "2014-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2014-01-03T00:00:00Z")
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.ByDay()
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
	min, _ := time.Parse(time.RFC3339, "2010-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2013-01-01T00:00:00Z")
	d := dateIncrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.ByYear()
	assert.Len(t, d.ayamlSeq.Results, 4)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "201"+strconv.Itoa(i)+"-01-01T00:00:00Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}

func TestDateDecrementByYear(t *testing.T) {
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
	min, _ := time.Parse(time.RFC3339, "2010-01-01T00:00:00Z")
	max, _ := time.Parse(time.RFC3339, "2013-01-01T00:00:00Z")
	d := dateDecrement{
		ayamlSeq: as,
		key:      "created",
		layout:   time.RFC3339,
		min:      min,
		max:      max,
	}

	d.ByYear()
	assert.Len(t, d.ayamlSeq.Results, 4)
	for i, r := range d.ayamlSeq.Results {
		assert.Equal(t, "201"+strconv.Itoa(3-i)+"-01-01T00:00:00Z", r.fileData["valid_user"]["created"])
		assert.Equal(t, "UserName", r.fileData["valid_user"]["name"])
	}
}
