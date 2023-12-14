build:
	go build -ldflags="-s -w" goctlwt.go
	$(if $(shell command -v upx), upx goctlwt)

mac:
	GOOS=darwin go build -ldflags="-s -w" -o goctlwt-darwin goctlwt.go
	$(if $(shell command -v upx), upx goctlwt-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o goctlwt.exe goctlwt.go
	$(if $(shell command -v upx), upx goctlwt.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o goctlwt-linux goctlwt.go
	$(if $(shell command -v upx), upx goctlwt-linux)

image:
	docker build --rm --platform linux/amd64 -t superhuan/goctlwt:$(version) .
	docker tag superhuan/goctlwt:$(version) superhuan/goctlwt:latest
	docker push superhuan/goctlwt:$(version)
	docker push superhuan/goctlwt:latest
	docker build --rm --platform linux/arm64 -t superhuan/goctlwt:$(version)-arm64 .
	docker tag superhuan/goctlwt:$(version)-arm64 kevinwan/goctlwt:latest-arm64
	docker push superhuan/goctlwt:$(version)-arm64
	docker push superhuan/goctlwt:latest-arm64
