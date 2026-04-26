package hazm

import "embed"

//go:embed data/*.dat
var embeddedData embed.FS
