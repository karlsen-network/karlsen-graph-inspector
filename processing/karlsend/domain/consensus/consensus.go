package consensus

import (
	karlsendConsensus "github.com/karlsen-network/karlsend/v2/domain/consensus"
	consensusDatabase "github.com/karlsen-network/karlsend/v2/domain/consensus/database"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/datastructures/ghostdagdatastore"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/model"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/model/externalapi"
	"github.com/karlsen-network/karlsend/v2/domain/prefixmanager/prefix"
	"github.com/karlsen-network/karlsend/v2/infrastructure/db/database"
)

func New(consensusConfig *karlsendConsensus.Config, databaseContext database.Database, dbPrefix *prefix.Prefix, consensusEventsChan chan externalapi.ConsensusEvent) (*Consensus, bool, error) {
	karlsendConsensusFactory := karlsendConsensus.NewFactory()
	karlsendConsensusInstance, shouldMigrate, err := karlsendConsensusFactory.NewConsensus(consensusConfig, databaseContext, dbPrefix, consensusEventsChan)
	if err != nil {
		return nil, false, err
	}

	dbManager := consensusDatabase.New(databaseContext)
	pruningWindowSizeForCaches := int(consensusConfig.Params.PruningDepth())
	prefixBucket := consensusDatabase.MakeBucket(dbPrefix.Serialize())
	ghostdagDataStore := ghostdagdatastore.New(prefixBucket.Bucket([]byte{byte(0)}), pruningWindowSizeForCaches, true)

	return &Consensus{
		dbManager:         dbManager,
		karlsendConsensus: karlsendConsensusInstance,
		ghostdagDataStore: ghostdagDataStore,
	}, shouldMigrate, nil
}

type Consensus struct {
	dbManager         model.DBManager
	karlsendConsensus externalapi.Consensus
	ghostdagDataStore model.GHOSTDAGDataStore

	onBlockAddedListener      OnBlockAddedListener
	onVirtualResolvedListener OnVirtualResolvedListener
}

func (c *Consensus) ValidateAndInsertBlock(block *externalapi.DomainBlock, shouldValidateAgainstUTXO bool) error {
	err := c.karlsendConsensus.ValidateAndInsertBlock(block, shouldValidateAgainstUTXO)
	if err != nil {
		return err
	}
	if c.onBlockAddedListener != nil {
		c.onBlockAddedListener(block)
	}
	return nil
}

func (c *Consensus) ResolveVirtual(progressReportCallback func(uint64, uint64)) error {
	err := c.karlsendConsensus.ResolveVirtual(progressReportCallback)
	if err != nil {
		return err
	}
	if c.onVirtualResolvedListener != nil {
		c.onVirtualResolvedListener()
	}
	return nil
}

type OnBlockAddedListener func(*externalapi.DomainBlock)
type OnVirtualResolvedListener func()

func (c *Consensus) SetOnBlockAddedListener(listener OnBlockAddedListener) {
	c.onBlockAddedListener = listener
}

func (c *Consensus) SetOnVirtualResolvedListener(listener OnVirtualResolvedListener) {
	c.onVirtualResolvedListener = listener
}

func (c *Consensus) BlockGHOSTDAGData(blockHash *externalapi.DomainHash) (*externalapi.BlockGHOSTDAGData, error) {
	return c.ghostdagDataStore.Get(c.dbManager, model.NewStagingArea(), blockHash, false)
}
