package ayaml

import (
	"time"
)

const (
	// doesn't care leap second & year for now
	Day  = time.Hour * 24
	Year = Day * 365
)

type DateCalculator interface {
	By(modify func(time time.Time) time.Time) *AyamlSeq
	BySecond() *AyamlSeq
	ByMinute() *AyamlSeq
	ByHour() *AyamlSeq
	ByDay() *AyamlSeq
	ByYear() *AyamlSeq
}

type dateIncrement struct {
	ayamlSeq *AyamlSeq
	key      string
	layout   string
	min      time.Time
	max      time.Time
}

type dateDecrement struct {
	ayamlSeq *AyamlSeq
	key      string
	layout   string
	min      time.Time
	max      time.Time
}

func (d *dateIncrement) By(modify func(minTime time.Time) time.Time) *AyamlSeq {
	for {
		data := d.ayamlSeq.Base.withCopy(SchemaData{d.key: d.min.Format(d.layout)})
		d.ayamlSeq.Results = append(d.ayamlSeq.Results, &data)
		d.min = modify(d.min)
		if d.min.After(d.max) {
			break
		}
	}

	return d.ayamlSeq
}

func (d *dateDecrement) By(modify func(maxTime time.Time) time.Time) *AyamlSeq {
	for {
		data := d.ayamlSeq.Base.withCopy(SchemaData{d.key: d.max.Format(d.layout)})
		d.ayamlSeq.Results = append(d.ayamlSeq.Results, &data)
		d.max = modify(d.max)
		if d.max.Before(d.min) {
			break
		}
	}

	return d.ayamlSeq
}

func (d *dateIncrement) BySecond() *AyamlSeq {
	return d.By(func(minTime time.Time) time.Time {
		return minTime.Add(time.Second)
	})
}

func (d *dateDecrement) BySecond() *AyamlSeq {
	return d.By(func(maxTime time.Time) time.Time {
		return maxTime.Add(-time.Second)
	})
}

func (d *dateIncrement) ByMinute() *AyamlSeq {
	return d.By(func(minTime time.Time) time.Time {
		return minTime.Add(time.Minute)
	})
}

func (d *dateDecrement) ByMinute() *AyamlSeq {
	return d.By(func(maxTime time.Time) time.Time {
		return maxTime.Add(-time.Minute)
	})
}

func (d *dateIncrement) ByHour() *AyamlSeq {
	return d.By(func(minTime time.Time) time.Time {
		return minTime.Add(time.Hour)
	})
}

func (d *dateDecrement) ByHour() *AyamlSeq {
	return d.By(func(maxTime time.Time) time.Time {
		return maxTime.Add(-time.Hour)
	})
}

func (d *dateIncrement) ByDay() *AyamlSeq {
	return d.By(func(minTime time.Time) time.Time {
		return minTime.Add(Day)
	})
}

func (d *dateDecrement) ByDay() *AyamlSeq {
	return d.By(func(maxTime time.Time) time.Time {
		return maxTime.Add(-Day)
	})
}

func (d *dateIncrement) ByYear() *AyamlSeq {
	return d.By(func(minTime time.Time) time.Time {
		year := minTime.Year()
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return minTime.Add(Year + Day)
		}
		return minTime.Add(Year)
	})
}

func (d *dateDecrement) ByYear() *AyamlSeq {
	return d.By(func(maxTime time.Time) time.Time {
		maxTime = maxTime.Add(-Year)
		year := maxTime.Year()
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return maxTime.Add(-Day)
		}
		return maxTime
	})
}
