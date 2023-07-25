build:
	go build . && docker build -t call-go .

run:
	docker run -p 8080:8080 call-go > logfile

