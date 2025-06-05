# Singleton

In simple terms: this design pattern ensures that only one object of a particular class is ever created and that single object is easily accessible from anywhere in the program.

## 🤔 How to think of it?

President of a country!

There’s only one official president at a time. If you ask for the president, everyone should refer to the same person.

## ⏱️ When do you need this pattern?

| Scenario | Example |
| -------- | ------- |
| 1\. Configuration or Settings Manager | You want the entire application to read and write to the same settings.<br>e.g. Dark mode setting, language preference, app config. |
| 2\. Logging \(Logger Class\) | You want all parts of the system to log messages to the same log file or console.<br>e.g. `Logger.log("Error occurred")` should write to the same log. |
| 3\. Database Connection | Maintaining multiple database connections can be expensive. One shared connection is better.<br>e.g. Accessing a PostgreSQL or MongoDB database. |
| 4\. Thread Pool or Resource Manager | You want one manager to coordinate shared system resources like memory or threads.<br>e.g. One `ThreadPoolManager` to manage background tasks. |
| 5\. Caching or State Store | You want a shared, in-memory data store (like a mini-database).<br>e.g. A singleton `Cache` object to avoid repeated computations or database calls. |
| 6\. File System or Print Spooler Access | Shared access to hardware or files can cause conflicts if not managed centrally.<br>e.g. A `PrinterManager` ensures only one job is sent to the printer at a time. |

**Rule of Thumb**: Use a Singleton when you need exactly one shared instance of a class to coordinate actions across your entire application.

## 💻 Code Example in Go:

```
package main

import (
	"fmt"
	"sync"
)

// Logger is the singleton type
type Logger struct{}

// Log is a method of the Logger
func (l *Logger) Log(message string) {
	fmt.Println("Log:", message)
}

// Singleton instance and sync.Once
var (
	loggerInstance *Logger
	once           sync.Once
)

// GetLogger returns the singleton instance
func GetLogger() *Logger {
	once.Do(func() {
		loggerInstance = &Logger{}
		fmt.Println("Logger initialized")
	})
	return loggerInstance
}

// Main usage
func main() {
	log1 := GetLogger()
	log2 := GetLogger()

	log1.Log("First message")
	log2.Log("Second message")

	// Check if they are the same instance
	fmt.Println("Same instance?", log1 == log2)
}
```

- - -

## 🏃 How to run the codebase?

* Ensure your current directory is `~/creational/singleton/code`
* Run `go run .`
