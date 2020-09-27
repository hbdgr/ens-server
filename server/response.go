package server

type resolveMsg struct {
	Name    string `json:"name"`
	EthAddr string `json:"eth_addr"`
}
