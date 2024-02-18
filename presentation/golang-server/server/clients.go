package server

type ZajoncStore struct {
	secretMessage string
}

func GetDefaultZajoncStore() *ZajoncStore {
	return &ZajoncStore{
		"zajonc secret message",
	}
}

func (z *ZajoncStore) GetZajonc() string {
	return z.secretMessage
}
