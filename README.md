# smsapi
A server that exposes the SMS functionality of the SIMCom SIM800 cellular modem to the Internet, allowing you to write applications that connect to it.

Some of this code could be rewritten to be better and more reliable (maybe using functions on structs...), but it does _work_.

## Setup
The first time you run it, a file called `config.toml` will be created in the current directory, with contents like this:

```toml
# smsapi configuration
[server]
Port = 4725
Token = "hello"
SerialPort = "/dev/ttyUSB0"
```

You should stop the server (control-C) and change these settings.
* `Port` determines the port the web server should listen on.
* `Token` is a secret token that is required for all requests. This is how requests are authenticated.
* `SerialPort` is the serial port the SIM800 is connected to. On Windows, this will be in the format `COM1` (with the exact number depending on your system). On Linux and Mac, this will probably be in the format `/dev/ttyUSB1` or `/dev/ttyACM1`.

## Usage
There are currently five available API routes. All of them except for `GET /` return information as JSON (with the `Content-Type` header set appropriately).

All API routes require the token defined in the `config.toml` file to be sent with your request (as either a GET parameter in the URL, or, in the case of POST requests, as a form value).

* `GET /` - returns the string "hello there\n\ni am a server\n\ni can send text messages" (with `\n` being actual newlines)
* `GET /modem/battery` - returns information about the battery level of the modem, in the format `{ "status": "ok", "percentage": 100, "voltage": 4123 }`. (voltage is in millivolts)
* `GET /modem/ccid` - returns the CCID code for the current SIM card, in the format `{ "status": "ok", "ccid": "<ccid goes here>" }`
* `GET /modem/carrier` - returns the current cell carrier the SIM800 is connected to, in the format `{ "status": "ok", "carrier": "T-Mobile USA" }`
* `POST /sms/send` - requires parameters `to` and `msg`. Sends an SMS message to the given phone number. Returns either `{ "status": "ok" }` on success or `{ "status": "error", "error": "<error string>" }` on failure.
