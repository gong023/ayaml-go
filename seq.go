package ayaml

type AyamlSeq struct {
	Base    *Ayaml
	Results []*Ayaml
}

func (as *AyamlSeq) Range(key string, start int, end int) RangeCalculator {
	if start <= end {
		increment := rangeIncrement{
			ayamlSeq: as,
			key:      key,
			min:      start,
			max:      end,
		}
		return RangeCalculator(&increment)
	}

	decrement := rangeDecrement{
		ayamlSeq: as,
		key:      key,
		min:      end,
		max:      start,
	}
	return RangeCalculator(&decrement)
}

func (as *AyamlSeq) Dump() ([]SchemaData, error) {
	var data []SchemaData
	for _, result := range as.Results {
		r, err := result.Dump()
		if err != nil {
			return nil, err
		}
		data = append(data, r)
	}
	return data, nil
}
