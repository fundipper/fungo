package cli

import "sync"

const (
	_VERSION = "v1.0.0"
	_SERVER  = "server run on http://localhost%s"
	_BUILD   = "build file to `public` successed"
	_BANNER  = `
    ____                                
   / __/  __  __   ____    ____ _  ____ 
  / /_   / / / /  / __ \  / __  / / __ \
 / __/  / /_/ /  / / / / / /_/ / / /_/ /
/_/     \__,_/  /_/ /_/  \__, /  \____/ 
                        /____/ %s
https://fungo.dev
simple and fast
____________________________O/_______
                            O\
`
)

var once sync.Once
