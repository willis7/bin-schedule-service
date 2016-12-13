# Bin Schedule Service

A RESTful API for letting consumers know when their bins are being collected.

**NOTE:** currently only works for East Devon

### Build Status
[![Go Report Card](https://goreportcard.com/badge/github.com/willis7/bin-schedule-service)](https://goreportcard.com/report/github.com/willis7/bin-schedule-service)
[![Build Status](https://travis-ci.org/willis7/bin-schedule-service.svg?branch=master)](https://travis-ci.org/willis7/bin-schedule-service)


### Architectue

Here's how the application is structured.

* **base**: base contains the basic (core) functionality for the api. Base contains all the source code that cannot be categorized into a submodule.


### Example Response

    {
        "Postcode": "EX5 3DX",
        "RecyclingDate": "Friday 25 November 2016",
        "RubbishDate": "Friday 25 November 2016"
    }
