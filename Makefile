test:
	go test ./...

build:
	go build \
	  -race \
	  -ldflags "-X gitlab.zerbytes.net/${CI_PROJECT_PATH}/vendor/github.com/prometheus/common/version.Version=$(cat VERSION) \
	  -X gitlab.zerbytes.net/${CI_PROJECT_PATH}/vendor/github.com/prometheus/common/version.Revision=${CI_COMMIT_SHA} \
	  -X gitlab.zerbytes.net/${CI_PROJECT_PATH}/vendor/github.com/prometheus/common/version.Branch=${CI_COMMIT_REF_NAME} \
	  -X gitlab.zerbytes.net/${CI_PROJECT_PATH}/vendor/github.com/prometheus/common/version.BuildUser=$(whoami)@$(hostname) \
	  -X gitlab.zerbytes.net/${CI_PROJECT_PATH}/vendor/github.com/prometheus/common/version.BuildDate=$(date +%Y%m%d-%H:%M:%S) \
	  -extldflags '-static'" \
	  -o app

.PHONY: test build
