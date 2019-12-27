ifdef RACE
	RACE_OPTION := -race
endif

test:
	@echo make test:
	GO111MODULE=on go test -v -cover $(RACE_OPTION) ./...