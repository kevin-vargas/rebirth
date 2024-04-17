
docker buildx build -f Dockerfile.${1}\
	--tag kevargas/nonsense:${1}-v${2} \
	--platform linux/arm64,linux/amd64 \
	--builder container \
	--push .