#!/bin/bash
# scripts/run_all_tests.sh
echo "Running Go tests..."
go test ./internal/... -v

echo "Running Frontend tests..."
cd frontend && npm run test:unit