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
	for i := r.min; i <= r.max; i++ {
		data := r.ayamlSeq.Base.withCopy(SchemaData{r.key: i})
		r.ayamlSeq.Results = append(r.ayamlSeq.Results, &data)
	}
	return r.ayamlSeq
}

func (r *rangeDecrement) ByOne() *AyamlSeq {
	for i := r.max; i >= r.min; i-- {
		data := r.ayamlSeq.Base.withCopy(SchemaData{r.key: i})
		r.ayamlSeq.Results = append(r.ayamlSeq.Results, &data)
	}
	return r.ayamlSeq
}
