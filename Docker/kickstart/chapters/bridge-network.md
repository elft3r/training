---
title: Bridge Networking
---

# Bridge Networking

In this lab you'll learn how to build, manage, and use **bridge** networks — the most common networking type for containers running on a single Docker host.

> **Tasks:**
>
> - [Task 1: Explore the default bridge network](#task-1-explore-the-default-bridge-network)
> - [Task 2: Run containers on the default bridge](#task-2-run-containers-on-the-default-bridge)
> - [Task 3: Create a user-defined bridge network](#task-3-create-a-user-defined-bridge-network)
> - [Task 4: Test DNS-based service discovery](#task-4-test-dns-based-service-discovery)
> - [Task 5: Network isolation between bridges](#task-5-network-isolation-between-bridges)
> - [Task 6: Configure port mapping for external access](#task-6-configure-port-mapping-for-external-access)
> - [Task 7: Cleanup](#task-7-cleanup)

## Task 1: Explore the default bridge network

Every Docker installation comes with three pre-built networks. List them with `docker network ls`.

```console
$ docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
1befe23acd58        bridge              bridge              local
726ead8f4e6b        host                host                local
ef4896538cc7        none                null                local
```

The **bridge** network uses the *bridge* driver and is scoped locally (single-host only). It's important to note that the network name and the driver name are both "bridge", but they are not the same thing — one is a specific network, the other is the driver (template) used to create it.

Under the hood, the *bridge* driver creates a Linux bridge (virtual switch) called **docker0** on the host. You can see it with the `ip` command.

> **NOTE:** In order to execute this command, you need to be in the Linux system, that is running Docker.

```console
$ ip addr show docker0
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN
    link/ether 02:42:f1:7f:89:a6 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 scope global docker0
       valid_lft forever preferred_lft forever
```

The **docker0** bridge acts as the gateway (`172.17.0.1`) for all containers connected to the default **bridge** network.

Inspect the **bridge** network to see its configuration.

```console
$ docker network inspect bridge
[
    {
        "Name": "bridge",
        "Id": "1befe23acd58...",
        "Scope": "local",
        "Driver": "bridge",
        "IPAM": {
            "Config": [
                {
                    "Subnet": "172.17.0.0/16",
                    "Gateway": "172.17.0.1"
                }
            ]
        },
        "Containers": {},
        ...
    }
]
```

## Task 2: Run containers on the default bridge

The **bridge** network is the default for new containers. If you don't specify `--network`, your container lands here.

Start two Alpine containers.

```console
$ docker container run -dit --name c1 alpine sh
$ docker container run -dit --name c2 alpine sh
```

> **NOTE:** We use Alpine because it includes `ping` out of the box — no need to install extra packages.

Verify both containers are connected to the **bridge** network.

{% raw %}
```console
$ docker network inspect bridge --format '{{range .Containers}}{{.Name}}: {{.IPv4Address}}{{"\n"}}{{end}}'
c1: 172.17.0.2/16
c2: 172.17.0.3/16
```
{% endraw %}

Test connectivity **by IP address** between the two containers.

```console
$ docker container exec c1 ping -c 3 172.17.0.3
PING 172.17.0.3 (172.17.0.3): 56 data bytes
64 bytes from 172.17.0.3: seq=0 ttl=64 time=0.112 ms
64 bytes from 172.17.0.3: seq=1 ttl=64 time=0.098 ms
64 bytes from 172.17.0.3: seq=2 ttl=64 time=0.102 ms
```

Containers on the default bridge can reach each other by IP. Now try reaching the other container **by name**.

```console
$ docker container exec c1 ping -c 3 c2
ping: bad address 'c2'
```

**This fails.** The default bridge network does **not** provide DNS-based service discovery. Containers can only communicate using IP addresses, which is fragile and impractical. This is one of the key reasons Docker recommends using **user-defined bridge networks** instead.

Also verify that containers can reach the internet.

```console
$ docker container exec c1 ping -c 3 docker.com
PING docker.com (141.193.213.20): 56 data bytes
64 bytes from 141.193.213.20: seq=0 ttl=37 time=15.2 ms
...
```

Clean up the default bridge containers.

```console
$ docker container rm -f c1 c2
```

## Task 3: Create a user-defined bridge network

User-defined bridge networks are the **recommended** approach for container networking. They provide:
- **Automatic DNS resolution** between containers (by name)
- **Better isolation** from other containers
- **Configurable subnets** and IP ranges

Create a user-defined bridge network.

```console
$ docker network create my_bridge
a1b2c3d4e5f6...
```

Inspect the new network.

{% raw %}
```console
$ docker network inspect my_bridge --format '{{(index .IPAM.Config 0).Subnet}}'
172.18.0.0/16
```
{% endraw %}

Docker automatically assigned a subnet. You can also specify one explicitly.

```console
$ docker network create \
  --subnet=10.0.0.0/24 \
  --gateway=10.0.0.1 \
  custom_bridge
```

Verify both networks exist.

```console
$ docker network ls --filter driver=bridge
NETWORK ID          NAME                DRIVER              SCOPE
1befe23acd58        bridge              bridge              local
a1b2c3d4e5f6        my_bridge           bridge              local
d7e8f9a0b1c2        custom_bridge       bridge              local
```

Remove the `custom_bridge` network since we won't use it further.

```console
$ docker network rm custom_bridge
```

## Task 4: Test DNS-based service discovery

Now start two containers on the **user-defined** `my_bridge` network.

```console
$ docker container run -dit --name web --network my_bridge nginx:alpine
$ docker container run -dit --name client --network my_bridge alpine sh
```

Test connectivity **by container name**.

```console
$ docker container exec client ping -c 3 web
PING web (172.18.0.2): 56 data bytes
64 bytes from 172.18.0.2: seq=0 ttl=64 time=0.089 ms
64 bytes from 172.18.0.2: seq=1 ttl=64 time=0.105 ms
64 bytes from 172.18.0.2: seq=2 ttl=64 time=0.100 ms
```

**It works!** Docker's built-in DNS server automatically resolves container names to their IP addresses on user-defined networks. This is how real-world microservices discover each other — by name, not by IP.

You can even use `wget` to reach the NGINX web server by name.

```console
$ docker container exec client wget -qO- http://web
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
...
```

## Task 5: Network isolation between bridges

Containers on **different** bridge networks are isolated from each other by default. This is a fundamental Docker security feature.

Start a container on the default **bridge** network.

```console
$ docker container run -dit --name isolated alpine sh
```

Try to ping the `web` container (which is on `my_bridge`) from `isolated` (which is on the default `bridge`).

```console
$ docker container exec isolated ping -c 3 172.18.0.2
PING 172.18.0.2 (172.18.0.2): 56 data bytes

--- 172.18.0.2 ping statistics ---
3 packets transmitted, 0 packets received, 100% packet loss
```

**No connectivity.** The two bridge networks are completely isolated from each other.

### Connecting a container to multiple networks

You can attach a container to additional networks using `docker network connect`.

```console
$ docker network connect my_bridge isolated
```

Now `isolated` is connected to **both** the default `bridge` and `my_bridge`. Verify.

```console
$ docker container exec isolated ping -c 3 web
PING web (172.18.0.2): 56 data bytes
64 bytes from 172.18.0.2: seq=0 ttl=64 time=0.110 ms
...
```

It can now reach containers on `my_bridge` by name. To remove the connection:

```console
$ docker network disconnect my_bridge isolated
```

Clean up the isolated container.

```console
$ docker container rm -f isolated
```

## Task 6: Configure port mapping for external access

Containers on a bridge network are not accessible from outside the Docker host by default. To expose a container's service externally, use **port mapping** with the `-p` flag.

The `web` container from Task 4 is already running NGINX on port 80, but it's only reachable from within the `my_bridge` network. Let's fix that.

Stop and re-create the `web` container with port mapping.

```console
$ docker container rm -f web
$ docker container run -d --name web --network my_bridge -p 8080:80 nginx:alpine
```

This maps port **8080** on the Docker host to port **80** inside the container.

Verify the port mapping.

```console
$ docker container ls --filter name=web
CONTAINER ID   IMAGE          COMMAND                  PORTS                  NAMES
b747d43fa277   nginx:alpine   "/docker-entrypoint.…"   0.0.0.0:8080->80/tcp   web
```

The `0.0.0.0:8080->80/tcp` mapping means port 8080 on **all host interfaces** forwards to port 80 in the container.

Test external access using `curl` from the Docker host.

```console
$ curl localhost:8080
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
...
```

You can also access it from a web browser by navigating to `http://<your-docker-host-ip>:8080`.

> **NOTE:** Port mapping uses **DNAT** (Destination NAT) under the hood — inbound connections arriving on host port 8080 are redirected to the container's port 80. Separately, **outbound** traffic from containers to the internet uses **SNAT/masquerade**, so it appears to originate from the host's IP address. This outbound masquerading applies to all bridge-connected containers, regardless of whether port mapping is configured.

## Task 7: Cleanup

Remove all containers and the custom network.

```console
$ docker container rm -f web client
$ docker network rm my_bridge
```

Verify cleanup.

```console
$ docker container ls -a --filter name=web --filter name=client
CONTAINER ID   IMAGE   COMMAND   CREATED   STATUS   PORTS   NAMES

$ docker network ls --filter driver=bridge
NETWORK ID          NAME                DRIVER              SCOPE
1befe23acd58        bridge              bridge              local
```

## Key Takeaways

| | Default Bridge | User-Defined Bridge |
|---|---|---|
| **DNS resolution** | No — containers must use IP addresses | Yes — containers can reach each other by name |
| **Isolation** | All containers share the same default network | Containers are isolated by network |
| **Recommended** | No — use only for quick tests | Yes — use for all real workloads |
| **Configuration** | Cannot be customized | Custom subnets, IP ranges, and gateways |
| **Connect/disconnect** | Cannot disconnect running containers from primary (default) network | Live connect/disconnect with `docker network connect/disconnect` |

**Best practices:**
- Always create user-defined bridge networks for your applications
- Use container names for service discovery instead of hard-coded IPs
- Use port mapping (`-p`) only for services that need external access
- Use separate networks to isolate groups of containers (e.g., frontend vs backend)

## Next Steps
For the next step in the tutorial, head over to [Docker Secrets](./secrets.md)
