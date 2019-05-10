# Go Template

A simple go service that uses keyckloak gatekeeper to keep it safe.  Includes docker setup for keycloak as well to make local dev easy.

Work In Progress, just learning go.

# setup

The first time you run your docker container, you will need to create a realm, you can do so manually,
or you can do so with this command.



# go setup

You need several environemtn variables for this to run:

```
export TZ="America/Los_Angeles"
export PROJECTDIR="/path to project"
export PGSQL_ROOT_PASSWORD="password"
export KC_ADMIN_PASSWORD="password"

```

# Keycloak setup

Keycloak is setup and connected via the docker container, but it still needs user data
for you to develop with.  You can launch it (once it is running) with this link:

[http://localhost:9081/auth/](http://localhost:9081/auth/)

## Importing a realm
To create an admin account and import a previously exported realm run:

```
docker run -e KEYCLOAK_USER=<USERNAME> -e KEYCLOAK_PASSWORD=<PASSWORD> \
    -e KEYCLOAK_IMPORT=/tmp/example-realm.json -v /tmp/example-realm.json:/tmp/example-realm.json jboss/keycloak
```

