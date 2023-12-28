# Karlsen Graph Inspector

Karlsen Graph Inspector (KGI) is comprised of four components:

* A PostgreSQL database
* A `processing` karlsen node (this is simply a karlsend wrapped in some
  extra logic)
* An `api` REST server
* A `web` server

How the components interact:

* The `processing` node connects to the Karlsen network the same way a
  regular karlsend node does and starts syncing just as a karlsend node
  would.
* While it's syncing, it writes metadata about every block to the
  PostgreSQL database.
* From the other end, the `web` server listens to http requests on some
  port.
* When a user navigates their browser to that port, the `web` server
  serves the KGI clientside logic, which includes the UI.
* The clientside logic calls the `api` REST server every so often.
* The `api` REST server queries the PostgreSQL database and returns it to
  the clientside.
* The clientside uses the response it received from the `api` REST
  server to update the UI.

## Deployment

Deploy a PostgreSQL database instance in any way you desire. Note the
address, port, username, password, and database name, since these will
be required later.

### Build KGI Sync

Make sure the go build environment is set up by running `go version`.
Build and install the processing binary:

```
cd processing
go build -o kgi-processing .
cd ..
mkdir -p kgi/sync/database
cp -r processing/database/migrations kgi/sync/database
cp processing/kgi-processing kgi/sync
```

### Build KGI API Server

Make sure the nodejs build environment is set up by running
`npm version`. Build and install API Server:

```
cd api
npm install
cd ..
mkdir -p kgi
cp -r api kgi
```

### Build KGI Web Frontend

Make sure the nodejs build environment is set up by running
`npm version`. The build required to configure the following
environment variables:

* `REACT_APP_API_ADDRESS` which is the the public address of
  your previously build KGI API Server will run.
* `REACT_APP_EXPLORER_ADDRESS` which is the public address of
  you explorer to view details of a specific block.

Build and install web frontend:

```
cd web
npm install
export REACT_APP_API_ADDRESS=api.karlsencoin.com:4455 
export REACT_APP_EXPLORER_ADDRESS=explorer.karlsencoin.com
npm run build
cd ..
mkdir -p kgi
cp -r web kgi
```

## Running

After setting up PostgreSQL and building all components, you can start
running them:

### Run KGI Sync

Running KGI Sync requires to configure the following parameters in the
`--connection-string` to be modified:

* `<psql_user>` which is the username for authentication.
* `<psql_pass>` which is the password for authentication.
* `<psql_host>` which is the PostgreSQL server host to connect to.
* `<psql_port>` which is the PostgreSQL server port to connect to.
* `<psql_db>` which is the database to use.

Navigate to wherever you copied `kgi-processing` binary:

```
cd kgi/sync
./kgi-processing --connection-string=postgres://<psql_user>:<psql_pass>@<psql_host>:<psql_port>/<psql_db>?sslmode=disable
```

### Run KGI API Server

Navigate to wherever you copied `api` folder to:

```
cd kgi/api
export API_PORT=4455
npm run start
```

### Run KGI Web Frontend

Navigate to wherever you copied `web` to:

```
cd kgi/web
npm install -g serve
serve --listen=8080
```
