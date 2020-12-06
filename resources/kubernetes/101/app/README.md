# Guide
This folder contains the source code to build this image, that is used for this Kubernetes basic yaml examples.
## Building the image
To build the image run:   
```
sudo /bin/bash build.sh
```
The image is located in Docker Hub and you can access it in next form:   
```
YOUR_USER/python-flask-distroless
```
## Running the image
t
To run a container based on your image run:   
```
sudo docker run -it -d YOUR_USER/python-flask-distroless
```
Then access your image with the next command:   
```
curl http://localhost:5000
```