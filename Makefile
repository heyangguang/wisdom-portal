## watch for trailing spaces!
THIS_MAKEFILE=$(lastword $(MAKEFILE_LIST))
REPOROOT = $(abspath  $(dir $(THIS_MAKEFILE)))

export GOPATH := $(REPOROOT)
PATH := $(GOPATH)/bin:$(PATH)

ifdef GOROOT
  #$(error Define GOROOT prior to running make)
  PATH := $(GOROOT)/bin:$(PATH)
endif

PATH := /usr/local/go/bin/:$(PATH)

# BUILDVER is set by OBS (rpm build system)
BUILDVER ?= MANUALMAKEFILE

#when runs on build host - actual golang compiler name is 'golang-go', not 'go'.
#GOCMD ?= golang-go
GOCMD = go

GOCLEAN := $(GOCMD) clean

.PHONY: build
build:
	mkdir -p $(REPOROOT)/bin
	cd $(REPOROOT); $(GOCMD) install

.PHONY: clean
clean:
	@rm -rf $(REPOROOT)/bin
	@rm -rf $(REPOROOT)/pkg