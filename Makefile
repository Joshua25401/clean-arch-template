build-local :
	docker build -t template-project .

run-local : build-local
	docker run --name template -p 8080:8080 -itd template-project

prune-local:
	docker container stop template
	docker container rm template
	docker rmi template-project
	docker system prune -f

run-remote:
	docker run --name template -p 8080:8080 -itd ghcr.io/joshua25401/clean-arch-template/template:develop

prune-remote:
	docker container stop template
	docker container rm template
	docker rmi ghcr.io/joshua25401/clean-arch-template/template:develop 
	docker system prune -f