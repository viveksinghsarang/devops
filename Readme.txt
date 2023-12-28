Pull the images from the dockerhub:
docker pull viveksinghsarang/node:v1
docker pull viveksinghsarang/golang:v1
docker pull viveksinghsarang/python:v1

If you want to create an image from scratch:-
First pull all the latest images of python, golang and node.
Then install all the dependencies:-

Install this at node image:-
npm install express axios body-parser

Install this at python image:-
pip install flask

Install this at golang image:-
go mod init go_service.go
go get -u github.com/gin-gonic/gin

To run the application:-

docker-compose up
