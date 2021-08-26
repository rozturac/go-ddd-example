# go-ddd-example
/Golang - Domain Driven Design

## Design

* Api
    * Implements restful service using echo framework
    * Implementation config file that has environment
* Domain
    * Define Aggregate
        * Entity, Value Object, Enumerations, DomainEvent
    * Define Interface
        * Repository interface for infrastructure
* Application
    * Write business logic
    * Implements event consumers
* Infrastructure
    * Implements repository interface
    * Implements dependency injections
    * Implements event dispatcher
    * Implements RMQC framework to work with rabbitmq
