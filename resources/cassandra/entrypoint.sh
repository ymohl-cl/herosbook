#!/bin/bash

if [ ! -f cass_is_init ]; then
	# start cassandra on background and get this pid number
	./docker-entrypoint.sh > logs 2>&1 &
	PID=$!

	# wait cassandra being up
	while ! cqlsh --execute="SELECT now() FROM system.local;"; do
		sleep 3
	done

	# insert dump cassandra
	for dump in $(ls /opt/data/*.cql); do
		cqlsh < $dump
	done

	kill $PID

	while netstat -ano | grep "127.0.0.1:7199"; do
		sleep 3
	done

	touch /cass_is_init
fi

./docker-entrypoint.sh