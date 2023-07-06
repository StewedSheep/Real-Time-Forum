#!/bin/bash

# Change directory to frontend
cd frontend

# Install dependencies
npm install

# Build the frontend
npm run build

# Move back to the parent directory
cd ..

# Run the Go program
go run .
