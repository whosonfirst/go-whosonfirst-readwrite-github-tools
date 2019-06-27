package main

import (
	"bytes"
	"context"
	"flag"
	"github.com/tidwall/sjson"
	"github.com/whosonfirst/go-whosonfirst-readwrite-github/reader"
	"github.com/whosonfirst/go-whosonfirst-readwrite-github/writer"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	var owner = flag.String("owner", "whosonfirst-data", "...")
	var repo = flag.String("repo", "whosonfirst-data", "...")
	var branch = flag.String("branch", "master", "...")
	var token = flag.String("token", "", "...")
	var action = flag.String("action", "", "...")

	var ceased_date = flag.String("ceased-date", "", "...")

	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := reader.NewGitHubAPIReader(ctx, *owner, *repo, *branch, *token)

	if err != nil {
		log.Fatal(err)
	}

	t := &writer.GitHubAPIWriterCommitTemplates{
		New:    "",
		Update: "Flag %s as ceased",
	}

	wr, err := writer.NewGitHubAPIWriter(ctx, *owner, *repo, *branch, *token, t)

	if err != nil {
		log.Fatal(err)
	}

	// exporter.NewExporter(wr)

	for _, path := range flag.Args() {

		old, err := r.Read(path)

		if err != nil {
			log.Fatal(err)
		}

		defer old.Close()

		body, err := ioutil.ReadAll(old)

		if err != nil {
			log.Fatal(err)
		}

		switch *action {
		case "ceased":

			dt := time.Now()

			if *ceased_date != "" {

				t, err := time.Parse("2006-01-02", *ceased_date)

				if err != nil {
					log.Fatal(err)
				}

				dt = t
			}

			body, err = sjson.SetBytes(body, "properties.edtf:cessation", dt.Format("2006-01-02"))

			if err != nil {
				log.Fatal(err)
			}

			body, err = sjson.SetBytes(body, "properties.mz:is_current", 0)

			if err != nil {
				log.Fatal(err)
			}

		default:
			log.Fatal("Unsupported action")
		}

		now := time.Now()
		body, err = sjson.SetBytes(body, "properties.wof:lastmodified", now.Unix())

		if err != nil {
			log.Fatal(err)
		}

		// exporter.ExportFeature(body)

		b := bytes.NewReader(body)
		out := ioutil.NopCloser(b)

		err = wr.Write(path, out)

		if err != nil {
			log.Fatal(err)
		}
	}

}
