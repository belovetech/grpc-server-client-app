.PHONY: gen run

GEN_PATH = gen/hello

run: gen
	@echo "Running client"
	python3 client.py

gen:
	@echo "Generating Python proto files"
	mkdir -p $(GEN_PATH)
	python -m grpc_tools.protoc -I. --python_out=$(GEN_PATH) --grpc_python_out=$(GEN_PATH) hello.proto

env:
	@echo "Creating virtual environment"
	python3 -m venv .venv


activate:
	@echo "Activating virtual environment"
	source .venv/bin/activate


install:
	@echo "Installing dependencies"
	pip install -r requirements.txt
