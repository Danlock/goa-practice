Workflow:

Download go's official dep tool. https://github.com/golang/dep/releases
run make

Note that if main.go is deleted and recreated, make will fail because main.go is mostly there as a guide and doesnt generate the proper import paths.
