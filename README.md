[![Circle CI](https://circleci.com/gh/echocat/caretakerd.svg?style=svg)](https://circleci.com/gh/echocat/caretakerd)
[![Go Report Card](https://goreportcard.com/badge/github.com/echocat/caretakerd)](https://goreportcard.com/report/github.com/echocat/caretakerd) [![Gitter](https://badges.gitter.im/echocat/caretakerd.svg)](https://gitter.im/echocat/caretakerd?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

# caretakerd

caretakerd is a minimal process supervisor build for easy use without depending on dependencies (no pun intended).

* [Documentation](#documentation)
* [Building](#building)
* [Contributing](#contributing)
* [Support](#support)
* [License](#license)

## Documentation

For general documentation please refer to the official homepage: [caretakerd.echocat.org](https://caretakerd.echocat.org).

For specific versions refer to: [caretakerd.echocat.org/all](https://caretakerd.echocat.org/all).

## Building

### Precondition

For building caretakerd you require only:

1. a compatible operating system (Linux, Windows or Mac OS X)
2. a working Java 8 installation.

There is no need for any working and installed Go installation (or anything else). The build system will download every dependency and build it automatically if necessary.

> **Hint:** The Go runtime build by the build system will be placed under ``~/.go-bootstrap``.

### Run

On Linux and Mac OS X:
```bash
# Build binaries only
./mvnw compile

# Run tests (includes compile)
./mvnw test

# Build resulting packages (including documentation - includes compile)
./mvnw package

# Set the target version number, increase the version number, do mvnw package,
# deploy everything to GitHub releases and set next development version number.
./mvnw release:prepare release:perform
```

On Windows:
```bash
# Build binaries only
mvnw compile

# Run tests (includes compile)
mvnw test

# Build resulting packages (including documentation - includes compile)
mvnw package

# Set the target version number, increase the version number, do mvnw package,
# deploy everything to GitHub releases and set next development version number.
mvnw release:prepare release:perform
```

### Build artifacts

* Compiled and lined binaries can be found under ``./target/gopath/bin/caretaker*``
* A generated documentation file can be found under ``./target/docs/caretakerd.html``
* Packaged TARZs and ZIPs can be found under ``./target/caretakerd-*.tar.gz`` and ``./target/caretakerd-*.zip``

## Contributing

caretakerd is an open source project of [echocat](https://echocat.org).
If you want to make this project even better, you can contribute to it on [Github](https://github.com/echocat/caretakerd)
by [fork us](https://github.com/echocat/caretakerd/fork).

If you commit code to this project you automatically accept that this code will be released under the [license](#license) of this project.

## Support

If you need support you should file a ticket in our [issue tracker](https://github.com/echocat/caretakerd/issues)
or join our chat at [echocat.slack.com/messages/caretakerd](https://echocat.slack.com/messages/caretakerd/).

## License

See [LICENSE](LICENSE) file.
