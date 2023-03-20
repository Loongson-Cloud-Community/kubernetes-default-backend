BIN=defaultbackend
VERSION=1.0
IMAGE=cr.loongnix.cn/kubernetes-sigs/defaultbackend:${VERSION}
$(BIN):
	CGO_ENABLED=0 go build -o ${BIN} ./main.go 

image:${BIN}
	docker build \
		--build-arg BIN=${BIN} \
		-t ${IMAGE} \
		.

