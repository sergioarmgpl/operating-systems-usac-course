docker login $1
docker build --no-cache -t $1/python-flask-distroless .