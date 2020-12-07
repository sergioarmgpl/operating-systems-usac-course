docker login -u $1
docker build -t $1/python-flask-distroless .
docker push $1/python-flask-distroless
