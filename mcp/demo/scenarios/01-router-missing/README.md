# Scenario 01 — "my router isn't working"

The most common Traefik question. Two failure modes, two services.

## Run

```bash
./demo.sh up 01-router-missing
```

## What's broken

| Service | Symptom | Root cause |
|---------|---------|------------|
| `api`   | router exists but in `warning`, traffic 404s | references middleware `does-not-exist@docker` that was never defined |
| `typo`  | router is completely absent from the API | the `rule` label is misspelled `rulee`, so Traefik ignores it |

## Demo in Claude Desktop

- "Why is my `api` router not working?"
- "I added a router for `typo.localhost` but it doesn't exist — what's wrong?"

Expected: the assistant cross-references `list_routers` (the `api` error is right
there; `typo` is missing entirely), the reload status, and the Traefik logs to
rank the cause — undefined middleware vs a label that never registered.

## Reset

```bash
./demo.sh down 01-router-missing
```
