package server

type resolveMsg struct {
	EthAddr string `json:"eth_addr"`
}

type reverseResolveMsg struct {
	Name string `json:"name"`
}

type subdomainsMsg struct {
	subdomains []string `json:"names"`
}

type errorMsg struct {
	ErrorMsg string `json:"error_message"`
}
