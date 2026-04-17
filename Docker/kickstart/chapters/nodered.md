---
title: Node-RED with Docker
---

# Node-RED with Docker

In this chapter you will learn the basics of running Node-RED inside a Docker container, building custom images, and composing a Node-RED stack with MongoDB.

> **Tasks:**
>
> - [Task 1: Configuring Node-RED](#task-1-configuring-node-red)
> - [Task 2: Persisting Data](#task-2-persisting-data)
> - [Task 3: Building Custom Images](#task-3-building-custom-images)
> - [Task 4: Building a Node-RED Stack](#task-4-building-a-node-red-stack)
> - [Task 5: Configuring Node-RED to see the new services](#task-5-configuring-node-red-to-see-the-new-services)

To get started, let's run the following in our terminal:

```console
$ docker container run -it -p 1880:1880 --name mynodered nodered/node-red-docker
```

We can now access the Node-RED UI via `http://<hostip>:1880`

## Task 1: Configuring Node-RED

Great! Let's configure the Node-RED container and add a node:

```console
# Open a shell in the container
$ docker container exec -it mynodered /bin/bash

# Once inside the container, npm install the nodes in /data
$ cd /data
$ npm install node-red-node-smooth
$ exit
```

Hit `Ctrl-p` `Ctrl-q` to detach from the container.

## Task 2: Persisting Data

In the last section, you saw a lot of Docker-specific jargon which might be confusing to some. So before you go further, let's clarify some terminology that is used frequently in the Docker ecosystem.

## Task 3: Building Custom Images

Creating a new Docker image, using the public Node-RED images as the base image, allows you to install extra nodes during the build process.

Create a file called Dockerfile with the content:

```dockerfile
FROM nodered/node-red-docker
RUN npm install node-red-node-twitter
```

Run the following command to build the image:

```console
$ docker image build -t mynodered:<tag> .
```

That will create a Node-RED image that includes the wordpos nodes.

## Task 4: Building a Node-RED Stack

In the Docker Swarm Chapter, we linked containers together using the Docker compose file.

For example, if you have a container that provides an MQTT broker container called mybroker, you can run the Node-RED container with the link parameter to join the two:

```console
$ docker container run -it -p 1880:1880 --name mynodered --link mybroker:broker nodered/node-red-docker
```

This will make broker a known hostname within the Node-RED container that can be used to access the service within a flow, without having to expose it outside of the Docker host.

We will now create a Docker compose stack using the Dockerfile from section 1.3.

1. Create a directory called nodered and change to the nodered directory
2. Create a file named `Dockerfile` in this directory with the following code:

```dockerfile
FROM nodered/node-red-docker
RUN npm install node-red-node-twitter
RUN npm install node-red-node-mongodb
```

1. Create a file name `docker-compose.yml`
2. Copy the below text into a file named `docker-compose.yml`

```yaml
version: '3.1'
networks:
  node-red:

services:
 nodered:
   build: .
   ports:
     - "1880:1880"
   volumes:
     - ./data:/data
     - ./public:/home/nol/node-red-static
   links:
    - mongodb:mongodb
   networks:
    - node-red

 mongodb:
   image: mongo
   ports:
     - "27017:27017"
   networks:
     - node-red
```

Run the command from the CLI:

```console
$ docker compose up
```

## Task 5: Configuring Node-RED to see the new services

Now, we will walk through how Node-RED can use the new MongoDB.

Access Node-RED `http://<hostip>:1880`

## Next Steps

For the next step in the tutorial, head over to the [Next Steps](./nextsteps.md) page.
