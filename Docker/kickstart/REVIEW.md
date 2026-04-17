---
nav_exclude: true
---

# Docker Kickstart Training – Remaining Review Notes

This document tracks **currently known open issues** in the Docker Kickstart training material.
Earlier pre-fix review notes that described the state of the repository prior to recent clean-up
have been removed to avoid confusion, since many of those issues have now been resolved or the
referenced files have been deleted.

The items below are intended as a lightweight to-do list for future improvements, not as a
historical record.

---

## 1. Docker Compose v1 usage

Some chapters still use the legacy `docker-compose` v1 CLI syntax instead of the modern
Docker Compose v2 plugin (`docker compose`):

- `mongo.md`
- `mongo-2.md`

**Suggested actions:**

- Update command examples to use `docker compose` rather than `docker-compose`.
- Adjust any installation/setup instructions to match the recommended way of installing
  Docker Compose v2 for the target environment.
- Verify that all screenshots, code listings, and narrative text use consistent, up-to-date
  syntax.

---

## 2. Voting App Swarm chapter

The `votingapp-swarm.md` chapter was reverted to its original state because the improvements
need more thorough review. The following items should be addressed when the chapter is revisited:

- Update EOL `postgres:9.4` to a current Postgres image
- Replace deprecated Docker for AWS/Azure beta references
- Add cross-platform IP address lookup instructions (macOS, Linux, Windows)
- Replace archived `docker/labs` networking link with an up-to-date reference
- Update "Docker Store" references to "Docker Hub"
- Review the inline `docker-stack.yml` for accuracy against the current upstream
  example-voting-app repo
- Verify that `dockersamples/examplevotingapp_vote:before` and `:after` image tags still
  exist on Docker Hub
- General copy-editing pass for clarity and modern Docker CLI usage

---

## 3. Content that depends on external repositories

These items rely on external repos that may change independently:

- `webapps-part2.md` clones `github.com/dockersamples/linux_tweet_app` — if this repo is
  modified or removed, the tutorial breaks
- `votingapp-compose.md` and `votingapp-swarm.md` clone
  `github.com/docker/example-voting-app` — the compose file shown inline may drift from the
  actual repo
- `votingapp-swarm.md` uses `dockersamples/examplevotingapp_vote:before` and `:after` image
  tags which may be removed from Docker Hub

---

## 4. Enhancements (not bugs)

- No `.dockerignore` coverage in the Dockerfile chapters
- No `USER` directive / non-root container discussion
- No container resource limits (`--memory`, `--cpus`) coverage
- The `mongo.md`, `mongo-2.md`, and `nodered.md` chapters exist but are not listed in the
  main `readme.md` index (they appear to be supplementary/optional content)
