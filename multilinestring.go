package geo

type MultiLineString []LineString

func NewMultiLineString(ls ...LineString) *MultiLineString {
	var ml MultiLineString
	for _, v := range ls {
		ml = append(ml, v)
	}
	return &ml
}

func (l MultiLineString) SetSrid(srid uint64) {
	SridMap[&l] = srid
}

func (l MultiLineString) ToWkt() string {
	return LineToWkt(l...)
}

func (ml MultiLineString) BoundingBox() Box {
	var pois []Point
	for _, v := range ml {
		for _, vv := range v {
			pois = append(pois, vv)
		}
	}
	return calBox(pois...)
}

func (ml MultiLineString) TypeString() string {
	return "MultiLineString"
}
