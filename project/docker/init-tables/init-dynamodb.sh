#!/bin/bash

echo "Creating tables..."

aws dynamodb create-table --cli-input-json file:///init-tables/user-table.json --endpoint-url http://dynamodb:8000
aws dynamodb create-table --cli-input-json file:///init-tables/product-table.json --endpoint-url http://dynamodb:8000
aws dynamodb create-table --cli-input-json file:///init-tables/order-table.json --endpoint-url http://dynamodb:8000
aws dynamodb create-table --cli-input-json file:///init-tables/order-item-table.json --endpoint-url http://dynamodb:8000

echo "Waiting for tables to be created..."
sleep 10

echo "Inserting initial data..."

aws dynamodb batch-write-item --request-items file:///init-tables/initial-data.json --endpoint-url http://dynamodb:8000

echo "Initialization completed."
