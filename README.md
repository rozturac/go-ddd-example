## Folder Design
~~~~
├── api
│   └── configs
│   └── controllers
│       └── v1
├── application
│   └── consts
│   └── users
│       └── consumers
│       └── mappers
│       └── models
├── domain
│   └── common
│   └── users
├── infrastructure
│   └── common
│       └── event-dispatcher
│       └── persistance
│   └── users
~~~~

### Api (Presentation Layer)

<p>
This layer is a part developed for client-side (mobile, web, etc.) applications to access our domain. It will forward the requests from this layer to the application layer and expose the response it receives from the application layer.

As you can see in the example project, our project will be accessible from the outside world using HTTP protocol over the controllers classes.
Here is the sample code snippet that forwards the request received by HTTP to the Application Layer and return the result it receives.
</p>

```go
func CreateGuestUser(group *echo.Group, userService users.UserService) {
    path := fmt.Sprintf("%s/GuestUser", _prefix)
    group.POST(path, func(context echo.Context) error {
        var (
            user *models.NewUserModel
            err  error
        )
        
        if user, err = userService.AddNewGuestUser(context2.Background()); err != nil {
            return err
        }
        
        return context.JSON(http.StatusCreated, user)
    })
}
```

### Application Layer

<p>
Application layer currently has application services, event consumers, mappers and data transfer objects. However, it can also contain cross-cutting concerns such as transaction management, logging, caching and exception handling. (soon)

Application layer only call a aggregate root from domain layer and just may use its funcs. After It can be save all of changed that has been on aggregate-root to any database system. 

As you can see in the sample code block, a new guest user is created in the application service and then saved to the database. Then, the related user information created is mapped to the user-created-model and sent to the upper layer. 
</p>

```go
func (service userService) AddNewGuestUser(ctx context.Context) (*models.NewUserModel, error) {
    user := users.NewGuestUser()
    
    if err := service.Repository.Add(ctx, user); err != nil {
        return nil, err
    }
    
    return mappers.MapNewUserModel(user), nil
}
```

### Domain Layer
Will be completed soon..

### Infrastructure Layer
Will be completed soon..

