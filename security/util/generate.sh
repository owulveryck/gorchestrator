echo "Generating Orchestrator certificate"
go run generate_cert.go -ca -host localhost
mv cert.pem ../certs/orchestrator/orchestrator.pem
mv key.pem ../certs/orchestrator/orchestrator_key.pem
echo "Generating Executor certificate"
go run generate_cert.go -ca -host localhost
mv cert.pem ../certs/executor/executor.pem
mv key.pem ../certs/executor/executor_key.pem
