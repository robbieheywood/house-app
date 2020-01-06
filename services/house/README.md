# House Server

This is the main server.
It fields requests from users.
At the moment, it does an incredibly primitive auth check and then replies with 'hello-world'.

## Commands

Build image: `DOCKER_BUILDKIT=1; docker build $REPO -f $REPO/services/house/Dockerfile -t eu.gcr.io/tensile-imprint-156310/house:$TAG`

Run container locally: `docker run -p 8080:8080 eu.gcr.io/tensile-imprint-156310/house:$TAG`

Push image: `docker push eu.gcr.io/tensile-imprint-156310/house:$TAG`