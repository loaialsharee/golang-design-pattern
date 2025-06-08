# Adapter

In simple terms: this design pattern is a structural design pattern that **decouples an abstraction from its implementation**, so they can vary **independently**.

## 🤔 How to think of it?

You havea remote control (as the abstraction) and a TV (as the implementation).

You can have a smart remote or a basic remote.
Also, you can have Samsung TV, LG, etc.

By keeping the **remote logic** separate from the **TV logic**, you can mix and match without rewriting everything.

## ⏱️ When do you need this pattern?

| Scenario | Example |
| -------- | ------- |
| 1\. UI Rendering Across Platforms | You want the same UI components to work on multiple operating systems (Windows, macOS, Linux), but each OS renders differently.<br><br>e.g. <br>**Abstraction**: `Window`, `Button`, `Dialog`<br>**Implementation**: `WindowsRenderer`, `MacRenderer`, `LinuxRenderer` |
| 2\. Device Control \(Remote / Hardware APIs\) | Different remotes (basic, smart) controlling different devices (TVs, Radios, etc.)<br>e.g.<br>**Abstraction**: `RemoteControl`<br>**Implementation**: `TV`, `Radio` |
| 3\. Document Exporters or Converters | You support multiple document types (PDF, Word, HTML) and output formats (local file, cloud storage, email).<br>e.g.<br>**Abstraction**: `DocumentExporter`<br>**Implementation**: `FileSystemOutput`, `S3Uploader`, `EmailSender`<br><br>Now you can create combinations like:<br>* PDF to File* Word to S3* HTML to Email |
| 4\. Multichannel Notification System | Send the same notification through different channels (Email, SMS, Push)<br>e.g. <br>* **Abstraction**: `Notification` (Alert, Reminder)* **Implementation**: `EmailSender`, `SMSSender`, `PushSender`Send any type of alert through any delivery method, independently. |
| 5\. Payment Gateways | You support different payment types (CreditCard, PayPal, Crypto) and different providers (Stripe, Braintree, Coinbase)<br>e.g.<br>* **Abstraction**: `Payment`* **Implementation**: `StripeProcessor`, `CoinbaseProcessor`You can swap either side without rewriting both. |

**Rule of Thumb**: Use **Adapter** when you need to **fit a square peg into a round hole** — without changing either.

## 💻 Code Example in Go:

```
package main

import "fmt"

//////////////////////////
// Implementation Layer //
//////////////////////////

type Device interface {
	On()
	Off()
	SetVolume(percent int)
}

type TV struct{}

func (tv *TV) On() {
	fmt.Println("TV is now ON")
}

func (tv *TV) Off() {
	fmt.Println("TV is now OFF")
}

func (tv *TV) SetVolume(percent int) {
	fmt.Printf("TV volume set to %d%%\n", percent)
}

type Radio struct{}

func (r *Radio) On() {
	fmt.Println("Radio is now ON")
}

func (r *Radio) Off() {
	fmt.Println("Radio is now OFF")
}

func (r *Radio) SetVolume(percent int) {
	fmt.Printf("Radio volume set to %d%%\n", percent)
}

////////////////////////
// Abstraction Layer  //
////////////////////////

type RemoteControl interface {
	TogglePower()
	VolumeUp()
	VolumeDown()
}

type BasicRemote struct {
	device Device
	volume int
}

func (r *BasicRemote) TogglePower() {
	fmt.Println("BasicRemote toggling power...")
	r.device.On()
}

func (r *BasicRemote) VolumeUp() {
	r.volume += 10
	r.device.SetVolume(r.volume)
}

func (r *BasicRemote) VolumeDown() {
	r.volume -= 10
	r.device.SetVolume(r.volume)
}

////////////////////////
// Usage              //
////////////////////////

func main() {
	// You can plug any device into the remote
	tv := &TV{}
	radio := &Radio{}

	fmt.Println("Controlling TV:")
	tvRemote := &BasicRemote{device: tv}
	tvRemote.TogglePower()
	tvRemote.VolumeUp()

	fmt.Println("\nControlling Radio:")
	radioRemote := &BasicRemote{device: radio}
	radioRemote.TogglePower()
	radioRemote.VolumeUp()
}
```

- - -

## 🏃 How to run the codebase?

* Ensure your current directory is `~/structural/bridge/code`
* Run `go run .`
