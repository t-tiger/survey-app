# Survey application

This project contains both server and client applications.

Server application is written in Go, client application is written in TypeScript and uses Next.js as a framework.

## How to start

Following command involves building docker image and running docker containers.

This command requires several minutes (perhaps close to 10 minutes) to finish because both server and client applications have to install dependent libraries and build source codes.

```bash
$ make build-and-start
```

Then open `http://localhost:3000`.

## Remove docker containers

Since above command starts several docker containers, please run following command to remove containers after using applications.

```bash
$ make remove-containers
```

## Entity–relationship diagram

![Entity–relationship diagram](./ER.png)