package standard

import (
	// standard Caddy modules
	_ "github.com/rspier/caddy/v2/caddyconfig/caddyfile"
	_ "github.com/rspier/caddy/v2/modules/caddyhttp/standard"
	_ "github.com/rspier/caddy/v2/modules/caddypki"
	_ "github.com/rspier/caddy/v2/modules/caddypki/acmeserver"
	_ "github.com/rspier/caddy/v2/modules/caddytls"
	_ "github.com/rspier/caddy/v2/modules/caddytls/distributedstek"
	_ "github.com/rspier/caddy/v2/modules/caddytls/standardstek"
	_ "github.com/rspier/caddy/v2/modules/filestorage"
	_ "github.com/rspier/caddy/v2/modules/logging"
)
