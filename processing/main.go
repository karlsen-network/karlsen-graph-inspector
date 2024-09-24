package main

import (
	"fmt"
	"github.com/karlsen-network/karlsend/v2/infrastructure/network/rpcclient"

	databasePackage "github.com/karlsen-network/karlsen-graph-inspector/processing/database"
	configPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/infrastructure/config"
	"github.com/karlsen-network/karlsen-graph-inspector/processing/infrastructure/logging"
	processingPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/processing"
	versionPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/version"
	"github.com/karlsen-network/karlsend/v2/version"
)

func main() {
	fmt.Println("===================================================")
	fmt.Println("Karlsen Graph Inspector (KGI)   -   Processing Tier")
	fmt.Println("===================================================")

	config, err := configPackage.LoadConfig()
	if err != nil {
		logging.LogErrorAndExit("Could not parse command line arguments.\n%s", err)
	}

	logging.Logger().Infof("Application version %s", versionPackage.Version())
	logging.Logger().Infof("Embedded karlsend version %s", version.Version())
	logging.Logger().Infof("Network %s", config.ActiveNetParams.Name)

	database, err := databasePackage.Connect(config.DatabaseConnectionString)
	if err != nil {
		logging.LogErrorAndExit("Could not connect to database %s: %s", config.DatabaseConnectionString, err)
	}
	defer database.Close()

	rpcAddress, err := config.NetParams().NormalizeRPCServerAddress(config.RPCServer)
	if err != nil {
		panic(err)
	}
	rpcClient, err := rpcclient.NewRPCClient(rpcAddress)
	if err != nil {
		panic(err)
	}

	_, err = processingPackage.NewProcessing(config, database, rpcClient)
	if err != nil {
		logging.LogErrorAndExit("Could not initialize processing: %s", err)
	}

	<-make(chan struct{})
}
