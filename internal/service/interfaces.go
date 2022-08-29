package service

type (
	Fetch interface {
		GetInfoByIIN(iin string)
	}
)
