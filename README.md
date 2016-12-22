# Bin Schedule Service

A RESTful API for letting consumers know when their bins are being collected.

**NOTE:** currently only works for East Devon

### Build Status
[![Go Report Card](https://goreportcard.com/badge/github.com/willis7/bin-schedule-service)](https://goreportcard.com/report/github.com/willis7/bin-schedule-service)
[![Build Status](https://travis-ci.org/willis7/bin-schedule-service.svg?branch=master)](https://travis-ci.org/willis7/bin-schedule-service)
[![Coverage Status](https://coveralls.io/repos/github/willis7/bin-schedule-service/badge.svg?branch=master)](https://coveralls.io/github/willis7/bin-schedule-service?branch=master)


### Architectue

Here's how the application is structured.

* **services**: Business logic and calls to external services.
* **models**:  Describes the data model of the application
* **handlers**:  Implements the application’s application handlers
* **routers**: Implements the HTTP request routers for the API


### Example Response

    {
        "Postcode": "EX15 5DX",
        "RecyclingDate": "Friday 25 November 2016",
        "RubbishDate": "Friday 25 November 2016"
    }
