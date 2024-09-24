package database

import (
	"github.com/karlsen-network/karlsen-graph-inspector/v2/processing/infrastructure/config"
	"github.com/karlsen-network/karlsen-graph-inspector/v2/processing/infrastructure/logging"
	"github.com/karlsen-network/karlsend/v2/infrastructure/db/database"
	"github.com/karlsen-network/karlsend/v2/infrastructure/db/database/ldb"
	"path/filepath"
)

const (
	databaseDirectoryName = "database"
	levelDBCacheSizeMiB   = 256
)

var (
	log = logging.Logger()
)

func Open(config *config.Config) (database.Database, error) {
	databaseDirectory := filepath.Join(config.AppDir, databaseDirectoryName)
	log.Infof("Loading database from '%s'", databaseDirectory)
	return ldb.NewLevelDB(databaseDirectory, levelDBCacheSizeMiB)
}
