---
title: "8 - Concurrency"
parent: 12-Factor App
grand_parent: Docker
nav_order: 9
---


# 8 - Concurrency

Horizontal scalability with the processes model.

The app can be seen as a set of processes of different types
* web server
* worker
* cron

Each process needs to be able to scale horizontally, it can have its own internal multiplexing.

## What does that mean for our application ?

The messageApp only have one type of process (http server), it's doing the multiplexing using Node.js http server.

This process can be easily scalable (stateless process).

[Previous](07_port_binding.md) - [Next](09_disposability.md)
