# VERSION:		  0.1
# DESCRIPTION:	Runs a simple docker microservice
# USAGE:
#
# # Build image
# docker build -t bin-schedule-svc .
#
# # Run container
# docker run -it --rm --name my-bin-schedule-svc -p 8080:8080 bin-schedule-svc
#
FROM golang:1.7.3-onbuild

EXPOSE 8080
