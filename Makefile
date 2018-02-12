#! /usr/bin/make

#Using gopkg for major versioning of goagen outside of dep tool
DEPEND= gopkg.in/goadesign/goa.v1 \
	gopkg.in/goadesign/goa.v1/goagen

all: depend regen

regen: clean generate fix-vendor

depend:
	@echo "Go getting goagen..."
	@go get $(DEPEND)
	@echo "Dep ensuring packages..."
	@dep ensure

clean:
	@echo "Cleaning old generated files..."
	@rm -rf generated

generate:
	@echo "Generating and renaming vendor to avoid goagen conflicts......"
	@mv vendor/ temp-vendor/
	@goagen app     -d github.com/danlock/goa-practice/design -o generated || $(MAKE) fix-vendor
	@goagen swagger -d github.com/danlock/goa-practice/design -o generated/public
	@goagen schema -d github.com/danlock/goa-practice/design -o generated/public
	@goagen client  -d github.com/danlock/goa-practice/design -o generated

fix-vendor:
	@mv temp-vendor/ vendor/ 

build:
	@echo "Building binary..."
	@go build -o bin/goa-practice

#regenerate main and controller files, these files need to be deleted or they will cause errors
scaffolding:
	@$(MAKE) clean
	@$(MAKE) generate
	@echo "Generating scaffolding..."
	@goagen main -d github.com/danlock/goa-practice/design -o generated/main
	@$(MAKE) fix-vendor

delete-scaffolding:
	@echo "Deleting scaffolding..."
	@rm -rf generated/main
