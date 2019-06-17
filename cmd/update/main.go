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
	
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := reader.NewGitHubAPIReader(ctx, *owner, *repo, *branch, *token)

	if err != nil {
		log.Fatal(err)
	}

	t := &writer.GitHubAPIWriterCommitTemplates{
		New: "",
		Update: "Flag %s as ceased",
	}
		
	wr, err := writer.NewGitHubAPIWriter(ctx, *owner, *repo, *branch, *token, t)

	if err != nil {
		log.Fatal(err)
	}

	// exporter.NewExporter(wr)
	
	dt := time.Now() // sudo make this an option...

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
		
			body, err = sjson.SetBytes(body, "properties.edtf:cessation", dt.Format("20060102"))

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
