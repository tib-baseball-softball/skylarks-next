# Backend-specific documentation

## Pocketbase Schema

- the file `pb_schema.json` is an export of the currently configured collections access rules.
- it acts as a source of truth for the backend schema
- in addition, pocketbase migrations are saved in the corresponding folder
- for type-safe access to custom models, RecordProxy structs are used
- TS types are generated on the frontend side