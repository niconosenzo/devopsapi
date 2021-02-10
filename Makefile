# Manually create the app binary
.PHONY: install
install:
	go build -o api .

# Run the suit (unit) tests
.PHONY: test
test:
	go test ./tests/ -v

# Create the container image to my public repo
.PHONY: build
build:
	docker build -t devops-api-initial .
	docker tag devops-api-initial quay.io/nnosenzo/devopsapi:latest
	docker push quay.io/nnosenzo/devopsapi:latest

# Deploy API pod and service to a K8s cluster
.PHONY: deploy
deploy:
	kubectl create deployment devopsapi --image=quay.io/nnosenzo/devopsapi:latest && kubectl create service clusterip devopsapi --tcp=3000:3000
