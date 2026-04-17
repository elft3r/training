---
---

# MongoDB with Node-RED

In this lab you'll learn how to deploy MongoDB with Docker.

> **Tasks:**
>
> - [Task 1: Building Custom Images](#task-1-building-custom-images)
> - [Task 2: Building a Node-RED Stack](#task-2-building-a-node-red-stack)
> - [Task 3: Configuring Node-RED to see the new services](#task-3-configuring-node-red-to-see-the-new-services)
> - [Task 4: Replicated MongoDB](#task-4-replicated-mongodb)
> - [Task 5: Cleanup](#task-5-cleanup)

In this lab the terms *service task* and *container* are used interchangeably.
In all examples in the lab a *service tasks* is a container that is running as
part of a service.

## Overview Node-RED, MongoDB, and Docker
Node-RED is a flow-based programming tool based on NodeJs. The interface allows for easily dragging and dropping nodes and wiring them together. The real power of Node-RED is when it is combined with Docker. Docker allows easily to provison services which Node-RED can connect to.

In this chapter we will cover the basics of running Node-RED inside of a Docker container. Once we have accomplished the deplyoment of Node-RED with Docker we will then add a MongoDB database which we will connect to from inside Node-RED.

To get started, let's run the following in our terminal:
```console
$ docker container run -it -p 1880:1880 --name mynodered nodered/node-red-docker
```

We can now access the Node-RED UI via `http://<hostip>:1880`


## Task 1: Building Custom Images

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


## Task 2: Building a Node-RED Stack
In the Docker Swarm Chapter, we linked containers together using the Docker compose file.

For example, if you have a container that providesis your APP and you would like to connect it to a MongoDB to persist the data. In this example, we link the Node-RED container with MongoDB so node-red can store data:

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
   volumes:
      - /path/to/mongodb-persistence:/bitnami
   ports:
     - "27017:27017"
   networks:
     - node-red
```

* Run the command from the CLI: `docker compose up`


## Task 3: Configuring Node-RED to see the new services
Now, we will walk through how Node-RED can use the new MongoDB

Access Node-RED `http://<hostip>:1880`


## Task 4: Replicated MongoDB
To setup a MongoDB cluster is quite easy. To understand how a MongoDB cluster works review the [MongoDB replication documentation](https://docs.mongodb.com/manual/replication/)

```yaml
services:
  mongodb-primary:
    image: 'bitnami/mongodb:latest'
    environment:
      - MONGODB_REPLICA_SET_MODE=primary
      - MONGODB_ROOT_PASSWORD=password123
      - MONGODB_REPLICA_SET_KEY=replicasetkey123
    volumes:
      - 'mongodb_master_data:/bitnami'

  mongodb-secondary:
    image: 'bitnami/mongodb:latest'
    depends_on:
      - mongodb-primary
    environment:
      - MONGODB_REPLICA_SET_MODE=secondary
      - MONGODB_PRIMARY_HOST=mongodb-primary
      - MONGODB_PRIMARY_PORT_NUMBER=27017
      - MONGODB_PRIMARY_ROOT_PASSWORD=password123
      - MONGODB_REPLICA_SET_KEY=replicasetkey123

  mongodb-arbiter:
    image: 'bitnami/mongodb:latest'
    depends_on:
      - mongodb-primary
    environment:
      - MONGODB_REPLICA_SET_MODE=arbiter
      - MONGODB_PRIMARY_HOST=mongodb-primary
      - MONGODB_PRIMARY_PORT_NUMBER=27017
      - MONGODB_PRIMARY_ROOT_PASSWORD=password123
      - MONGODB_REPLICA_SET_KEY=replicasetkey123

volumes:
  mongodb_master_data:
    driver: local
```

In the above example we can easily scale our MongoDB cluster:

```console
$ docker compose scale mongodb-primary=1 mongodb-secondary=3 mongodb-arbiter=1
```

The above command scales up the number of secondary nodes to 3

## Task 5: Cleanup

This is the final cleanup where we will delete all the containers, networks, and volumes from your Docker Host. **Only if you want to**

1. First stop the running MongoDB stack

```console
$ docker compose rm
```

2. Prune Docker of all containers, images, networks, and basically everything else.

```console
$ docker system prune
WARNING! This will remove:
        - all stopped containers
        - all networks not used by at least one container
        - all dangling images
        - all build cache
Are you sure you want to continue? [y/N]
```


## Next Steps

For additional resources, head over to the [Next Steps](./nextsteps.md) page.
