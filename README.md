# Skylarks Diamond Planner

Progressive Web App for the [Berlin Skylarks Baseball & Softball Club](https://www.tib-baseball.de/).
Built with Pocketbase and SvelteKit with the static adapter.

## Concept / Background

Main rationale to create this project: At project start, different types of data were being processed in different web
services:

* The main administrative platform [Baseball & Softball Manager / BSM](https://bsm.baseball-softball.de/) of the German
  Baseball &
  Softball Federation (DBV).
    * Everything related to organised play is processed there (clubs, games, leagues, teams, player, stats)
    * Accessed via REST API
    * _External_: no club access to any internal logic
* [The current Skylarks website](https://www.tib-baseball.de/), served as a [TYPO3 CMS](https://typo3.org/) website (
  PHP-based).
    * Processes mostly typical CMS _content_ (articles, info pages), but also additional team data that is distinct from
      BSM data: club teams, player profiles (with more data than what is available in BSM),
      training times, game reports
    * Mainly displayed in the TYPO3 frontend (server-side templating) directly, but is also partly accessible via REST
      API
    * _Internal_: custom-built, full club access

See also: [Project History](docs/History.md)

## Project goal

* collect data from all relevant sources and displays it in a user-friendly way
* Progressive Web App with mobile-first design
* provide logic to carry out team organisation tasks
  (practice and game attendance, teams and members, stats)
* → replace existing external tools for those tasks
* become _THE_ central hub for team activities

### Design decisions

* This tool will process administrative data for club and team events (events, announcements, comments, attendance data
  and
  statistics)
* this data will be distinct from both CMS and BSM data:
    * BSM data relates to the club's organised play ⇒ main use case is _external_
    * CMS data is strictly presentational, the CMS does not care about administrative logic ⇒ main use case is
      _external_
    * ⇒ here, the main use case is _internal_

## Project requirements

* basic familiarity with the concepts of SvelteKit and Svelte – there is an
  excellent [official tutorial](https://learn.svelte.dev/tutorial/welcome-to-svelte) available
* Go `v1.25` or higher
* Node.js `v24` or higher
* access to Berlin Skylarks environment secrets and API keys

## Local Development

1. Clone the repository

```bash
git clone git@github.com:tib-baseball-softball/skylarks-next.git
git submodule update --init
cd skylarks-next
```

2. Set up environment

```bash
cp backend/.env.dist backend/.env
cp .env.dist .env
```

* BSM API key (from BSM user account with at the club admin scope for Berlin Skylarks)
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

- Go Backend: http://127.0.0.1:8090/
- Vite Frontend: http://localhost:5173/

### Production

- unified URL: https://app.berlinskylarks.de/

## Deployment

- Fully automated via GitHub Actions, deploys on every push to branch `main` or `stage`, excluding documentation files
  and folders.
- app is deployed via Docker-Compose in a Portainer stack on the team VPS
- Traefik router (not deployed via this repository) acts as a reverse proxy and forwards requests
- ⇒ Traefik labels need to be set in Compose files

### Server Setup / Manual Deployment

What the GitHub Action does:

1. Build the app according to the Dockerfile
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
