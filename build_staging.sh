docker login
docker build \
    --build-arg BUILD_MODE="build:staging" \
    --platform=linux/amd64 \
    --tag obnoxieux/skylarks-diamond-planner:staging .
docker push obnoxieux/skylarks-diamond-planner:staging
