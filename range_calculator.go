package ayaml

type RangeCalculator interface {
	ByOne() *AyamlSeq
	//By(fn func (int) int) *AyamlSeq
}

type rangeIncrement struct {
	ayamlSeq *AyamlSeq
	key      string
	min      int
	max      int
}

type rangeDecrement struct {
	ayamlSeq *AyamlSeq
	key      string
	min      int
	max      int
}

func (r *rangeIncrement) ByOne() *AyamlSeq {
	length := len(r.ayamlSeq.Results)
	index := 0
	for v := r.min; v <= r.max; v++ {
		if length > index {
			data := r.ayamlSeq.Results[index].withCopy(SchemaData{r.key: v})
			r.ayamlSeq.Results[index] = &data
		} else {
			data := r.ayamlSeq.Base.withCopy(SchemaData{r.key: v})
			r.ayamlSeq.Results = append(r.ayamlSeq.Results, &data)
		}
		index++
	}
	return r.ayamlSeq
}

func (r *rangeDecrement) ByOne() *AyamlSeq {
	length := len(r.ayamlSeq.Results)
	index := 0
	for v := r.max; v >= r.min; v-- {
		if length > index {
			data := r.ayamlSeq.Results[index].withCopy(SchemaData{r.key: v})
			r.ayamlSeq.Results[index] = &data
		} else {
			data := r.ayamlSeq.Base.withCopy(SchemaData{r.key: v})
			r.ayamlSeq.Results = append(r.ayamlSeq.Results, &data)
		}
		index++
	}
	return r.ayamlSeq
}
