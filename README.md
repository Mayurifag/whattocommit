# What to commit

## Description

<https://github.com/ngerakines/commitment> analog built on Golang using
fasthttp. Made for personal self-education purposes and self-hosting. Latest
version from default branch pushed to Docker Hub via GitHub Actions.

## Provision

Make sure you have Docker installed and have `9321` port not occupied.

### Development

To provision the application, run the following command:

```bash
make provision
```

### Fast launch latest main branch version

To launch the latest version from main branch, run the following command:

```bash
docker run --rm -p 9321:8080 mayurifag/whattocommit:main
```

## Examples

App will be available on <http://localhost:9321>.

* <http://localhost:9321/all> - all messages
* <http://localhost:9321/> - random message

## k6

Little performance test using [k6](https://k6.io/):

```bash
k6 run k6.js
```
