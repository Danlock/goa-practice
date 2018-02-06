#! /usr/bin/make
#


#Using gopkg for major versioning of goagen outside of dep tool
DEPEND= gopkg.in/goadesign/goa.v1 \
	gopkg.in/goadesign/goa.v1/goagen

CURRENT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: depend clean generate generate-controllers fix-vendor build

depend:
	@echo "Downloading dependencies..."
	@go get $(DEPEND)
	@dep ensure

clean:
	@echo "Cleaning old generated files and renaming vendor to avoid goagen conflicts..."
	@rm -rf generated
	@mv vendor/ temp-vendor/

generate:
	@echo "Generating..."
	@goagen app     -d github.com/danlock/goa-practice/design -o generated
	@goagen swagger -d github.com/danlock/goa-practice/design -o generated/public
	@goagen schema -d github.com/danlock/goa-practice/design -o generated/public
	@goagen client  -d github.com/danlock/goa-practice/design -o generated

#regenerate main and controller files, delete duplicate controllers, only needed on initial run because main doesnt get regenerated properly
generate-main:
	@echo "Generating scaffolding for main..."
	@goagen main     -d github.com/danlock/goa-practice/design
	@goagen controller     -d github.com/danlock/goa-practice/design -o controller --app-pkg ../generated/app --regen
	@ls controller/ | xargs rm

generate-controllers:
	@echo "Regenerating scaffolding for controllers..."
	@goagen controller     -d github.com/danlock/goa-practice/design -o controller --app-pkg ../generated/app --regen


fix-vendor:
	@echo "Fixing vendor directory..."
	@	@mv temp-vendor/ vendor/ 

build:
	@echo "Building binary..."
	@go build -o bin/goa-practice
