# go-whosonfirst-readwrite-github-tools

This is work in progress and, if you're reading this, it might work... until it doesn't.

## Tools

### update

Update one or more records in GitHub, using the GitHub API. Should this code be abstracted in to something like a `go-whosonfirst-update` package that works `go-whosonfirst-readwrite.Reader` and `Writer` instances? Maybe, but not today...

```
go run cmd/update/main.go -token {GITHUB_API_TOKEN} -action ceased -ceased-date '2019-06-29' -repo whosonfirst-data-venue-us-ca 571/513/137/571513137.geojson
```

Which results in this:

https://github.com/whosonfirst-data/whosonfirst-data-venue-us-ca/commit/5844892fabbc298dfa7ff773230d8f27463acdb1

There are a few things to notice here:

* `Flag data/588/390/025/588390025.geojson as ceased` is a fine commit message but it's not great.

* There is only a "commit" writer (the `-readwrite-github-` part) but there should also be a "pull request" writer.

* It does not format or "export" anything which is okay for flagging things as ceased but will need to be addressed as this tool learns to do more more update actions.

#### Actions

There is only one right now and it is `ceased`.

## See also

* https://github.com/whosonfirst/go-whosonfirst-readwrite
* https://github.com/whosonfirst/go-whosonfirst-readwrite-github