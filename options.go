package api

import (
	"os"
	"strings"

	"github.com/cdvelop/gotools"
	. "github.com/cdvelop/output"
)

func (c *config) processOptions(options ...string) {
	for _, option := range options {

		switch {

		case strings.Contains(option, "cache:"):
			var cache_option string
			err := gotools.ExtractTwoPointArgument(option, &cache_option)
			if err == nil {

				var seconds string

				switch cache_option {

				case "week":
					seconds = "604800" // 7 días

				case "month":
					seconds = "2592000" // 1 mes

				case "year":
					seconds = "31536000" // 1 año
				}

				c.static_cache = "public, max-age=" + seconds
			}

		}
	}

	for _, arg := range os.Args {
		if arg == "dev" {
			c.developer_mode = true
			c.static_cache = "no-cache"
			PrintWarning("*** Api en Modo Desarrollo ***\n")
		}
	}

	if !c.developer_mode {
		PrintOK("*** Api en Modo Producción ***\n")
	}

}