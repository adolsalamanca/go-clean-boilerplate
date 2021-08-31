# Go Clean Architecture Boilerplate

## Description
This is a project aims to serve developers as a boilerplate for new Go projects. It tries to use DDD & Clean Architecture.
I really encourage you to read the articles from references, but I leave here a summary of my thoughts:

* Clean architecture attempts to create better projects, simplifying  complexity but also make them more scalable, easily extendable, and reduce maintenance cost.
* The code should be decoupled from UI frameworks, databases, queues, data streams... etc.
* Thanks to layer differentiation, the code is more robust and testable in isolation.
* Inner layers, where your business rules reside, should not know anything about outer ones.

## Layers
You can find the different layers below, from outside to inside.
* Infrastructure Layer (a.k.a.s Frameworks & Drivers) : stores external logic, could be related to DB, Queues, other systems. 
    * code: ```$PROJECT/internal/infrastructure```

* Interface layer (a.k.a.s Interface adapters) : this layer serves as a connector between infrastructure and our application.
    * code: ``` $PROJECT/internal/interface```

* Application layer (a.k.a.s Use cases) : as its name in the diagram indicates, it contains our application business use cases, in this specific app it stores the web server.
    * code: ```$PROJECT/internal/application```

* Domain layer (a.k.a.s Entities) : this is where our core resides, our specific system entities and sacred logic will be, this layer will have no clue about the rest of the system and shouldn't be something frequently changed.
    * code: ```$PROJECT/internal/domain```

<br>

Apart from those pretty well differentiated layers, there's a few folders:
* Assets: used to store assets from the repo(icons, images...). Here we only have Uncle's Bob Clean Architecture diagram.
* Cmd/App: this is a Go convention to store our main application.
* Config: serve to store our configuration files, here we only have stuff related to metrics, statsd and prometheus specifically.
* Pkg: another Go convention, here we save code that could be potentially used by external applications.
* Test: additional stuff, in this project we have a parallel docker-compose to run tests and also a sql script for DB initialization(this one could be inside a migrations folder to be run by Goose or another migrations tool instead)


#### Architecture Diagram (by Uncle Bob)

![Clean Architecture](https://github.com/adolsalamanca/go-clean-boilerplate/blob/main/assets/cleanArchitecture.jpeg)

## Additional notes
As you can see, the project only have some acceptance tests, and not tests for each layer.
Note that this is a side effect of having a simple CRUD application, as the project does not have much logic I preferred not to add them.
In case the project evolves to have more business logic, I would start adding more tests to each layer, so we can test the behavior of the system in isolation. 


## Running the code

##### Tests

```bash
$ make test
```


#### Project

```bash
$ make up
```

## Author
* **Adolfo Rodriguez** - *go-clean-boilerplate* - [adolsalamanca](https://github.com/adolsalamanca)


## Articles
* [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
* [SOLID Go Design](https://dave.cheney.net/2016/08/20/solid-go-design)

## References
* [Non-official standard Layout](https://github.com/golang-standards/project-layout)

## License
This project is licensed under MIT License.


