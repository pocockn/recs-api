<h1 align="center">recs-api</h1>

<p align="center">
   A GO REST API to rate and review recommendations from a Spotify playlist.
<p>
# recs-api

# recs-api

A GO REST API to rate and review recommendations from Spotify. Songs get pulled from Spotify into the DB from another microservice (currently in dev)

## Getting Started

```
make run
```

### Prerequisites

What things you need to install the software and how to install them

```
make install
```

### Installing

```
make install
```

## Running the tests

```
make test
```

## Deployment

Each commit to master is built and tested within Gitlab, only tagged builds are released on Github and on Dockerhub.

## Built With

* [Go](https://golang.org/) - Main language
* [Gitlab](https://gitlab.com/) - Dependency Management
* [Echo](https://github.com/labstack/echo) - HTTP Framework
* [GORM](https://gorm.io/) - ORM Library

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/pocockn/recs-api/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
