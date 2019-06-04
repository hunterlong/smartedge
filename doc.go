// Package main is an application to create a code challenge for Smart Edge. This app will encrypt a message from a RSA private key and output a JSON message with
// the complete message, the signature of the message, and a copy of the public key that encrypted the message.
//
// Usage
//
// To use this application, simply input 1 argument as your message!
//		smartedge "hello world"
// The application will print a JSON object that will return the message and it's signature based on the private key (private.pem).
//
// Testing
//
//		go test -v
//
// Docker
//
// You can also build a small Alpine Linux image by running:
//		docker build -t smartedge .
// Then you can run it with:
//		docker run -it smartedge
package main
