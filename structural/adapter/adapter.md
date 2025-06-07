# Adapter

In simple terms: this design pattern is like a **translator** between two things that don’t naturally fit together. The **Adapter pattern** shines when you need to **bridge incompatible systems** — especially when you're dealing with legacy code, third-party libraries, or changes in your architecture.

## 🤔 How to think of it?

You have a laptop charger which supports only US Sockets.

You want to travel to Europe, but your charger will not be compatible with the European sockets.

You'll have to get an **adapter** for you to be able to charge!

## ⏱️ When do you need this pattern?

| Scenario | Example |
| -------- | ------- |
| 1\. Integrating Legacy Code | You have an old class or module that doesn’t match your new system’s interface.<br>e.g. You upgraded your app to use a new logging interface, but the old logger has a different method signature. An adapter lets the old logger work with the new system. |
| 2\. Using Third\-Party Libraries | You’re using a library that doesn’t match your code’s expectations.<br>e.g. You want to use a third-party payment gateway, but your app expects all payments to implement a specific `PaymentProcessor` interface. |
| 3\. Connecting Systems with Different Data Formats | Two components use different formats or APIs, but need to talk to each other.<br>e.g. One API returns dates as strings (`"2025-06-01"`), but your app expects `Date` objects. Use an adapter to translate formats. |
| 4\. Switching Implementations Without Changing Client Code | You want to change what's underneath (e.g., from local storage to cloud storage), but keep the **same interface**.<br>e.g. Switch from writing files locally to uploading to AWS S3, without changing the rest of your app. |
| 5\. Cross\-Platform or Device Support | You support multiple platforms (e.g., mobile, desktop) where the APIs differ.<br>e.g. Mobile and web use different image loaders. An adapter can present a consistent interface to your app. |
| 6\. Facade \+ Adapter Combo | You use Adapter inside a **Facade pattern** to wrap multiple incompatible subsystems and expose one clean API.<br>e.g. A `MediaPlayerFacade` uses adapters for MP3, MP4, and FLAC players that each have different method names. |

**Rule of Thumb**: Use **Adapter** when you need to **fit a square peg into a round hole** — without changing either.

## 💻 Code Example in Go:

```
package main

import "fmt"

// Step 1: The target interface (what the new system expects)
type Printer interface {
	Print(string)
}

// Step 2: The legacy class (incompatible)
type OldPrinter struct{}

func (op *OldPrinter) PrintText(text string) {
	fmt.Println("OldPrinter printing:", text)
}

// Step 3: The adapter — makes OldPrinter match the Printer interface
type OldPrinterAdapter struct {
	LegacyPrinter *OldPrinter
}

func (a *OldPrinterAdapter) Print(text string) {
	// Delegate the call to the legacy method
	a.LegacyPrinter.PrintText(text)
}

// Step 4: New system working with Printer interface
func usePrinter(p Printer, message string) {
	p.Print(message)
}

func main() {
	// Create legacy printer
	legacy := &OldPrinter{}

	// Wrap it in an adapter
	adapter := &OldPrinterAdapter{
		LegacyPrinter: legacy,
	}

	// Use it like it's a modern printer
	usePrinter(adapter, "Hello from the Adapter Pattern!")
}
```

- - -

## 🏃 How to run the codebase?

* Ensure your current directory is `~/structural/adapter/code`
* Run `go run .`
