# Backend-specific documentation

## Pocketbase Schema

- the file `pb_schema.json` is an export of the currently configured collections access rules.
- it acts as a source of truth for the backend schema
- unfortunately, it is not involved in the build process and must be imported into the production instance manually
  after each change.
- TS types are generated on the frontend side