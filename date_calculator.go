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
	By(layout string, duration time.Duration) *AyamlSeq
	BySecond(layout string) *AyamlSeq
	ByMinute(layout string) *AyamlSeq
	ByHour(layout string) *AyamlSeq
	ByDay(layout string) *AyamlSeq
	ByYear(layout string) *AyamlSeq
}

type dateIncrement struct {
	ayamlSeq *AyamlSeq
	key      string
	min      string
	max      string
}

type dateDecrement struct {
	ayamlSeq *AyamlSeq
	key      string
	min      string
	max      string
}

func (d *dateIncrement) By(layout string, modify func(minTime time.Time) time.Time) *AyamlSeq {
	// note: err is ignored here
	minTime, _ := time.Parse(layout, d.min)
	maxTime, _ := time.Parse(layout, d.max)

	for {
		data := d.ayamlSeq.Base.withCopy(SchemaData{d.key: minTime.Format(layout)})
		d.ayamlSeq.Results = append(d.ayamlSeq.Results, &data)
		minTime = modify(minTime)
		if minTime.After(maxTime) {
			break
		}
	}

	return d.ayamlSeq
}

func (d *dateIncrement) BySecond(layout string) *AyamlSeq {
	return d.By(layout, func(minTime time.Time) time.Time {
		return minTime.Add(time.Second)
	})
}

func (d *dateIncrement) ByMinute(layout string) *AyamlSeq {
	return d.By(layout, func(minTime time.Time) time.Time {
		return minTime.Add(time.Minute)
	})
}

func (d *dateIncrement) ByHour(layout string) *AyamlSeq {
	return d.By(layout, func(minTime time.Time) time.Time {
		return minTime.Add(time.Hour)
	})
}

func (d *dateIncrement) ByDay(layout string) *AyamlSeq {
	return d.By(layout, func(minTime time.Time) time.Time {
		return minTime.Add(Day)
	})
}

func (d *dateIncrement) ByYear(layout string) *AyamlSeq {
	return d.By(layout, func(minTime time.Time) time.Time {
		year := minTime.Year()
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return minTime.Add(Year + Day)
		}
		return minTime.Add(Year)
	})
}
