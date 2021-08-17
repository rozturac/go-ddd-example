# go-ddd-example
/Golang - Domain Driven Design

## Design

* Api
    * Implements restful service using echo framework
    * Implementation config file that has environment
* Domain
    * Define aggregate
        * Entity, Value Object, Enumerations
    * Define interface
        * Repository interface for infrastructure
* Application
    * Write business logic
* Infrastructure
    * Implements repository interface
    * Implements dependency injections
