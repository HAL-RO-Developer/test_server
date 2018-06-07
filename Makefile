build:
	docker build -t test_server .
run:
	docker run -it --rm -p 8080:8080 test_server
