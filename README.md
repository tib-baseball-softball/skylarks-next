# Skylarks-FE-Next
New experimental frontend for the [website of the Berlin Skylarks Baseball & Softball Club](https://www.tib-baseball.de/).
Built with SvelteKit, Skeleton UI and Tailwind CSS.

## Concept / Background

As of now, different types of data are being processed in different backends:
* The main backend [Baseball & Softball Manager / BSM](https://bsm.baseball-softball.de/) of the German Baseball & Softball Federation (DBV).
  * Everything related to organised play is processed there (clubs, games, leagues, teams, player, stats)
  * Accessed via REST API (provided by BSM - _external_)
* The current Skylarks website, served as a [TYPO3 CMS](https://typo3.org/) website (PHP-based).
  * Processes additional data that is distinct from BSM data: club teams, player profiles (with more data than what is available in BSM), 
  training times, game reports
  * Mainly displayed in the TYPO3 frontend directly, but can also be accessed via REST API (custom - _internal_)

## Project goal
This project is intended as a pure frontend that collects data from these sources and displays it in a user-friendly way. 
Since usage is expected to be predominantly mobile, focus is on providing a Progressive Web App with mobile-first design.

## Project requirements
* basic familiarity with the concepts of SvelteKit and Svelte - there is an excellent [official tutorial](https://learn.svelte.dev/tutorial/welcome-to-svelte) available
* Node.js `v20` or higher
* access to Berlin Skylarks environment secrets and API keys

## Local Development

1. Clone the repository

```bash
git clone git@github.com:tib-baseball-softball/skylarks-fe-next.git
cd skylarks-fe-next
```

2. Set up environment

```bash
cp .env.dist .env
```
* BSM API key (from BSM user account with at the club admin scope for Berlin Skylarks)
* API Auth Header (from `.env` on TYPO3 host server)
* `PUBLIC_BACKEND_URL` can either be set to a TYPO3 dev environment running locally or the production URL

3. Get project dependencies

```bash
npm install`
# yarn/pnpm/bun could also work, untested
```

4. Start the dev server (watches for file changes automatically)

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

Default local port is http://localhost:5173/

## Deployment

Fully automated via GitHub Actions, deploys on every push to branch `main`, excluding documentation files and folders.

### Server Setup / Manual Deployment

What the GitHub Action does:
1. `npm ci && npm run build` for production build output
2. sets static environment variables
3. copies build output to server via `rsync`
4. App is served as Node.js application listening on a specific port
5. Process is managed by [PM2](https://pm2.keymetrics.io/)
6. The [Caddy Webserver](https://caddyserver.com/) is used as reverse proxy.


## Bugs and Problems
* Please open an issue in this repository.

## Terms of use
Licensed under [AGPL-3](LICENSE).

---

> Explore other templates from [The Good Docs Project](https://thegooddocsproject.dev/). Use our [feedback form](https://thegooddocsproject.dev/feedback/?template=Readme) to give feedback on this template.
---
