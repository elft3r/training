# Code Block Formatting Comparison

This file compares different approaches for showing shell commands and their output.
The goal is to find the format that is **easiest to copy-paste** while still being clear for learners.

We use the same two examples for each option so you can compare them side by side.

---

## Current Format (baseline)

Command and output mixed in one block, `$` prefix on commands:

```
$ docker image pull alpine
Unable to find image 'alpine:latest' locally
latest: Pulling from library/alpine
88286f41530e: Pull complete
Digest: sha256:f006ecbb824d87947d0b51ab8488634bf69fe4094959d935c0c103f4820a417d
Status: Downloaded newer image for alpine:latest
```

```
$ docker image ls

REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
alpine                  latest              3fd9065eaf02        2 weeks ago         7.8MB
hello-world             latest              f2a91732366c        2 months ago        13.3kB
```

---

## Option A: Separate code blocks

Command in a `bash` block (no `$`), output in a plain block below:

```bash
docker image pull alpine
```

```
Unable to find image 'alpine:latest' locally
latest: Pulling from library/alpine
88286f41530e: Pull complete
Digest: sha256:f006ecbb824d87947d0b51ab8488634bf69fe4094959d935c0c103f4820a417d
Status: Downloaded newer image for alpine:latest
```

```bash
docker image ls
```

```
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
alpine                  latest              3fd9065eaf02        2 weeks ago         7.8MB
hello-world             latest              f2a91732366c        2 months ago        13.3kB
```

---

## Option B: Command block + blockquote output

Command in a `bash` block, output in a quoted block:

```bash
docker image pull alpine
```

> **Output:**
> ```
> Unable to find image 'alpine:latest' locally
> latest: Pulling from library/alpine
> 88286f41530e: Pull complete
> Digest: sha256:f006ecbb824d87947d0b51ab8488634bf69fe4094959d935c0c103f4820a417d
> Status: Downloaded newer image for alpine:latest
> ```

```bash
docker image ls
```

> **Output:**
> ```
> REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
> alpine                  latest              3fd9065eaf02        2 weeks ago         7.8MB
> hello-world             latest              f2a91732366c        2 months ago        13.3kB
> ```

---

## Option C: Console language hint

Same combined block as today, but with `console` syntax highlighting:

```console
$ docker image pull alpine
Unable to find image 'alpine:latest' locally
latest: Pulling from library/alpine
88286f41530e: Pull complete
Digest: sha256:f006ecbb824d87947d0b51ab8488634bf69fe4094959d935c0c103f4820a417d
Status: Downloaded newer image for alpine:latest
```

```console
$ docker image ls

REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
alpine                  latest              3fd9065eaf02        2 weeks ago         7.8MB
hello-world             latest              f2a91732366c        2 months ago        13.3kB
```

---

## Option D: Command block + collapsible output

Command in a `bash` block, output hidden in a collapsible `<details>` section:

```bash
docker image pull alpine
```

<details>
<summary>Expected output</summary>

```
Unable to find image 'alpine:latest' locally
latest: Pulling from library/alpine
88286f41530e: Pull complete
Digest: sha256:f006ecbb824d87947d0b51ab8488634bf69fe4094959d935c0c103f4820a417d
Status: Downloaded newer image for alpine:latest
```
</details>

```bash
docker image ls
```

<details>
<summary>Expected output</summary>

```
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
alpine                  latest              3fd9065eaf02        2 weeks ago         7.8MB
hello-world             latest              f2a91732366c        2 months ago        13.3kB
```
</details>

---

## Interactive Session Example

For comparison, here is how an **interactive container session** looks with each approach.

### Current / Option C (console)

```console
$ docker container run -it alpine /bin/sh
/ # ls
bin    dev    etc    home   lib    media  mnt    opt    proc   root   run    sbin   srv    sys    tmp    usr    var
/ # uname -a
Linux 97916e8cb5dc 4.4.27-moby #1 SMP Wed Oct 26 14:01:48 UTC 2016 x86_64 Linux
/ # exit
```

### Option A (separate blocks)

```bash
docker container run -it alpine /bin/sh
```

```
/ # ls
bin    dev    etc    home   lib    media  mnt    opt    proc   root   run    sbin   srv    sys    tmp    usr    var
/ # uname -a
Linux 97916e8cb5dc 4.4.27-moby #1 SMP Wed Oct 26 14:01:48 UTC 2016 x86_64 Linux
/ # exit
```

---

## Multi-line Command Example

### Current

```
$ docker container run \
--detach \
--name mydb \
--env MARIADB_ROOT_PASSWORD=my-secret-pw \
mariadb:latest

Unable to find image 'mariadb:latest' locally
latest: Pulling from library/mariadb
...
```

### Option A (separate blocks)

```bash
docker container run \
  --detach \
  --name mydb \
  --env MARIADB_ROOT_PASSWORD=my-secret-pw \
  mariadb:latest
```

```
Unable to find image 'mariadb:latest' locally
latest: Pulling from library/mariadb
...
```

### Option D (collapsible output)

```bash
docker container run \
  --detach \
  --name mydb \
  --env MARIADB_ROOT_PASSWORD=my-secret-pw \
  mariadb:latest
```

<details>
<summary>Expected output</summary>

```
Unable to find image 'mariadb:latest' locally
latest: Pulling from library/mariadb
...
```
</details>
