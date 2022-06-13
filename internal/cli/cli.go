package cli

import "sync"

const (
	_VERSION = "v0.3.7"
	_SERVER  = "server run on http://localhost%s"
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
