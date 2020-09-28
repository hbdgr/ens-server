package eth

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func (e *NameService) getAllEvents(ctx context.Context) error {
	filterQuery := createFilterQuery()
	rawLogs, err := e.client.FilterLogs(ctx, filterQuery)
	if err != nil {
		return err
	}

	fmt.Println("Gotlogs, count: ", len(rawLogs))

	// for _, rawLog := range rawLogs {

	// fmt.Printf("%x\n", rawLog.Data)
	// fmt.Printf("%x\n", rawLog.Topics)
	// log := NewOwnerLog{}
	// err := json.Unmarshal(rawLog.Topics, &log)
	// if err != nil {
	// 	continue
	// }
	//
	// fmt.Println("unmarshaled log:", log)
	// }

	fmt.Println("Events finished")
	return nil
}

type NewOwnerLog struct {
	node  hexutil.Bytes
	label hexutil.Bytes
}

func createFilterQuery() ethereum.FilterQuery {
	hbdgr1234Node := "0x7fbac75b026ecf8e067ee87b458a5a6ab37fcb84c5ff3b889de27360dfef59e8"
	ensGorlinContractAddres := "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e"

	var addresses []common.Address
	addresses = append(addresses, common.HexToAddress(ensGorlinContractAddres))

	// NewOwner(bytes32 indexed node, bytes32 indexed label, address owner)
	eventName := "NewOwner(bytes32,bytes32,address)"

	query := []interface{}{
		eventName,
		hbdgr1234Node,
	}

	topics, err := abi.MakeTopics([][]interface{}{query}...)
	if err != nil {
		panic(err)
	}

	return ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    topics,
	}
}
