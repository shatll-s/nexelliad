package model

import "github.com/Nexellia-Network/nexelliad/domain/consensus/model/externalapi"

// BlockBuilder is responsible for creating blocks from the current state
type BlockBuilder interface {
	BuildBlock(coinbaseData *externalapi.DomainCoinbaseData,
		transactions []*externalapi.DomainTransaction) (block *externalapi.DomainBlock, coinbaseHasRedReward bool, err error)
}
