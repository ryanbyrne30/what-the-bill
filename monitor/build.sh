#!/bin/bash

O_DIR=.bin

O_QUEUE_SERVER=serve_queue
O_QUEUE_DISPATCH=dispatch_queue
O_US_BILLS_SERVER=serve_us_bills
O_US_BILLS_FETCH=fetch_us_bills

go build -o $O_DIR/$O_QUEUE_SERVER queue/cmd/server/*.go
go build -o $O_DIR/$O_QUEUE_DISPATCH queue/cmd/dispatch/*.go
go build -o $O_DIR/$O_US_BILLS_SERVER consumer_us_bills/*.go
go build -o $O_DIR/$O_US_BILLS_FETCH monitor_us_bills/*.go