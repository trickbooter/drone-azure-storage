# drone-azure-storage

Drone plugin to publish files and artifacts to Azure Storage. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

[![Build Status](http://138.68.151.154/api/badges/Retail_Development/drone-azure-storage/status.svg)](http://138.68.151.154/Retail_Development/drone-azure-storage)

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-azure-storage <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com",
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "account_key": "123889asd89u8hsfdjh98128hh",
        "storage_account": "my-storage-account",
        "container": "my-container",
        "source": "folder/to/upload"
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-azure-storage <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com",
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "account_key": "123889asd89u8hsfdjh98128hh",
        "storage_account": "my-storage-account",
        "container": "my-container",
        "source": "folder/to/upload"
    }
}
EOF
```
