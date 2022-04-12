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

docker

    docker pull fundipper/fungo

    
# command

fungo is easy to use, only have 5 commands.

## site

create a new site

    fungo site your-site-name

docker

    docker run -it --rm -v $PWD:/fungo fundipper/fungo site your-site-name

## file

create a new file

    fungo file your-file-model your-file-name

docker

    docker run -it --rm -v $PWD:/fungo -w /fungo/my-fungo-site fundipper/fungo file your-file-model your-file-name

ps:

default file model include [`article` , `page`, `document`]

you can define your own file model with yaml in `source`

## theme

create a new theme

    fungo theme your-theme-name

docker

    docker run -it --rm -v $PWD:/fungo -w /fungo/my-fungo-site fundipper/fungo theme your-theme-name

ps:

usually, you can get theme from fungo theme store or some open source repository

if you need your own one, use this command to create it

it's not too complicated, only needs tailwindcss

## serve

run serve mode

    fungo serve

docker

    docker run -it --rm -v $PWD:/fungo -w /fungo/my-fungo-site -p 3000:3000 fundipper/fungo serve

## build

run build mode

    fungo build

docker

    docker run -it --rm -v $PWD:/fungo -w /fungo/my-fungo-site fundipper/fungo build

# document

ours [official website](https://fungo.dev/), [document](https://fungo.dev/doc/overview/), [theme](https://fungo.dev/theme/), [blog](https://fungo.dev/post/) are all generate based on `fungo`.

how to use fungo ? see [https://fungo.dev/doc/overview/](https://fungo.dev/doc/overview/).

# feature

- [x] Support

  - [x] Article

  - [x] Page

  - [x] Document

  - [x] Customize

  - [x] I18N

- [x] Markdown

  - [x] TOC

  - [x] Highlighting

  - [x] Lazyload

- [x] Feeds

  - [x] RSS 2.0

  - [x] Atom 1.0

  - [x] JSON 1.1

- [x] I18N

  - [x] Content

  - [x] Template

- [x] SEO

  - [x] sitemap

  - [ ] robot.txt

# Thanks

- Language & Framework

  - [golang](https://go.dev/)

  - [tailwindcss](https://www.tailwindcss.com/)

  - [alpinejs](https://alpinejs.dev/)

- Tools & Libraries

  - [cobra](https://github.com/spf13/cobra)

  - [viper](https://github.com/spf13/viper)

  - [goldmark](https://github.com/yuin/goldmark)

  - [ristretto](https://github.com/dgraph-io/ristretto)

  - [httprouter](https://github.com/julienschmidt/httprouter)

  - [fsnotify](https://github.com/fsnotify/fsnotify)

  - [copy](https://github.com/otiai10/copy)

  - [go-git](https://github.com/go-git/go-git)

  - [etree](https://github.com/beevik/etree)
