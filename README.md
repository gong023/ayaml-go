## ayaml-go

![build](https://circleci.com/gh/gong023/ayaml-go.svg?style=shield)

ayaml-go is imported from [Ayaml](https://github.com/gong023/Ayaml)

ayaml-go gives you factory-girl like DSL to build test data.

Note that ayaml-go doesn't build test data as it is. Connection to database should be delegated to the other codes. (this is why this library is simple and flexible


## sample

build data from below yml

```yaml
valid_user:
  id: 100000000
  name: UserName
  created: '2014-01-01T00:00:00Z'
```

```go
a, _ := ayaml.New("./your_fixture.yml")


data, _ := a.Schema("valid_user").Dump()
data == SchemaData{
	"id":      100000000,
	"name":    "UserName",
	"created": "2014-01-01T00:00:00Z",
}

data, _ = a.Schema("valid_user").With(SchemaData{"id": 1}).Dump()
data == SchemaData{
	"id":      1,
	"name":    "UserName",
	"created": "2014-01-01T00:00:00Z",
}
```

build sequential data

```go
a, err := ayaml.New("./your_fixture.yml")
if err != nil {
	t.FailNow()
}

user := a.Schema("valid_user")

// increment id.
data, _ := ayaml.Seq(user).Range("id", 10, 12).ByOne().Dump()
data == []SchemaData{
	{
		"id": 10,
		"name": "UserName",
		"created": "2014-01-01T00:00:00Z",
	},
	{
		"id": 11,
		"name": "UserName",
		"created": "2014-01-01T00:00:00Z",
	},
	{
		"id": 12,
		"name": "UserName",
		"created": "2014-01-01T00:00:00Z",
	},
}

// decrement id.
data, _ = ayaml.Seq(user).Range("id", 10, 8).ByOne().Dump()
data == []SchemaData{
	{
		"id": 10,
		"name": "UserName",
		"created": "2014-01-01T00:00:00Z",
	},
	{
		"id": 9,
		"name": "UserName",
		"created": "2014-01-01T00:00:00Z",
	},
	{
		"id": 8,
		"name": "UserName",
		"created": "2014-01-01T00:00:00Z",
	},
}

// increment date string.
// you can specify duration 'byDay','byWeek','byMonth','byYear','bySecond'
data, _ = ayaml.Seq(user).Between("created", time.RFC3339, "2014-01-01T00:00:00Z", "2014-03-01T00:00:00Z").ByMonth().Dump()
data == []SchemaData{
	{
		"id": 100000000,
		"name": "UserName",
		"created": "2014-01-01T00:00:00Z",
	},
	{
		"id": 100000000,
		"name": "UserName",
		"created": "2014-02-01T00:00:00Z",
	},
	{
		"id": 100000000,
		"name": "UserName",
		"created": "2014-03-01T00:00:00Z",
	},
}

// decrement date string.
data, _ = ayaml.Seq(user).Between("created", time.RFC3339, "2014-03-01T00:00:00Z", "2014-01-01T00:00:00Z").ByMonth().Dump()
data == []SchemaData{
	{
		"id": 100000000,
		"name": "UserName",
		"created": "2014-03-01T00:00:00Z",
	},
	{
		"id": 100000000,
		"name": "UserName",
		"created": "2014-02-01T00:00:00Z",
	},
	{
		"id": 100000000,
		"name": "UserName",
		"created": "2014-01-01T00:00:00Z",
	},
}

// make numeric and date column sequential.
d := ayaml.Seq(user).Range("id", 10, 12).ByOne()
d = d.Between("created", time.RFC3339, "2014-01-01T00:00:00Z", "2014-03-01T00:00:00Z").ByMonth()
data := d.Dump()
data == []SchemaData{
	{
		"id": 10,
		"name": "UserName",
		"created": "2014-01-01T00:00:00Z",
	},
	{
		"id": 11,
		"name": "UserName",
		"created": "2014-02-01T00:00:00Z",
	},
	{
		"id": 13,
		"name": "UserName",
		"created": "2014-03-01T00:00:00Z",
	},
}
```
