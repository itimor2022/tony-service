build:
	docker build --platform=linux/amd64 -t tangsengdaodaoserver .
push:
	docker tag tangsengdaodaoserver itimor2022/tangsengdaodaoserver:zy1.2
	docker push itimor2022/wukongim/wukongchatserver:zy1.2
deploy:
	docker build --platform=linux/amd64 -t tangsengdaodaoserver .
	docker tag tangsengdaodaoserver itimor2022/tangsengdaodaoserver:zy1.2
	docker push itimor2022/tangsengdaodaoserver:zy1.2
run-dev:
	docker-compose build;docker-compose up -d
stop-dev:
	docker-compose stop
env-test:
	docker-compose -f ./testenv/docker-compose.yaml up -d 