build:
	docker build --platform=linux/amd64 -t mft .
push:
	docker tag mft itimor2022/mft:1.3
	docker push itimor2022/mft:1.3
deploy:
	docker build --platform=linux/amd64 -t mft .
	docker tag mft itimor2022/mft:1.3
	docker push itimor2022/mft:1.3
run-dev:
	docker-compose build;docker-compose up -d
stop-dev:
	docker-compose stop
env-test:
	docker-compose -f ./testenv/docker-compose.yaml up -d 