
linux:
	mkdir -p build/linux
	GOOS=linux go build -o build/linux/kepler-supply-api .