.PHONY: proto

UID := $(shell id -u)
GID := $(shell id -g)

migration_up:
	docker run --rm -v $(PWD)/services/link_management/external/database/migrations:/migrations --user $(UID):$(GID) --network host migrate/migrate -path=/migrations/ -database "mysql://root:password@tcp(localhost:3306)/link_management" up

migration_down:
	docker run --rm -v $(PWD)/services/link_management/external/database/migrations:/migrations --user $(UID):$(GID) --network host migrate/migrate -path=/migrations/ -database "mysql://root:password@tcp(localhost:3306)/url_management" down 1

migration_create:
	docker run --rm -v $(PWD)/services/link_management/external/database/migrations:/migrations --user $(UID):$(GID) --network host migrate/migrate create -ext sql -dir /migrations  init

run_url_management:
	go run services/url-management/main.go

run_url_redirect:
	go run services/link_redirect/main.go

terraform_fmt:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(UID):$(GID) --network host hashicorp/terraform:light fmt

terraform_init:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(UID):$(GID) --network host hashicorp/terraform:light init -upgrade

terraform_plan:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(UID):$(GID) --network host hashicorp/terraform:light plan -out plan

terraform_apply:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(UID):$(GID) --network host hashicorp/terraform:light apply -auto-approve "plan"

terraform_destroy:
	docker run --rm -v $(PWD)/infra/terraform:/infra -w /infra --user $(UID):$(GID) --network host hashicorp/terraform:light destroy -auto-approve

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative services/$(service)/external/grpc/proto/*.proto