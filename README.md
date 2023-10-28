# What to commit

## Description

<https://github.com/ngerakines/commitment> analog built on Golang using
fasthttp. Made for personal self-education purposes and self-hosting. Latest
version from default branch pushed to Docker Hub via GitHub Actions.

## Provision

To provision the application, run the following command:

```bash
make provision
```

App will be available on <http://localhost:9321>.

## Example

* <http://localhost:9321/all> - all messages
* <http://localhost:9321/> - random message

## k6

Little performance test using [k6](https://k6.io/):

```bash
k6 run k6.js
```
