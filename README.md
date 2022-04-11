         ____                                
        / __/  __  __   ____    ____ _  ____ 
       / /_   / / / /  / __ \  / __  / / __ \
      / __/  / /_/ /  / / / / / /_/ / / /_/ /
     /_/     \__,_/  /_/ /_/  \__, /  \____/ 
                            /____/ v1.0.0   

# Overview

[fungo](https://fungo.dev/ "fungo") is a simple and fast open-source static site generators base on golang.

# install

go get

    go get github.com/fundipper/fungo

# tree

    .
    ├── config.toml
    ├── content
    │   ├── document 
    │   ├── page 
    │   ├── media
    │   ├── post 
    ├── source
    └── theme
        └── fungo
            ├── assets
            ├── i18n
            ├── package-lock.json
            ├── package.json
            ├── tailwind.config.js
            ├── templates
            ├── theme.toml
            └── watch.sh

# command

fungo is easy to use, only have 5 commands.

## site

create a new site 

    fungo site your-site-name

## theme

create a new theme (if you need your own template) 

    fungo theme your-theme-name

## file

create a new file

    fungo file your-file-model your-file-name

## serve
run serve mode

    fungo serve

## build

run build mode

    fungo build

# Thanks

- [x] cmd 

https://github.com/spf13/cobra

- [x] config

https://github.com/spf13/viper

- [x] router

https://github.com/julienschmidt/httprouter

- [x] markdown

https://github.com/yuin/goldmark

- [x] cache

https://github.com/dgraph-io/ristretto

- [x] render

go/template

- [x] copy

https://github.com/otiai10/copy

- [x] git

https://github.com/go-git/go-git

- [x] watch file

https://github.com/fsnotify/fsnotify

- [x] sitemap

https://github.com/beevik/etree

- [x] feeds

