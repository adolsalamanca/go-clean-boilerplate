# Go Clean Architecture Boilerplate

## Description
This is a project aims to serve developers to organize Go projects while trying to use both DDD & Clean Architecture.
I really encourage you to read the articles from references, but I leave here a summary of my thoughts:

* Clean architecture attempts to create better projects, simplifying  complexity but also make them more scalable, easily extendable, and reduce maintenance cost.
* The code should be decoupled from UI frameworks, databases, queues, data streams... etc.
* Thanks to layer differentiation, the code is more robust and testable in isolation.
* Inner layers, where your business rules reside, should not know anything about outer ones.

## Layers
You can find the different layers below, from outside to inside.
* Infrastructure Layer (a.k.a.s Frameworks & Drivers) -> ```code is under $PROJECT/internal/infrastructure```
* Interface layer (a.k.a.s Interface adapters) -> ```code is under $PROJECT/internal/interface```
* Application layer (a.k.a.s Use cases) -> ```code is under $PROJECT/internal/application```
* Entities layer (a.k.a.s Entities) -> ```code is under $PROJECT/internal/domain```
  

#### Architecture Diagram (by Uncle Bob)

![Clean Architecture](https://github.com/adolsalamanca/go-clean-boilerplate/blob/main/cleanArchitecture.jpeg)

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

## License
This project is licensed under MIT License.


