echo "Generating Orchestrator certificate"
go run generate_cert.go -ca -host 127.0.0.1 -target orchestrator
mv orchestrator.pem ../certs/orchestrator/
mv orchestrator_key.pem ../certs/orchestrator/
echo "Generating Executor certificate"
go run generate_cert.go -ca -host 127.0.0.1 -target executor
mv executor.pem ../certs/executor/
mv executor_key.pem ../certs/executor/
