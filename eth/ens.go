package eth

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ens "github.com/wealdtech/go-ens/v3"
)

type NameService struct {
	client       *ethclient.Client
	contractAddr common.Address
}

func NewNameService(ethURL, ensContractAddr string) (*NameService, error) {
	log.Println("Eth connecting to:", ethURL)
	client, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, err
	}
	log.Println("Eth connection ready")

	contractAddr := common.HexToAddress(ensContractAddr)
	ns := &NameService{
		client:       client,
		contractAddr: contractAddr,
	}

	err = ns.getAllEvents(context.Background())

	ns.playground()

	return ns, err
}

// Resolve ENS name into Ethereum address
func (e *NameService) Resolve(name string) (string, error) {
	address, err := ens.Resolve(e.client, name)
	if err != nil {
		return "", err
	}

	log.Printf("ENS: Resolved name [%s], address: %s", name, address.String())

	return address.String(), nil
}

// ReverseResolve address to ENS name
func (e *NameService) ReverseResolve(addressStr string) (string, error) {
	if ok := common.IsHexAddress(addressStr); !ok {
		return "", fmt.Errorf("Address in bad hex format: [%s]", addressStr)
	}

	address := common.HexToAddress(addressStr)
	name, err := ens.ReverseResolve(e.client, address)
	if err != nil {
		return "", err
	}

	log.Printf("ENS: Reverse Resolved address [%s], name: %s", address.String(), name)

	return name, nil
}
