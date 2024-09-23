package consensus

import "github.com/karlsen-network/karlsend/v2/domain/consensus/model/externalapi"

func (c *Consensus) BuildBlock(coinbaseData *externalapi.DomainCoinbaseData, transactions []*externalapi.DomainTransaction) (*externalapi.DomainBlock, error) {
	return c.karlsendConsensus.BuildBlock(coinbaseData, transactions)
}

func (c *Consensus) BuildBlockTemplate(coinbaseData *externalapi.DomainCoinbaseData, transactions []*externalapi.DomainTransaction) (*externalapi.DomainBlockTemplate, error) {
	return c.karlsendConsensus.BuildBlockTemplate(coinbaseData, transactions)
}

func (c *Consensus) ValidateTransactionAndPopulateWithConsensusData(transaction *externalapi.DomainTransaction) error {
	return c.karlsendConsensus.ValidateTransactionAndPopulateWithConsensusData(transaction)
}

func (c *Consensus) GetBlock(blockHash *externalapi.DomainHash) (*externalapi.DomainBlock, bool, error) {
	return c.karlsendConsensus.GetBlock(blockHash)
}

func (c *Consensus) GetBlockHeader(blockHash *externalapi.DomainHash) (externalapi.BlockHeader, error) {
	return c.karlsendConsensus.GetBlockHeader(blockHash)
}

func (c *Consensus) GetBlockInfo(blockHash *externalapi.DomainHash) (*externalapi.BlockInfo, error) {
	return c.karlsendConsensus.GetBlockInfo(blockHash)
}

func (c *Consensus) GetHashesBetween(lowHash, highHash *externalapi.DomainHash, maxBlueScoreDifference uint64) (
	[]*externalapi.DomainHash, *externalapi.DomainHash, error) {

	return c.karlsendConsensus.GetHashesBetween(lowHash, highHash, maxBlueScoreDifference)
}

func (c *Consensus) GetAnticone(blockHash, contextHash *externalapi.DomainHash, maxBlocks uint64) (hashes []*externalapi.DomainHash, err error) {
	return c.karlsendConsensus.GetAnticone(blockHash, contextHash, maxBlocks)
}

func (c *Consensus) GetPruningPointUTXOs(expectedPruningPointHash *externalapi.DomainHash, fromOutpoint *externalapi.DomainOutpoint, limit int) ([]*externalapi.OutpointAndUTXOEntryPair, error) {
	return c.karlsendConsensus.GetPruningPointUTXOs(expectedPruningPointHash, fromOutpoint, limit)
}

func (c *Consensus) GetVirtualUTXOs(expectedVirtualParents []*externalapi.DomainHash, fromOutpoint *externalapi.DomainOutpoint, limit int) ([]*externalapi.OutpointAndUTXOEntryPair, error) {
	return c.karlsendConsensus.GetVirtualUTXOs(expectedVirtualParents, fromOutpoint, limit)
}

func (c *Consensus) PruningPoint() (*externalapi.DomainHash, error) {
	return c.karlsendConsensus.PruningPoint()
}

func (c *Consensus) ClearImportedPruningPointData() error {
	return c.karlsendConsensus.ClearImportedPruningPointData()
}

func (c *Consensus) AppendImportedPruningPointUTXOs(outpointAndUTXOEntryPairs []*externalapi.OutpointAndUTXOEntryPair) error {
	return c.karlsendConsensus.AppendImportedPruningPointUTXOs(outpointAndUTXOEntryPairs)
}

func (c *Consensus) GetVirtualSelectedParent() (*externalapi.DomainHash, error) {
	return c.karlsendConsensus.GetVirtualSelectedParent()
}

func (c *Consensus) CreateHeadersSelectedChainBlockLocator(lowHash, highHash *externalapi.DomainHash) (externalapi.BlockLocator, error) {
	return c.karlsendConsensus.CreateHeadersSelectedChainBlockLocator(lowHash, highHash)
}

func (c *Consensus) CreateFullHeadersSelectedChainBlockLocator() (externalapi.BlockLocator, error) {
	return c.karlsendConsensus.CreateFullHeadersSelectedChainBlockLocator()
}

func (c *Consensus) GetSyncInfo() (*externalapi.SyncInfo, error) {
	return c.karlsendConsensus.GetSyncInfo()
}

func (c *Consensus) Tips() ([]*externalapi.DomainHash, error) {
	return c.karlsendConsensus.Tips()
}

func (c *Consensus) GetVirtualInfo() (*externalapi.VirtualInfo, error) {
	return c.karlsendConsensus.GetVirtualInfo()
}

func (c *Consensus) IsValidPruningPoint(blockHash *externalapi.DomainHash) (bool, error) {
	return c.karlsendConsensus.IsValidPruningPoint(blockHash)
}

func (c *Consensus) GetVirtualSelectedParentChainFromBlock(blockHash *externalapi.DomainHash) (*externalapi.SelectedChainPath, error) {
	return c.karlsendConsensus.GetVirtualSelectedParentChainFromBlock(blockHash)
}

func (c *Consensus) IsInSelectedParentChainOf(blockHashA *externalapi.DomainHash, blockHashB *externalapi.DomainHash) (bool, error) {
	return c.karlsendConsensus.IsInSelectedParentChainOf(blockHashA, blockHashB)
}

func (c *Consensus) GetHeadersSelectedTip() (*externalapi.DomainHash, error) {
	return c.karlsendConsensus.GetHeadersSelectedTip()
}

func (c *Consensus) Anticone(blockHash *externalapi.DomainHash) ([]*externalapi.DomainHash, error) {
	return c.karlsendConsensus.Anticone(blockHash)
}

func (c *Consensus) GetBlockRelations(blockHash *externalapi.DomainHash) (
	parents []*externalapi.DomainHash, children []*externalapi.DomainHash, err error) {

	return c.karlsendConsensus.GetBlockRelations(blockHash)
}

func (s *Consensus) GetBlockAcceptanceData(blockHash *externalapi.DomainHash) (externalapi.AcceptanceData, error) {
	return s.karlsendConsensus.GetBlockAcceptanceData(blockHash)
}

func (s *Consensus) GetBlocksAcceptanceData(blockHashes []*externalapi.DomainHash) ([]externalapi.AcceptanceData, error) {
	return s.karlsendConsensus.GetBlocksAcceptanceData(blockHashes)
}

func (c *Consensus) GetBlockEvenIfHeaderOnly(blockHash *externalapi.DomainHash) (*externalapi.DomainBlock, error) {
	return c.karlsendConsensus.GetBlockEvenIfHeaderOnly(blockHash)
}

func (c *Consensus) EstimateNetworkHashesPerSecond(startHash *externalapi.DomainHash, windowSize int) (uint64, error) {
	return c.karlsendConsensus.EstimateNetworkHashesPerSecond(startHash, windowSize)
}

func (c *Consensus) GetVirtualDAAScore() (uint64, error) {
	return c.karlsendConsensus.GetVirtualDAAScore()
}

func (c *Consensus) Init(shouldNotAddGenesis bool) error {
	return c.karlsendConsensus.Init(shouldNotAddGenesis)
}

func (c *Consensus) PruningPointAndItsAnticone() ([]*externalapi.DomainHash, error) {
	return c.karlsendConsensus.PruningPointAndItsAnticone()
}

func (c *Consensus) ValidateAndInsertImportedPruningPoint(newPruningPoint *externalapi.DomainHash) error {
	return c.karlsendConsensus.ValidateAndInsertImportedPruningPoint(newPruningPoint)
}

func (c *Consensus) CreateBlockLocatorFromPruningPoint(highHash *externalapi.DomainHash, limit uint32) (externalapi.BlockLocator, error) {
	return c.karlsendConsensus.CreateBlockLocatorFromPruningPoint(highHash, limit)
}

func (c *Consensus) PopulateMass(transaction *externalapi.DomainTransaction) {
	c.karlsendConsensus.PopulateMass(transaction)
}

func (c *Consensus) ValidateAndInsertBlockWithTrustedData(block *externalapi.BlockWithTrustedData, validateUTXO bool) error {
	return c.karlsendConsensus.ValidateAndInsertBlockWithTrustedData(block, validateUTXO)
}

func (c *Consensus) GetMissingBlockBodyHashes(highHash *externalapi.DomainHash) ([]*externalapi.DomainHash, error) {
	return c.karlsendConsensus.GetMissingBlockBodyHashes(highHash)
}

func (c *Consensus) ImportPruningPoints(pruningPoints []externalapi.BlockHeader) error {
	return c.karlsendConsensus.ImportPruningPoints(pruningPoints)
}

func (c *Consensus) BuildPruningPointProof() (*externalapi.PruningPointProof, error) {
	return c.karlsendConsensus.BuildPruningPointProof()
}

func (c *Consensus) ValidatePruningPointProof(pruningPointProof *externalapi.PruningPointProof) error {
	return c.karlsendConsensus.ValidatePruningPointProof(pruningPointProof)
}

func (c *Consensus) ApplyPruningPointProof(pruningPointProof *externalapi.PruningPointProof) error {
	return c.karlsendConsensus.ApplyPruningPointProof(pruningPointProof)
}

func (c *Consensus) PruningPointHeaders() ([]externalapi.BlockHeader, error) {
	return c.karlsendConsensus.PruningPointHeaders()
}

func (c *Consensus) ArePruningPointsViolatingFinality(pruningPoints []externalapi.BlockHeader) (bool, error) {
	return c.karlsendConsensus.ArePruningPointsViolatingFinality(pruningPoints)
}

func (c *Consensus) BlockDAAWindowHashes(blockHash *externalapi.DomainHash) ([]*externalapi.DomainHash, error) {
	return c.karlsendConsensus.BlockDAAWindowHashes(blockHash)
}

func (c *Consensus) TrustedDataDataDAAHeader(trustedBlockHash, daaBlockHash *externalapi.DomainHash, daaBlockWindowIndex uint64) (*externalapi.TrustedDataDataDAAHeader, error) {
	return c.karlsendConsensus.TrustedDataDataDAAHeader(trustedBlockHash, daaBlockHash, daaBlockWindowIndex)
}

func (c *Consensus) TrustedBlockAssociatedGHOSTDAGDataBlockHashes(blockHash *externalapi.DomainHash) ([]*externalapi.DomainHash, error) {
	return c.karlsendConsensus.TrustedBlockAssociatedGHOSTDAGDataBlockHashes(blockHash)
}

func (c *Consensus) TrustedGHOSTDAGData(blockHash *externalapi.DomainHash) (*externalapi.BlockGHOSTDAGData, error) {
	return c.karlsendConsensus.TrustedGHOSTDAGData(blockHash)
}

func (c *Consensus) IsChainBlock(blockHash *externalapi.DomainHash) (bool, error) {
	return c.karlsendConsensus.IsChainBlock(blockHash)
}

func (c *Consensus) VirtualMergeDepthRoot() (*externalapi.DomainHash, error) {
	return c.karlsendConsensus.VirtualMergeDepthRoot()
}

func (c *Consensus) IsNearlySynced() (bool, error) {
	return c.karlsendConsensus.IsNearlySynced()
}
