go-blog
-------

**go-blog** is a web application that gives user ability to write, delete and modify posts. It was written for learning purposes.

Requirements
------------

Go-blog was developed under Fedora Workstation and it is only one supported operating system to run it. However I am pretty sure, that, you can easly run this applicatio under other Linux, BSD-like or Windows operating system, if it only provide support for technologies listed at the end of `README` file.

Installation
------------

Before installing application, you should set `$SESSION_KEY` envriomental variable.

Follow below instructions to install go-blog completely in containers.

    git clone https://github.com/thinkofher/go-blog.git && cd go-blog
    podman-compose up -d

You can also use `docker-compose` and it'll propably work, but it's not supported.

And that's all! Now go to `http://0.0.0.0:8080` adress in your browser and start using app.

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
- [podman-compose](https://github.com/containers/podman-compose)
- [Open Iconic](https://useiconic.com/open/)
