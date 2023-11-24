package main

import (
	"fmt"

	databasePackage "github.com/karlsen-network/karlsen-graph-inspector/processing/database"
	configPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/infrastructure/config"
	"github.com/karlsen-network/karlsen-graph-inspector/processing/infrastructure/logging"
	karlsendPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/karlsend"
	processingPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/processing"
	versionPackage "github.com/karlsen-network/karlsen-graph-inspector/processing/version"
	"github.com/karlsen-network/karlsend/version"
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
	logging.Logger().Infof("Embedded kalrsend version %s", version.Version())
	logging.Logger().Infof("Network %s", config.ActiveNetParams.Name)

	database, err := databasePackage.Connect(config.DatabaseConnectionString)
	if err != nil {
		logging.LogErrorAndExit("Could not connect to database %s: %s", config.DatabaseConnectionString, err)
	}
	defer database.Close()

	karlsend, err := karlsendPackage.New(config)
	if err != nil {
		logging.LogErrorAndExit("Could not create karlsend: %s", err)
	}
	processing, err := processingPackage.NewProcessing(config, database, karlsend)
	if err != nil {
		logging.LogErrorAndExit("Could not initialize processing: %s", err)
	}

	karlsend.SetOnVirtualResolvedListener(func() {
		err := processing.ResyncVirtualSelectedParentChain()
		if err != nil {
			logging.LogErrorAndExit("Could not resync the virtual selected parent chain: %s", err)
		}
	})
	karlsend.SetOnConsensusResetListener(func() {
		err := processing.ResyncDatabase()
		if err != nil {
			logging.LogErrorAndExit("Could not resync database: %s", err)
		}
	})
	err = karlsend.Start()
	if err != nil {
		logging.LogErrorAndExit("Could not start karlsend: %s", err)
	}

	<-make(chan struct{})
}
