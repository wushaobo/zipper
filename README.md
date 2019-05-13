# Zipper

- [Docker Image](https://hub.docker.com/r/wushaobo/zipper/)
- Demo [page](http://bl.ocks.org/wushaobo/65699d67346e1d0aa951828ff6a2121e), [source](https://gist.github.com/wushaobo/65699d67346e1d0aa951828ff6a2121e)
- [CI](https://travis-ci.org/wushaobo/zipper)

## Features
It accepts file urls in the request and return a key to allow the client to start downloading a zip of those files immediately.

- Sync download immediately, no waiting
- Fit for modern browser and IE
- Handle duplicate file names by renaming
- Support folders, including reserved empty folders
- Not support compression in zip


## How to run

### Requirements

- docker
- docker-compose


### Deployment Architecture

- redis
- zipper

### Configuration

Create a config.env, with some config for http auth and redis connection.

```
HTTP_LISTEN_PORT=24080
HTTP_APP_KEY=INSECURE-APP-KEY
HTTP_SECRET_KEY=INSECURE-SECRET-KEY
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=1
```

### Run

You can run it with this kind of docker-compose file.

```
version: '3'

services:
  zipper:
    image: wushaobo/zipper:latest
    command: /opt/zipper/run.sh
    ports:
      - "24080:80"
    env_file:
      - config.env
    volumes:
      - "/docker/logs/zipper:/var/log/zipper"
```


## Contribute
Suggestions and issues are welcomed.

### Development
Run containers as daemon to start the services of dependencies.

```
./devtool.sh start-deps
```

Rebuild and run zipper container in foreground.

```
./devtool.sh build-and-run
```

### Build

[The CI of this repo](https://travis-ci.org/wushaobo/zipper) is on travis-ci.org .

It will build docker image with the following command, 

```
./build.sh build-image
```

then push to docker registry with the following command,

```
./build.sh push-to-registry
```
