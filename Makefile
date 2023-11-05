build :
	docker build -t template-project .

run :
	docker run --name template -p 8080:8080 -itd template-project