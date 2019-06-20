# go-whosonfirst-readwrite-github-tools

This is work in progress and, if you're reading this, it might work... until it doesn't.

## Tools

### update

Update one or more records in GitHub, using the GitHub API. Should this code be abstracted in to something like a `go-whosonfirst-update` package that works `go-whosonfirst-readwrite.Reader` and `Writer` instances? Maybe, but not today...

```
./bin/update -action ceased -repo whosonfirst-data-venue-us-ca -token {GITHUB_API_TOKEN} 588/390/025/588390025.geojson
```

Which results in this:

https://github.com/whosonfirst-data/whosonfirst-data-venue-us-ca/commit/7755aaeeb7cfb8d3c0d1dabaaa2148eb8205ae91

There are a few things to notice here:

* There's no way to specify the date something "ceased" to be. [Lucca's](https://spelunker.whosonfirst.org/id/588390025/) actually closed at the end of April, 2019 not June 20, 2019.

* `Flag data/588/390/025/588390025.geojson as ceased` is a fine commit message but it's not great.

* There is only a "commit" writer (the `-readwrite-github-` part) but there should also be a "pull request" writer.

* It does not format or "export" anything which is okay for flagging things as ceased but will need to be addressed as this tool learns to do more more update actions.

#### Actions

There is only one right now and it is `ceased`.

## See also

* https://github.com/whosonfirst/go-whosonfirst-readwrite
* https://github.com/whosonfirst/go-whosonfirst-readwrite-github