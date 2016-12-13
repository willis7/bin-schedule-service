# VERSION:		  0.1
# DESCRIPTION:	Runs a simple docker microservice
# USAGE:
#
# # Build image
# docker build -t trash-collection-svc .
#
# # Run container
# docker run -it --rm --name my-trash-collection-svc -p 8080:8080 trash-collection-svc
#
FROM golang:1.7.3-onbuild

EXPOSE 8080
