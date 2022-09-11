### Test app with docker
    docker volume create scanner
    docker build -t scanner:dev .
    docker run --rm --network host --mount source=scanner,target=/var/tmp/htolic-scanner scanner:dev 192.168.50.51

### Push image to dockerhub
    docker tag scanner:dev htolic/scanner:dev
    docker push

### Prepare minikube
    minikube start --driver=docker

### Run scanner inside minikube
    kubectl apply -f scanner.yaml
