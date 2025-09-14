# Project History

This project was pretty much designed on the fly. The most basic concepts on wht it should be changed several times
and quite significantly (inexperienced solo developer conundrum). I try to make some sense of it here.

## Initial Concept Phase - frontend for headless CMS (early 2024)

- In the beginning the idea was to provide a simple frontend to the existing TYPO3 CMS website (classic server-rendered
  HTML)
  output.
- The idea was to use [a dedicated extension](https://docs.typo3.org/p/friendsoftypo3/headless/4.6/en-us/)
  to provide dual HTML/JSON responses.
- this idea was completely abandoned
    - the existing website fulfills its job to deliver static CMS content fast, reliable and SEO-friendly - no need to
      change anything

## Progressive Web App with feature parity to iOS (mid to late 2024)

- at this time the existing iOS app was considered to be rather feature-complete
- Android was in its infancy (early alpha version)
- next target: achieve feature-parity with native apps for Android users and non-smartphone users
- main features at the time: load and display team-related data from BSM
- remains a core concept

## Central team hub with dedicated backend (late 2024 - present)

- request from team administration: provide tool to replace kadermanager.de and spielerplus.de as game/practice admin
  tools
- desire to bring team data in-house instead of relying on external providers
- connect with existing BSM services to provide smart functionality (automatic import and sync of games for each team)
- this required a dedicated backend - [Pocketbase](https://pocketbase.io/) was chosen, used as a Go framework for custom
  extensions

## Test drive in 2025 Verbandsliga season

- MVP finished and rolled out in March 2025
- used throughout the season by the Skylarks second team
- bug fixes and improvements carried out incrementally based on user feedback
- planned: full club rollout for 2026 season

