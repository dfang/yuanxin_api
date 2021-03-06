.PHONY: install test build serve clean pack deploy ship local_image

TAG=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test:
	go test ./...

vet:
	go vet ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o news .

serve: build
	./news

clean:
	rm ./news

local_image:
	GOOS=linux make build
	docker build -t "dfang/yuanxin:latest" .

start_containers: local_image
	docker-compose up --force-recreate -d

pack:
	GOOS=linux make build
	docker build -t dfang/yuanxin:$(TAG) .
	docker tag dfang/yuanxin:$(TAG) dfang/yuanxin:latest

upload:
	git push
	docker push dfang/yuanxin:$(TAG)
	docker push dfang/yuanxin:latest

pull_code: 
	git pull

pull_image:
	docker pull dfang/yuanxin:$(TAG)
	docker tag  dfang/yuanxin:$(TAG) dfang/yuanxin:latest

make pull:
	make pull_code
	make pull_image

deploy:
	# envsubst < k8s/deployment.yml | kubectl apply -f -
	ssh root@39.108.66.107 "cd /root/yuanxin_api; make pull; docker-compose up -d"


ship: test pack upload clean deploy

