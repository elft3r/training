# Docker Kickstart Training - Review

## Overall Assessment

The Docker Kickstart training is a well-structured, progressive curriculum that takes learners from basic container operations through to orchestration with Docker Swarm. The hands-on, task-based approach is effective.

---

## Resolved Issues

The following issues have been identified and fixed:

- ✅ Updated Docker Desktop installation links in `setup.md`
- ✅ Fixed dead documentation links in `alpine.md`
- ✅ Fixed `$MARIADB_ROOT_PASSWORD` env var expansion bug in `alpine.md` (was expanding on host shell where the variable doesn't exist)
- ✅ Updated `VIRTUAL SIZE` → `SIZE` column header in `alpine.md`
- ✅ Fixed container list output in `alpine.md` to match actual tutorial commands
- ✅ Added macOS/Windows note for `cat /etc/issue` in `alpine.md`
- ✅ Fixed typo "not install" → "not installed" in `alpine.md`
- ✅ Fixed `docker container ps` → `docker container ls` throughout
- ✅ Replaced archived docker/labs repo link in `webapps-part1.md`
- ✅ Fixed broken grammar in `webapps-part1.md` ("The command summary the above command", "will publish instruct")
- ✅ Replaced `0.0.0.0` browser URLs with `localhost` throughout
- ✅ Standardized Docker CLI to modern syntax (`docker container run`, `docker image ls`, etc.) across all chapters
- ✅ Fixed port 80/8080 text-vs-command contradiction in `webapps-part2.md`
- ✅ Fixed port 8080/8081 text-vs-command contradiction in `webapps-part2.md`
- ✅ Removed incorrect `PUSH` Dockerfile instruction reference in `webapps-part2.md`
- ✅ Fixed typo "linux_tweet app" → "linux_tweet_app" in `webapps-part2.md`
- ✅ Added PowerShell instructions to Windows notes in `webapps-part2.md` and `devops.md`
- ✅ Updated Dockerfile best practices link in `webapps-part2.md`
- ✅ Fixed step numbering jump (3→13) in `devops.md`
- ✅ Replaced hardcoded `vegasbrianc` username with placeholder in `devops.md`
- ✅ Fixed `printf` vs `echo` for reliable escape sequences in `multistage-builds.md`
- ✅ Fixed hardcoded `~/Training` path in `multistage-builds.md`
- ✅ Fixed port 5002→5000 in `docker compose ps` output in `votingapp-compose.md`
- ✅ Fixed stray `"` characters in `ls -l` output in `secrets.md`
- ✅ Removed empty "Docker & Kubernetes" section in `secrets.md`
- ✅ Added link to `nextsteps.md` from `secrets.md` (was a dead-end)
- ✅ Updated Docker version prerequisite in `secrets.md`
- ✅ Enhanced `nextsteps.md` with useful links to other training content
- ✅ Fixed missing `$` prompt prefix in `images-and-volumes.md`
- ✅ Updated `docker images` → `docker image ls` in `images-and-volumes.md`
- ✅ Updated `docker inspect` → `docker volume inspect` / `docker container inspect` in `images-and-volumes.md`
- ✅ Fixed `docker run` → `docker container run` throughout `images-and-volumes.md`
- ✅ Multi-stage build coverage added (`multistage-builds.md`)
- ✅ GitHub Actions versions updated in `devops.md`
- ✅ Flask app Dockerfile updated to Python 3
- ✅ Navigation links between chapters fixed
- ✅ "Docker Store" references updated to "Docker Hub"

---

## Remaining Known Issues

### Content that depends on external repositories

These items rely on external repos that may change independently:

- `webapps-part2.md` clones `github.com/dockersamples/linux_tweet_app` — if this repo is modified or removed, the tutorial breaks
- `votingapp-compose.md` and `votingapp-swarm.md` clone `github.com/docker/example-voting-app` — the compose file shown inline may drift from the actual repo
- `votingapp-swarm.md` uses `dockersamples/examplevotingapp_vote:before` and `:after` image tags which may be removed from Docker Hub

### Voting App Swarm Chapter — Needs Full Resort

The `votingapp-swarm.md` chapter was reverted to its original state because the improvements need more thorough review. The following items should be addressed when the chapter is revisited:

- Update EOL `postgres:9.4` → a current Postgres image
- Replace deprecated Docker for AWS/Azure beta references
- Add cross-platform IP address lookup instructions (macOS, Linux, Windows)
- Replace archived `docker/labs` networking link with an up-to-date reference
- Update "Docker Store" references to "Docker Hub"
- Review the inline `docker-stack.yml` for accuracy against the current upstream example-voting-app repo
- Verify that `dockersamples/examplevotingapp_vote:before` and `:after` image tags still exist on Docker Hub
- General copy-editing pass for clarity and modern Docker CLI usage

### Enhancements (not bugs)

- No `.dockerignore` coverage in the Dockerfile chapters
- No `USER` directive / non-root container discussion
- No container resource limits (`--memory`, `--cpus`) coverage
- The `mongo.md`, `mongo-2.md`, and `nodered.md` chapters exist but are not listed in the main `readme.md` index (they appear to be supplementary/optional content)
