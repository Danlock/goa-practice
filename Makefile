#! /usr/bin/make

#Using gopkg for major versioning of goagen outside of dep tool
DEPEND= gopkg.in/goadesign/goa.v1 \
	gopkg.in/goadesign/goa.v1/goagen

all: depend clean generate fix-vendor

regenerate: clean generate fix-vendor

depend:
	@echo "Downloading dependencies..."
	@go get $(DEPEND)
	@dep ensure

clean:
	@echo "Cleaning old generated files and renaming vendor to avoid goagen conflicts..."
	@rm -rf generated
	@rm -rf temp-vendor/ #Only would be made as a result of previous botched builds
	@mv vendor/ temp-vendor/

generate:
	@echo "Generating..."
	@goagen app     -d github.com/danlock/goa-practice/design -o generated
	@goagen swagger -d github.com/danlock/goa-practice/design -o generated/public
	@goagen schema -d github.com/danlock/goa-practice/design -o generated/public
	@goagen client  -d github.com/danlock/goa-practice/design -o generated

fix-vendor:
	@echo "Fixing vendor directory..."
	@mv temp-vendor/ vendor/ 

build:
	@echo "Building binary..."
	@go build -o bin/goa-practice

#regenerate main and controller files, these files need to be deleted/merged or they will cause errors
generate-scaffolding:
	@$(MAKE) clean
	@$(MAKE) generate
	@echo "Generating scaffolding..."
	@goagen main -d github.com/danlock/goa-practice/design -o generated/main
	@$(MAKE) fix-vendor

delete-scaffolding:
	@echo "Deleting scaffolding..."
	@rm -rf generated/main
