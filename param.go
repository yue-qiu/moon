package moon

type Params map[string]string

func NewParams() Params {
	return make(map[string]string)
}

func (p Params) Count() (count int) {
	for _, _ = range p {
		count++
	}
	return count
}

func (p Params) Get(key string) (val string) {
	return p[key]
}
