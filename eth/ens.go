package eth

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

type NameService struct {
	client *ethclient.Client
}

func NewNameService(ethURL string) (*NameService, error) {
	client, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, err
	}
	log.Println("Eth connection ready")

	return &NameService{client: client}, nil
}

func (e *NameService) ResolveTest() (string, error) {
	testDomain := "hbdgr1234.eth"
	address, err := ens.Resolve(e.client, testDomain)
	if err != nil {
		return "", nil
	}

	log.Printf("ENS: Resolved test domain [%s], address: %s",
		testDomain, address.String())

	return address.String(), nil
}
