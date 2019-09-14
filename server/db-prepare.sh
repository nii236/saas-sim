#!/bin/bash

source db-prepare-userpass.sh

../bin/migrate -database "postgres://$LOCAL_DEV_DB_USER:$LOCAL_DEV_DB_PASS@$LOCAL_DEV_DB_HOST:$LOCAL_DEV_DB_PORT/$LOCAL_DEV_DB_DATABASE?sslmode=disable" -path ./migrations drop
../bin/migrate -database "postgres://$LOCAL_DEV_DB_USER:$LOCAL_DEV_DB_PASS@$LOCAL_DEV_DB_HOST:$LOCAL_DEV_DB_PORT/$LOCAL_DEV_DB_DATABASE?sslmode=disable" -path ./migrations up
