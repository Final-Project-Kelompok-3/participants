#!/usr/bin/env bash

go run main.go -migrate=migrate

go run main.go -migrate=status

go run main.go -migrate=seed

go run main.go -migrate=status

go run main.go