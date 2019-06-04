# Smart Edge Challenge
This application will encrypt a message based off an RSA Private Key named `private.pem`. If this file is not found, it will automatically create a RSA key for you and encrypt your message.

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/hunterlong/smartedge) [![Build Status](https://travis-ci.com/hunterlong/smartedge.svg?branch=master)](https://travis-ci.com/hunterlong/smartedge)

## Details
- [x] Given a string input of up to 250 characters, return a JSON response compliant to the schema defined below.
  - JSON is compliant, message is limited to 250 characters
- [x] You are responsible for generating a public/private RSA or ECDSA keypair and persisting the keypair on the filesystem
  - App tries to open `private.pem`, if it's not found, it will generate a new RSA key and save to `private.pem`.
- [x] Subsequent invocations of your application should read from the same files
  - App will open private key from `private.pem`.
- [x] Document your code, at a minimum defining parameter types and return values for any public methods
  - Docs can be viewed on: Golang Docs
- [x] Include Unit Test(s) with instructions on how a Continuous Integration system can execute your test(s)
  - `main_test.go` includes tests for all functions and methods.
- [x] You may only use first order libraries, you may not use any third party libraries or packages.
  - App uses `crypto` and `encoding` packages
  
## Code Summary
- [SHA256 digest of input](https://github.com/hunterlong/smartedge/blob/master/main.go#L138)
- [Creating RSA Key](https://github.com/hunterlong/smartedge/blob/master/main.go#L62)
- [Openning RSA Key](https://github.com/hunterlong/smartedge/blob/master/main.go#L147)
- [Using string for RSA Key](https://github.com/hunterlong/smartedge/blob/master/main.go#L163)
- [Print as JSON](https://github.com/hunterlong/smartedge/blob/master/main.go#L92)
  
# Docker
This application has a multi-stage build for the Docker image. The first image will compile the code and create a binary. The second step copies that binary into a small Alpine Linux image.

To run this application in Docker, follow the commands below:

##### Build
```bash
docker build -t smartedge .
```

##### Run
```bash
docker run -it smartedge
```

##### Mounting Volume
```bash
docker run -it -v /app:/my_local_directory smartedge
```

##### Run with Custom Message
```bash
docker run -it smartedge smartedge "this will be encrypted!"
```

#### Additional Information
The container will start in the `/app` directory. You can mount this folder to your localhost to share private keys.

# Continuous Integration

Travis-CI will run the Continuous Integration Testing for this application.

[![Build Status](https://travis-ci.com/hunterlong/smartedge.svg?branch=master)](https://travis-ci.com/hunterlong/smartedge)
