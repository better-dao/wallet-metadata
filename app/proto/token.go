package proto

// token contract meta:
type TokenContractMeta struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Logo     string `json:"logo"`
	Decimals int64  `json:"decimals"`
	Erc20    bool   `json:"erc20"`
}

func (m *TokenContractMeta) From() {

}

func (m *TokenContractMeta) To() []byte {
	return nil
}
