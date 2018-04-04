<h1 align="center">
  <img src="https://i.imgur.com/EY4EVJF.png" width="300"></a>
  <br>
  ResamVitalized
</h1>
<p align="center">
  <a href="https://travis-ci.org/ResamVi/resamvitalized"> <img src="https://travis-ci.org/ResamVi/resamvitalized.svg?branch=master"></a>
  <a href="https://codeclimate.com/github/ResamVi/resamvitalized/maintainability"><img src="https://api.codeclimate.com/v1/badges/d3dd3d9b21beb3ffd798/maintainability" /></a>
</p>

A re-engineered version of my old website [https://resamvi.de/](https://resamvi.de/) using more advanced technologies.

## Getting Started

To get the website running:
```
npm install
npm build
docker-compose up
```

Afterwards visit `http://localhost:9001/` in your favourite browser.

This sets up all the required dependencies, builds the website from scratch and
runs all services (database, server, nginx) to run the website.

### Prerequisites

* [Node.js](https://nodejs.org/en/) - Comes with `npm` to install the dependencies.
* [docker-compose](https://docs.docker.com/compose/install/) - To run the container.

### Developing

For development install the dependencies with

```
npm install
```

To build the website either run `npm build` or
```
npm run watch
```
to start the webpack watcher that will re-build the website on each change you do in the source code.

```
docker-compose up
```

Will run the following three services:
* nginx on port 9001
* mongodb on port 27017
* go on port 8080


## Built With

[Vue.js](https://vuejs.org/)

<img src="https://vuejs.org/images/logo.png" width="100">

[nginx](https://nginx.org/en/)

<img src="https://nginx.org/nginx.png">

[go](https://golang.org/)

<img src="https://upload.wikimedia.org/wikipedia/commons/2/23/Golang.png" width="200">

## Authors

* **Julien Midedji** - [ResamVi](https://github.com/ResamVi)

## License

This project is licensed under the MPL  - see the [LICENSE](LICENSE) file for details
