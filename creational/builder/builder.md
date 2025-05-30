# Builder

In simple terms: this design pattern helps you build complex objects step by step.

## 🤔 How to think of it?

You want make your own burger!

You want to choose the bun, the patty, the sauce, and the toppings. The pattern allows you to customize, without bothering yourself to know how the burger is put together internally.

## ⏱️ When do you need this pattern?

| Scenario | Example |
| -------- | ------- |
| 1\. Creating Objects with Many Optional Fields | When a struct has lots of fields, and not all of them are always used. |
| 2\. Avoiding Complex Constructors | If you have a constructor similar to this: ``func NewUser(name string, email string, phone string, address string, age int, isActive bool) *User``. Builder makes it more readable and less error-prone. |
| 3\. Immutable Objects | If you want to build an object once and not change it later, the builder pattern helps construct the object safely before it becomes immutable. |
| 4\. Creating Variants of the Same Object | You can easily create different versions of the same base object, like:<br>`admin := NewUserBuilder().SetRole("admin").SetPermissions(all).Build()`, and `guest := NewUserBuilder().SetRole("guest").SetPermissions(readOnly).Build()` |
| 5\. Readability and Fluent APIs | Builders create fluent APIs, where you can chain method calls for clarity, like `car := NewCarBuilder().SetMake("Honda").SetModel("Civic").SetColor("Black").Build` |
| 6\. Testing and Mocks | Builders are helpful for setting up test data with minimal boilerplate, like `testUser := NewUserBuilder().SetName("Test User").SetEmail("test@example.com").Build()` |
| 7\. Avoiding Struct Tag Confusion in JSON/DB Models | When populating complex structs from different sources (API, DB), a builder can give you more control over field mapping and default values.\` |

**Rule of Thumb**: Use the Builder Pattern when object construction feels repetitive, unclear, or overloaded with options.

## 💻 Code Example in Go:

```
package main

import (
"fmt"
)

// User is the object we want to build
type User struct {
Name     string
Email    string
Age      int
Phone    string
Address  string
IsActive bool
}

// UserBuilder helps build User objects
type UserBuilder struct {
user User
}

// NewUserBuilder returns a new instance of UserBuilder
func NewUserBuilder() *UserBuilder {
return &UserBuilder{}
}

// SetName sets the name field
func (b *UserBuilder) SetName(name string) *UserBuilder {
b.user.Name = name
return b
}

// SetEmail sets the email field
func (b *UserBuilder) SetEmail(email string) *UserBuilder {
b.user.Email = email
return b
}

// SetAge sets the age field
func (b *UserBuilder) SetAge(age int) *UserBuilder {
b.user.Age = age
return b
}

// SetPhone sets the phone number
func (b *UserBuilder) SetPhone(phone string) *UserBuilder {
b.user.Phone = phone
return b
}

// SetAddress sets the address
func (b *UserBuilder) SetAddress(address string) *UserBuilder {
b.user.Address = address
return b
}

// SetIsActive sets whether the user is active
func (b *UserBuilder) SetIsActive(isActive bool) *UserBuilder {
b.user.IsActive = isActive
return b
}

// Build returns the fully constructed User
func (b *UserBuilder) Build() User {
return b.user
}

func main() {
// Use the builder to create a user
user := NewUserBuilder().
SetName("Alice").
SetEmail("alice@example.com").
SetAge(30).
SetPhone("123-456-7890").
SetIsActive(true).
Build()

fmt.Printf("User: %+v\n", user)
}
```

- - -

## 🏃 How to run the codebase?

* Ensure your current directory is `~/creational/builder/code`
* Run `go run .`
