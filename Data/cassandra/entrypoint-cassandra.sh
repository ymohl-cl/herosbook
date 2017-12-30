#!/bin/bash

./docker-entrypoint.sh cassandra &

# dump cql
while ! cqlsh --execute="SELECT now() FROM system.local;"; do
	sleep 1
done

for dump in $(ls /opt/data/*.cql); do
	cqlsh < $dump
done

# This loop keep the container alive.
while :; do
	sleep 100
done
