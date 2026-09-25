# Mailpit for testing mail setup in development

docker run --rm -it \
  --name mailpit \
  -p 127.0.0.1:1025:1025 \
  -p 127.0.0.1:8025:8025 \
  axllent/mailpit:latest \
  --smtp-auth-accept-any \
  --smtp-auth-allow-insecure