
# Shell color codes...
#     ref -> https://stackoverflow.com/a/5947802/667301
CLR_GREEN=\033[0;32m
CLR_CYAN=\033[0;36m
CLR_YELLOW=\033[0;33m
CLR_RED=\033[0;31m
CLR_END=\033[0;0m

.DEFAULT_GOAL := all

all:
	@echo "$(CLR_GREEN)>> Building the 'routeviews_go' Go 'go.mod' file.$(CLR_END)"
	@echo "$(CLR_CYAN)    >> Removing the old ./routeviews_go binary.$(CLR_END)"
	-echo "    Removing the old ./routeviews_go executable"
	-rm -rf ./routeviews_go
	@echo "$(CLR_CYAN)    >> Removing the old ./go.mod file.$(CLR_END)"
	-rm -rf ./go.mod
	@echo "$(CLR_CYAN)    >> Building a new go.mod file.$(CLR_END)"
	-echo "module routeviews_go" > ./go.mod
	-go mod tidy -v
	@echo "$(CLR_CYAN)    >> vetting...$(CLR_END)"
	cd src && go vet .
	@echo "$(CLR_CYAN)    >> compiling...$(CLR_END)"
	-cd src && go build -ldflags "-s -w" -o ../routeviews_go .
	# Install all python dependencies
.PHONY: all

