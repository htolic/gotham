### Test app with docker
    docker volume create scanner
    docker build -t scanner:dev .
    docker run --rm --network host --mount source=scanner,target=/home/gorunner/scanner scanner:dev 192.168.50.51

### Push image to dockerhub
    docker tag scanner:dev htolic/scanner:dev
    docker push htolic/scanner:dev

### Prepare minikube
    minikube start --driver docker
    minikube mount temp:/data/scanner --uid 1001 --gid 1001

### Run scanner inside minikube
    kubectl apply -f scanner.yaml
