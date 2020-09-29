package eth

import (
	"context"
	"ens_feed/model"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	ens "github.com/wealdtech/go-ens/v3"
	"github.com/wealdtech/go-ens/v3/contracts/resolver"
)

type NameService struct {
	client       *ethclient.Client
	contractAddr common.Address

	registry *ens.Registry
}

func NewNameService(ethURL, ensContractAddr string) (*NameService, error) {
	log.Println("Eth connecting to:", ethURL)
	client, err := ethclient.Dial(ethURL)
	if err != nil {
		return nil, err
	}
	log.Println("Eth connection ready")

	registry, err := ens.NewRegistry(client)
	if err != nil {
		return nil, err
	}

	contractAddr := common.HexToAddress(ensContractAddr)
	return &NameService{
		client:       client,
		contractAddr: contractAddr,
		registry:     registry,
	}, nil
}

// resolve node to the domain name, uses resolver contract from given address
// uses directly `addr(bytes32 node)` method from the resolver contract
// it allows to resolve name without knowing to what address it is pointing to
func (e *NameService) resolveWith(node [32]byte, resolverAddr common.Address) (string, error) {
	resolverContract, err := resolver.NewContract(resolverAddr, e.client)
	if err != nil {
		return "", err
	}

	addr, err := resolverContract.Addr(nil, node)
	if err != nil {
		return "", err
	}

	return ens.ReverseResolve(e.client, addr)
}

// uses registry contract directly, to get owner of the domain name for given node
// without knowing the domain name itself
func (e *NameService) nodeOwner(node [32]byte) (common.Address, error) {
	return e.registry.Contract.Owner(nil, node)
}

// uses registry contract directly, to get resolver of the domain name for given node
// without knowing the domain name itself
func (e *NameService) nodeResolver(node [32]byte) (common.Address, error) {
	return e.registry.Contract.Resolver(nil, node)
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

// Subnames returns available subdomains for given parent name
func (e *NameService) Subnames(ctx context.Context, parentName string) ([]string, error) {
	subnamesInfo, err := e.SubnamesInfo(ctx, parentName)
	if err != nil {
		return nil, err
	}

	subnames := make([]string, 0, len(subnamesInfo.Subnames))
	for _, subInfo := range subnamesInfo.Subnames {
		if subInfo.Subname != "" {
			subnames = append(subnames, subInfo.Subname)
		}
	}

	return subnames, nil
}

// SubnamesInfo returns detailed info about subdomains for given parent name
func (e *NameService) SubnamesInfo(ctx context.Context, parentName string) (*model.SubnamesInfo, error) {
	logs, err := e.newOwnerEventsFor(ctx, parentName)
	if err != nil {
		return nil, err
	}

	sns := &model.SubnamesInfo{Parent: parentName}
	sns.Subnames = make([]*model.SubnameInfo, 0, len(logs))

	for _, l := range logs {
		s, err := e.subnameInfoFromLog(ctx, &l)
		if err != nil {
			return nil, err
		}

		// deleted name - omitting
		if s.Owner == "0x0000000000000000000000000000000000000000" {
			continue
		}

		sns.Subnames = append(sns.Subnames, s)

		log.Printf("ENS: SubnamesInfo: %#v\n", s)
	}
	return sns, nil
}
