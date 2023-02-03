package gallery

import "embed"

//go:generate sh -c "cd ui; npm run build"
//go:embed all:ui/build/*
var UIBuild embed.FS
