package eth

import (
	"context"
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

	ns := &NameService{client: client}

	err = ns.getAllEvents(context.Background())

	return ns, err
}

// Resolve ENS name into Ethereum address
func (e *NameService) Resolve(name string) (string, error) {
	address, err := ens.Resolve(e.client, name)
	if err != nil {
		return "", nil
	}

	log.Printf("ENS: Resolved test domain [%s], address: %s",
		name, address.String())

	return address.String(), nil
}
