.PHONY: build serve clean

build:
	go run .

serve: build
	python3 -m http.server 8000 --directory dist

clean:
	rm -rf dist
