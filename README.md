# gallery

A web album server with Go, SvelteKit, and lightGallery.

## Development

```console
$ cd ui
$ npm run dev
```

```console
$ go run -tags dev cmd/gallery/main.go cmd/gallery/dev.go -d ~/Desktop/
```

## Production

```console
$ go install github.com/ichiban/gallery/cmd/gallery@latest
```