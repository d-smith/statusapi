bin:
	GOOS=linux go build -o main
	zip statusapi.zip main

clean:
	rm -f statusapi.zip
