# Skylarks-Next

Progressive Web App for the [Berlin Skylarks Baseball & Softball Club](https://www.tib-baseball.de/).
Built with Pocketbase, SvelteKit, Skeleton UI and Tailwind CSS.

## Concept / Background

As of now, different types of data are being processed in different backends:

* The main backend [Baseball & Softball Manager / BSM](https://bsm.baseball-softball.de/) of the German Baseball &
  Softball Federation (DBV).
    * Everything related to organised play is processed there (clubs, games, leagues, teams, player, stats)
    * Accessed via REST API
    * _External_: no club access to any internal logic
* [The current Skylarks website](https://www.tib-baseball.de/), served as a [TYPO3 CMS](https://typo3.org/) website (
  PHP-based).
    * Processes mostly typical CMS content (articles, info pages), but also additional team data that is distinct from
      BSM data: club teams, player profiles (with more data than what is available in BSM),
      training times, game reports
    * Mainly displayed in the TYPO3 frontend directly, but is also partly accessible via REST API
    * _Internal_: custom-built, full club access
* The backend in this repository
    * Go-based (package `github.com/tib-baseball-softball/skylarks-next`)
    * processes administrative data for club and team events (events, announcements, comments, attendance data and
      statistics)
    * distinct from both CMS and BSM data
    * _Internal_: custom-built, full club access

See also: [Project History](docs/History.md)

## Project goal

* collect data from all relevant sources and displays it in a user-friendly way
* Progressive Web App with mobile-first design
* provide backend and frontend logic to carry out team organisation tasks
  (practice and game attendance, teams and members, stats)
* -> replace existing external tools for those tasks
* become _THE_ central hub for team activities

## Project requirements

* basic familiarity with the concepts of SvelteKit and Svelte - there is an
  excellent [official tutorial](https://learn.svelte.dev/tutorial/welcome-to-svelte) available
* basic familiarity with Go and TypeScript
* Node.js `v24` or higher
* access to Berlin Skylarks environment secrets and API keys

## Local Development

1. Clone the repository

```bash
git clone git@github.com:tib-baseball-softball/skylarks-fe-next.git
cd skylarks-fe-next
```

2. Set up environment

```bash
cp backend/.env.dist backend/.env
cp .env.dist .env
```

* BSM API key (from BSM user account with at the club admin scope for Berlin Skylarks)
* API Auth Header (from `.env` on TYPO3 host server)
* `PUBLIC_TYPO3_URL` can either be set to a TYPO3 dev environment running locally or the production URL

3. Get project dependencies

backend

```bash
cd backend
go mod download`
go run . serve
```

frontend

```bash
pnpm install`
# yarn/deno/bun could also work, untested
```

4. Start the dev server (frontend watches for file changes automatically, backend needs to be restarted manually)
   backend

```bash
go run . serve
```

frontend

```bash
pnpm dev

```

## URLs

### Local/Development

- BE: http://127.0.0.1:8090/
- FE: http://localhost:5173/

### Production

- BE: https://pb.berlinskylarks.de/
- FE: https://app.berlinskylarks.de/

## Deployment

- Fully automated via GitHub Actions, deploys on every push to branch `main`, excluding documentation files and folders.
- app is deployed via Docker-Compose in a Portainer stack
- Traefik router (not deployed via this repository) acts a reverse proxy and forwards requests
- => Traefik labels need to be set in compose files

### Server Setup / Manual Deployment

What the GitHub Action does:

1. Build the app according to the Dockerfiles
2. Pushes docker images to [Docker Hub](https://hub.docker.com/repositories/obnoxieux)
3. Triggers [Portainer](https://docs.portainer.io/) deployment via Webhook

## Bugs and Problems

* Please open an issue in this repository.

## Terms of use

Licensed under [AGPL-3](LICENSE).

---

> Explore other templates from [The Good Docs Project](https://thegooddocsproject.dev/). Use
> our [feedback form](https://thegooddocsproject.dev/feedback/?template=Readme) to give feedback on this template.
---
