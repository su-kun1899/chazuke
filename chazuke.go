package chazuke

type JSONContainer struct {
}

func (*JSONContainer) Get(key string) *JSONContainer {
	return &JSONContainer{}
}

func (*JSONContainer) Value() (string, error) {
	return "example", nil
}

func New(json string) *JSONContainer {
	return &JSONContainer{}
}
