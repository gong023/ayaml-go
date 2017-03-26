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
	AyamlSeq *AyamlSeq
	Key      string
	Min      int
	Max      int
}

func (r *rangeIncrement) ByOne() *AyamlSeq {
	for i := r.min; i <= r.max; i++ {
		data := r.ayamlSeq.Base.With(SchemaData{r.key: i})
		r.ayamlSeq.Results = append(r.ayamlSeq.Results, data)
	}
	return r.ayamlSeq
}

func (r *rangeDecrement) ByOne() *AyamlSeq {
	for i := r.Max; i >= r.Min; i-- {
		data := r.AyamlSeq.Base.With(SchemaData{r.Key: i})
		r.AyamlSeq.Results = append(r.AyamlSeq.Results, data)
	}
	return r.AyamlSeq
}
