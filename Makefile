migration_up:
	docker run --rm -v $(PWD)/services/url-management/infra/migrations:/migrations --user $(shell id -u):$(shell id -g) --network host migrate/migrate -path=/migrations/ -database "mysql://root:password@tcp(localhost:3306)/url_management" up

migration_down:
	docker run --rm -v $(PWD)/services/url-management/infra/migrations:/migrations --user $(shell id -u):$(shell id -g) --network host migrate/migrate -path=/migrations/ -database "mysql://root:password@tcp(localhost:3306)/url_management" down 1

migration_create:
	docker run --rm -v $(PWD)/services/url-management/infra/migrations:/migrations --user $(shell id -u):$(shell id -g) --network host migrate/migrate create -ext sql -dir /migrations  init

run_url_management:
	go run services/url-management/main.go

run_url_redirect:
	go run services/url-redirect/main.go

terraform_fmt:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(shell id -u):$(shell id -g) --network host hashicorp/terraform:light fmt

terraform_init:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(shell id -u):$(shell id -g) --network host hashicorp/terraform:light init -upgrade

terraform_plan:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(shell id -u):$(shell id -g) --network host hashicorp/terraform:light plan -out plan

terraform_apply:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(shell id -u):$(shell id -g) --network host hashicorp/terraform:light apply -auto-approve "plan"

terraform_destroy:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(shell id -u):$(shell id -g) --network host hashicorp/terraform:light destroy -auto-approve

