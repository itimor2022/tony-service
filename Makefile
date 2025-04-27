build:
	docker build --platform=linux/amd64 -t qmjk .
push:
	docker tag qmjk itimor2022/qmjk:1.2
	docker push itimor2022/qmjk:1.2
deploy:
	docker build --platform=linux/amd64 -t qmjk .
	docker tag qmjk itimor2022/qmjk:1.2
	docker push itimor2022/qmjk:1.2
run-dev:
	docker-compose build;docker-compose up -d
stop-dev:
	docker-compose stop
env-test:
	docker-compose -f ./testenv/docker-compose.yaml up -d 