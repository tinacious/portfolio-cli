build:
	go build

lint:
	goimports -w .
	gci write .

version:
	./scripts/create_git_version_tag.sh
