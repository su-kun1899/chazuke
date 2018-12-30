package chazuke

type Chazuke struct {
}

func (*Chazuke) Get(key string) *Chazuke {
	return &Chazuke{}
}

func (*Chazuke) Value() (string, error) {
	return "example", nil
}

func New(json string) *Chazuke {
	return &Chazuke{}
}
