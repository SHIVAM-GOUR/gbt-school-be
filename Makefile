export DEV_DSN = postgresql://postgres:password@localhost:5432/gbt_school_local
export PROD_DSN = postgresql://postgres:fTCNqYhMCcceJvrZJknDNNvxmgRzUXLc@switchyard.proxy.rlwy.net:30216/railway

# Run Migration down
.PHONY: run_migration_down
run_migration_down:
		migrate -path ./migrations -database "${DEV_DSN}" down


# Run the project (Dev server)
.PHONY: run_dev_server
run_dev_server:
		migrate -path ./migrations -database "${DEV_DSN}" up
		air

# Run the project (production)
.PHONY: run_prod_server
run_prod_server:
		migrate -path ./migrations -database "${PROD_DSN}" up
		go build -o main
		./main
