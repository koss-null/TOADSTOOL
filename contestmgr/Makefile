export GO111MODULE=on

# Parameters
PKG = github.com/koss-null/toadstool/contestmgr
NAME = contestmgr

# ATTENSION: this variables are copy-pasted, so there can be some stupid templates left there
# Replace backslashes with forward slashes for use on Windows.
E :=
BSLASH := \$E
FSLASH := /

# Directories
WD := $(subst $(BSLASH),$(FSLASH),$(shell pwd))
MD := $(subst $(BSLASH),$(FSLASH),$(shell dirname "$(realpath $(lastword $(MAKEFILE_LIST)))"))
PKGDIR = $(MD)
CMDDIR = $(PKGDIR)/cmd
RPTDIR ?= $(WD)/reports
GP = $(subst $(BSLASH),$(FSLASH),$(GOPATH))

# Parameters
VERSION ?= $(shell git -C "$(MD)" describe --tags --dirty=-dev)
COMMIT_ID := $(shell git -C "$(MD)" rev-parse HEAD | head -c8)
BUILD_TAGS ?= release
CMDPKG = $(PKG)/cmd
# CMDS := $(shell find "$(CMDDIR)/" -mindepth 1 -maxdepth 1 -type d | sed 's/ /\\ /g' | xargs -n1 basename)
BENCHCPUS ?= 1,2,4

# Commands
GOCMD = go
ARCHES ?= 386 amd64
OSES ?= linux 
OUTTPL = ./build/$(NAME)
LDFLAGS = -X $(PKG).Version=$(VERSION) -X $(PKG).Build=$(COMMIT_ID)
GOBUILD = $(GOCMD) build -o "$(OUTTPL)" ./...
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install -a -tags "$(BUILD_TAGS)" -ldflags "$(LDFLAGS)"
GOTEST = $(GOCMD) test -v -tags "$(BUILD_TAGS)"
GOLINT = gometalinter --deadline=30s --tests --disable=aligncheck --disable=gocyclo --disable=gotype
GODEP = $(GOCMD) get -d -t
GOFMT = goreturns -w
GOBENCH = $(GOCMD) test -v -tags "$(BUILD_TAGS)" -cpu=$(BENCHCPUS) -run=NOTHING -bench=. -benchmem -outputdir "$(RPTDIR)"
GZCMD = tar -czf
ZIPCMD = zip
SHACMD = sha256sum
SLOCCMD = cloc --by-file --xml --exclude-dir="vendor" --include-lang="Go"
XUCMD = go2xunit


all: debug dep format lint test build 

clean:
	rm -rf $(RPTDIR)
	$(GOCLEAN) $(PKG)

dep:
	$(GODEP) $(PKG)/...

format:
	$(GOFMT) "$(PKGDIR)"

lint: dep 
	$(GOLINT) "$(PKGDIR)" | tee "$(RPTDIR)/lint.out"

test: clean dep
	$(GOTEST) $$(go list "$(PKG)/..." | grep -v /vendor/) | tee "$(RPTDIR)/test.out"

setup-dirs:
	mkdir -p "$(RPTDIR)"

build:  dep
# todo: need to fix "dep" and add it as a dependency step
	$(GOBUILD) 
# $(RPTDIR)/build-$@.out"

debug:
	$(info MD=$(MD))
	$(info WD=$(WD))
	$(info PKG=$(PKG))
	$(info PKGDIR=$(PKGDIR))
	$(info VERSION=$(VERSION))
	$(info COMMIT_ID=$(COMMIT_ID))
	$(info BUILD_TAGS=$(BUILD_TAGS))
	$(info CMDS=$(CMDS))
	$(info BUILD_TARGETS=$(BUILD_TARGETS))
	$(info INSTALL_TARGETS=$(INSTALL_TARGETS))