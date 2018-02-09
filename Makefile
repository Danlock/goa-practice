#! /usr/bin/make
#


#Using gopkg for major versioning of goagen outside of dep tool
DEPEND= gopkg.in/goadesign/goa.v1 \
	gopkg.in/goadesign/goa.v1/goagen

all: depend clean generate fix-vendor build

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

fix-vendor:
	@echo "Fixing vendor directory..."
	@	@mv temp-vendor/ vendor/ 

build:
	@echo "Building binary..."
	@go build -o bin/goa-practice

#regenerate main and controller files, these files need to be deleted or they will conflict with real main and controllers outside of generated/
generate-scaffolding:
	@echo "Generating scaffolding..."
	@goagen main     -d github.com/danlock/goa-practice/design -o generated/main

delete-scaffolding:
	@echo "Deleting scaffolding..."
	@rm -rf generated/main
