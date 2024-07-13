**This project provides caching strategies. Currently iomplementing time based eviction policy but support for LRU as well.**

It also included **gitlab-ci.yml** file and docker file (docker-compose) for building and running the project.

To run the project and see output, follow the below steps:

```bash
docker-compose up
```

To run tests
```bash
go test -v ./...
```