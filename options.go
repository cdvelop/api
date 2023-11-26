package api

import (
	"os"

	"github.com/cdvelop/strings"
)

func (c *config) processOptions(options ...string) {
	for _, option := range options {

		switch {

		case strings.Contains(option, "cache:") != 0:
			var cache_option string
			err := strings.ExtractTwoPointArgument(option, &cache_option)
			if err == "" {

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
			c.production_mode = false
		}

		if arg == "no-cache" {
			c.static_cache = "no-cache"
		}
	}

}
