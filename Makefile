
build-image-%: cmd/%/Dockerfile
	docker build -f $< .
