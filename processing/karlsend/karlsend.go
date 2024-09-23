package karlsend

import (
	configPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/infrastructure/config"
	"github.com/karlsen-network/karlsen-graph-inspector/processing/infrastructure/database"
	//"github.com/karlsen-network/karlsen-graph-inspector/processing/infrastructure/logging"
	domainPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/karlsend/domain"
	consensusPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/karlsend/domain/consensus"
	"github.com/karlsen-network/karlsend/v2/app/protocol"
	"github.com/karlsen-network/karlsend/v2/domain/consensus/model/externalapi"
	karlsendConfigPackage "github.com/karlsen-network/karlsend/v2/infrastructure/config"
	"github.com/karlsen-network/karlsend/v2/infrastructure/network/addressmanager"
	"github.com/karlsen-network/karlsend/v2/infrastructure/network/connmanager"
	"github.com/karlsen-network/karlsend/v2/infrastructure/network/netadapter"
	"github.com/karlsen-network/karlsend/v2/infrastructure/network/netadapter/router"
	"net"
)

type Karlsend struct {
	config            *configPackage.Config
	domain            *domainPackage.Domain
	netAdapter        *netadapter.NetAdapter
	addressManager    *addressmanager.AddressManager
	connectionManager *connmanager.ConnectionManager
	protocolManager   *protocol.Manager
}

func New(config *configPackage.Config) (*Karlsend, error) {
	karlsendConfig := karlsendConfigPackage.DefaultConfig()
	karlsendConfig.ConnectPeers = config.ConnectPeers
	karlsendConfig.DNSSeed = config.DNSSeed
	karlsendConfig.GRPCSeed = config.GRPCSeed
	karlsendConfig.NetworkFlags = config.NetworkFlags
	karlsendConfig.Lookup = net.LookupIP

	//logging.UpdateLogLevels()

	databaseContext, err := database.Open(config)
	if err != nil {
		return nil, err
	}
	domain, err := domainPackage.New(config.NetworkFlags.ActiveNetParams, databaseContext)
	if err != nil {
		return nil, err
	}
	netAdapter, err := netadapter.NewNetAdapter(karlsendConfig)
	if err != nil {
		return nil, err
	}
	netAdapter.SetRPCRouterInitializer(func(router *router.Router, connection *netadapter.NetConnection) {})
	addressManager, err := addressmanager.New(addressmanager.NewConfig(karlsendConfig), databaseContext)
	if err != nil {
		return nil, err
	}
	connectionManager, err := connmanager.New(karlsendConfig, netAdapter, addressManager)
	if err != nil {
		return nil, err
	}
	protocolManager, err := protocol.NewManager(karlsendConfig, domain, netAdapter, addressManager, connectionManager)
	if err != nil {
		return nil, err
	}
	return &Karlsend{
		config:            config,
		domain:            domain,
		netAdapter:        netAdapter,
		addressManager:    addressManager,
		connectionManager: connectionManager,
		protocolManager:   protocolManager,
	}, nil
}

func (k *Karlsend) SetOnBlockAddedListener(listener consensusPackage.OnBlockAddedListener) {
	k.domain.SetOnBlockAddedListener(listener)
}

func (k *Karlsend) SetOnVirtualResolvedListener(listener consensusPackage.OnVirtualResolvedListener) {
	k.domain.SetOnVirtualResolvedListener(listener)
}

func (k *Karlsend) SetOnConsensusResetListener(listener domainPackage.OnConsensusResetListener) {
	k.domain.SetOnConsensusResetListener(listener)
}

func (k *Karlsend) BlockGHOSTDAGData(blockHash *externalapi.DomainHash) (*externalapi.BlockGHOSTDAGData, error) {
	return k.domain.BlockGHOSTDAGData(blockHash)
}

func (k *Karlsend) Start() error {
	err := k.netAdapter.Start()
	if err != nil {
		return err
	}
	k.connectionManager.Start()
	return nil
}

func (k *Karlsend) Domain() *domainPackage.Domain {
	return k.domain
}
