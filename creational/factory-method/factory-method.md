# Factory Method

In simple terms: this design pattern delegates the creation of objects to subclasses, so that you don't need to know the exact class or logic when creating an object.

## 🤔 How to think of it?

You go to a coffee shop!

You just tell the barista, "I want my coffee!"

The barista decides whether to prepare an Espresso, Latte, or Cappuccino based on some rules (your taste, time of day, etc.).

You don't need to know how it's made - you just order a coffee.

## ⏱️ When do you need this pattern?

| Scenario | Example |
| -------- | ------- |
| 1\. When the exact type of object isn't known until runtime | Logging system that creates different loggers based on configuration: `FileLogger`, `ConsoleLogger`, `RemoteLogger`.<br>`logger := GetLogger(config.Type)`<br>`logger.Log("Something happened!")` |
| 2\. To follow the Open/Closed Principle | When you want to **add new types** without changing existing code. Just add a new class that implements the interface and update the factory.<br>e.g. Adding a new payment mothod:<br>`type PaymentMethod interface {	Process(amount float64)}` |
| 3\. To encapsulate complex creation logic | When object creation involves logic, decisions, or dependencies, it's better to **isolate that in a factory** rather than spreading it in your codebase.<br>e.g. Creating a database connection with optional SSL, retry policies, etc.<br>`conn := NewDBConnectionFactory(config).Create()` |
| 4\. To return objects from a family of related classes | e.g. In a game, you may have different `Enemy` types: `Zombie`, `Alien`, `Robot`. A factory decides which to spawn. |
| 5\. When you want to avoid tight coupling to specific implementations | Using a factory, you depend on **interfaces** rather than concrete types.<br>So, instead of: <br>`user := &MySQLUserRepository{}`<br>Use: <br>`repo := GetUserRepository("mysql")` |
| 6\. When the creation logic may change in the future | Putting creation logic in a factory allows it to **evolve independently** from where it's used. |

**Rule of Thumb**: Use the Factory Method Pattern when you need to create objects without knowing exactly what type they are ahead of time, or when the creation logic varies depending on conditions

## 💻 Code Example in Go:

```
package main

import ("fmt")

// Notification is the common interface
type Notification interface {
	Send(message string)
}

// EmailNotification is one implementation
type EmailNotification struct{}

func (e EmailNotification) Send(message string) {
	fmt.Println("Sending EMAIL with message:", message)
}

// SMSNotification is another implementation
type SMSNotification struct{}

func (s SMSNotification) Send(message string) {
	fmt.Println("Sending SMS with message:", message)
}

// NotificationFactory is the factory method
func NotificationFactory(notificationType string) Notification {
	switch notificationType {
	case "email":
		return EmailNotification{}
	case "sms":
		return SMSNotification{}
	default:
		return nil
	}
}

// Main function using the factory
func main() {
	// Use factory to get a Notification
	notification := NotificationFactory("sms")
	if notification != nil {
		notification.Send("Hello from Factory Method!")
	} else {
		fmt.Println("Invalid notification type")
	}
}
```

- - -

## 🏃 How to run the codebase?

* Ensure your current directory is `~/creational/factory-method/code`
* Run `go run .`
