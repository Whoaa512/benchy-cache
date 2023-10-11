package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/valyala/fasthttp"
)

const dataDir = "./data"

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Method()) {
		case "GET":
			handleGET(ctx)
		case "PUT":
			handlePUT(ctx)
		default:
			ctx.Error("Unsupported method", fasthttp.StatusMethodNotAllowed)
		}
	}

	log.Fatal(fasthttp.ListenAndServe(":6942", requestHandler)) }

func handleGET(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/status":
		handleStatus(ctx)
	default:
		if strings.HasPrefix(string(ctx.Path()), "/cache/") {
			filePath := filepath.Join(dataDir, filepath.Clean(string(ctx.Path()[7:])))
			fasthttp.ServeFile(ctx, filePath)
		} else {
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
}

func handlePUT(ctx *fasthttp.RequestCtx) {
	if strings.HasPrefix(string(ctx.Path()), "/cache/") {
		filePath := filepath.Join(dataDir, filepath.Clean(string(ctx.Path()[7:])))
		err := os.WriteFile(filePath, ctx.Request.Body(), 0644)
		if err != nil {
			print(err.Error())
			ctx.Error("Error saving file", fasthttp.StatusInternalServerError)
		}
	} else {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func handleStatus(ctx *fasthttp.RequestCtx) {
	var totalSize int64
	var itemCount int

	err := filepath.WalkDir(dataDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			totalSize += info.Size()
			itemCount++
		}
		return nil
	})

	if err != nil {
		ctx.Error("Error reading directory", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString(fmt.Sprintf(`{"size_on_disk": %d, "number_of_items": %d}`, totalSize, itemCount))
}
