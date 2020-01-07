# gostarter
## Boilerplate For Golang Application/Service Project

### Setting Up
- Replace All Occurrences of `anjuspace/gostarter` with your username repository name
- Replace All Occurrences of `app` with your desired image name


### Adding New Libraries/Dependencies
```bash
go mod vendor
```
## Go Modules

This assumes the use of go modules (which will be the default for all Go builds
as of Go 1.13) and vendoring (which reasonable minds might disagree about).
You will need to run `go mod vendor` to create a `vendor` directory when you
have dependencies.

## Building

Run `make` or `make build` to compile your app.  This will use a Docker image
to build your app, with the current directory volume-mounted into place.  This
will store incremental state for the fastest possible build.  Run `make
all-build` to build for all architectures.

Run `make container` to build the container image.  It will calculate the image
tag based on the most recent git tag, and whether the repo is "dirty" since
that tag (see `make version`).  Run `make all-container` to build containers
for all supported architectures.

Run `make push` to push the container image to `REGISTRY`.  Run `make all-push`
to push the container images for all architectures.

Run `make clean` to clean up.

### Using GitHub Registry

Create and Push:

```bash
docker login docker.pkg.github.com -u <USERNAME> -p <GITHUB_TOKEN>
docker build -t  docker.pkg.github.com/anjuspace/gostarter/app:latest .
# make container
docker push docker.pkg.github.com/anjuspace/gostarter/app:latest
# make push
```

Pull and Run:

```bash
docker pull docker.pkg.github.com/anjuspace/gostarter/app:latest
docker run docker.pkg.github.com/anjuspace/gostarter/app:latest
```
