# go-whosonfirst-readwrite-github-tools

This is work in progress and, if you're reading this, does not work yet.

## Tools

### update

Update one or more records in GitHub, using the GitHub API. Should this code be abstracted in to something like a `go-whosonfirst-update` package that works `go-whosonfirst-readwrite.Reader` and `Writer` instances? Maybe, but not today...

```
./bin/update -repo whosonfirst-data-venue-us-ca -action ceased -token {GITHUB_API_TOKEN} 588/392/711/588392711.geojson
```

#### Actions

There is only one right now and it is `ceased`.

