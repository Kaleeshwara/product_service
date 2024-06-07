# Go parameters
GOCMD := go
GORUN := $(GOCMD) run
PKG_DIR := ./

# These envs will be coming from the docker config, Make sure not to push this
add-env:
    set PGHOST=localhost
    set PGPORT=5432
    set PGUSER=postgres
    set PGPASSWORD=password
    set PGDATABASE=product
    set PGSCHEMA=public

run: add-env
    $(GORUN) $(PKG_DIR)

gen-swagger:
    swag init -g $(PKG_DIR)main.go