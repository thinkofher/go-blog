go-blog
-------

**go-blog** is a web application that gives user ability to write, delete and modify posts. It was written for learning purposes.

Requirements
------------

Go-blog was developed under Fedora Workstation and it is only one supported operating system to run it. However I am pretty sure, that, you can easly run this applicatio under other Linux, BSD-like or Windows operating system, if it only provide support for technologies listed at the end of `README` file.

Installation
------------

Before installing application, you should set `$SESSION_KEY` envriomental variable.

Follow below instructions to install go-blog.

    git clone https://github.com/thinkofher/go-blog.git && cd go-blog
    go get -u .
    ./init_podman.sh
    ./download_fonts.sh
    podman start goblog-db
    go build .
    ./go-blog

You can also install [fresh](https://github.com/gravityblast/fresh) and use it instead of building application every time. It's really helpful when developing new features for web apps (make sure to add `$GOPATH/bin` to your `$PATH`).

Check out `config.go` file and modify it for your needs.

Used technologies
-----------------

- [gorilla/mux](https://github.com/gorilla/mux)
- [gorilla/sessions](https://github.com/gorilla/sessions/)
- [golang/x/crypto](https://golang.org/pkg/crypto/)
- [pq](https://github.com/lib/pq)
- [PostgreSQL](https://www.postgresql.org/)
- [Bootstrap](https://getbootstrap.com/)
- [podman](https://podman.io/)
- [Open Iconic](https://useiconic.com/open/)
