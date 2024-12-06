build:
	docker build --platform=linux/amd64 -t tangsengdaodaoserver .
push:
	docker tag tangsengdaodaoserver itimor2022/tangsengdaodaoserver:zy1.4
	docker push itimor2022/wukongim/wukongchatserver:zy1.4
deploy:
	docker build --platform=linux/amd64 -t tangsengdaodaoserver .
	docker tag tangsengdaodaoserver itimor2022/tangsengdaodaoserver:zy1.4
	docker push itimor2022/tangsengdaodaoserver:zy1.4
run-dev:
	docker-compose build;docker-compose up -d
stop-dev:
	docker-compose stop
env-test:
	docker-compose -f ./testenv/docker-compose.yaml up -d 