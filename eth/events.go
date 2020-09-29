package eth

import (
	"context"
	"ens_feed/model"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ens "github.com/wealdtech/go-ens/v3"
	"golang.org/x/crypto/sha3"
)

func childNameHash(parentNode, labelHash common.Hash) (hash [32]byte, err error) {
	sha := sha3.NewLegacyKeccak256()
	if _, err = sha.Write(parentNode.Bytes()); err != nil {
		return
	}
	if _, err = sha.Write(labelHash.Bytes()); err != nil {
		return
	}
	sha.Sum(hash[:0])
	return
}

func (e *NameService) subnameInfoFromLog(ctx context.Context, newOwnerLog *types.Log) (*model.SubnameInfo, error) {
	if len(newOwnerLog.Topics) < 3 {
		return nil, fmt.Errorf("ERROR: bad event log type for subdomainFromLog %#v", newOwnerLog)
	}

	parentNode := newOwnerLog.Topics[1]
	labelHash := newOwnerLog.Topics[2]

	childNode, err := childNameHash(parentNode, labelHash)
	if err != nil {
		return nil, err
	}

	owner, err := e.nodeOwner(childNode)
	if err != nil {
		return nil, err
	}

	resolver, err := e.nodeResolver(childNode)
	if err != nil {
		return nil, err
	}

	name, err := e.resolveWith(childNode, resolver)

	subname := &model.SubnameInfo{
		LabelHash:   labelHash.String(),
		SubnameNode: fmt.Sprintf("0x%x", childNode),
		Owner:       owner.String(),
		Resolver:    resolver.String(),
		Subname:     name,
	}

	if err != nil {
		errMsg := err.Error()
		subname.ErrorMsg = &errMsg
	}

	return subname, nil
}

func (e *NameService) newOwnerEventsFor(ctx context.Context, parentName string) (map[common.Hash]types.Log, error) {
	parentNameHash, err := ens.NameHash(parentName)
	if err != nil {
		return nil, err
	}

	filterQuery := e.createNewOwnerQuery(parentNameHash)
	rawLogs, err := e.client.FilterLogs(ctx, filterQuery)
	if err != nil {
		return nil, err
	}

	logsMap, err := latestNewOwnerLogs(rawLogs)
	if err != nil {
		return nil, err
	}

	log.Printf("ENS EVENTS: Got NewOwner event logs for %s, count: %d", parentName, len(rawLogs))
	return logsMap, nil
}

// only the newest logs are up to date for the same node and labelHash
// use map to easily override deprecated logs
func latestNewOwnerLogs(logs []types.Log) (map[common.Hash]types.Log, error) {
	if len(logs) > 0 && len(logs[0].Topics) < 3 {
		return nil, fmt.Errorf("ERROR: bad event log type for latestNewOwnerLogs %#v", logs[0])
	}

	logsMap := make(map[common.Hash]types.Log)

	for _, log := range logs {
		labelHash := log.Topics[2]
		logsMap[labelHash] = log
	}

	return logsMap, nil
}

func (e *NameService) createNewOwnerQuery(nodeFilter interface{}) ethereum.FilterQuery {
	// NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
	eventName := "NewOwner(bytes32,bytes32,address)"

	var addresses []common.Address
	addresses = append(addresses, e.contractAddr)

	query := [][]interface{}{
		{
			eventName,
		},
		{
			nodeFilter,
		},
	}

	topics, err := abi.MakeTopics(query...)
	if err != nil {
		panic(err)
	}

	return ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    topics,
	}
}

func (e *NameService) createNewResolverQuery(blockHash *common.Hash) ethereum.FilterQuery {
	// NewResolver (bytes32 indexed node, address resolver)
	eventName := "NewResolver(bytes32,address)"

	var addresses []common.Address
	addresses = append(addresses, e.contractAddr)

	query := [][]interface{}{
		{
			eventName,
		},
	}

	topics, err := abi.MakeTopics(query...)
	if err != nil {
		panic(err)
	}

	return ethereum.FilterQuery{
		BlockHash: blockHash,
		Addresses: addresses,
		Topics:    topics,
	}
}
