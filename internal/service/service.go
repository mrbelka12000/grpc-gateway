package service

type implementer struct{}

func New() *implementer {
	return &implementer{}
}

func (i *implementer) GetInfoByIIN(iin string) {
}
