# Sigmund

<img id="gopher" src="https://storage.googleapis.com/gopherizeme.appspot.com/gophers/7da7cd5ba32fae25e03301f30ba3a1296b47ca2e.png" alt="Sigmund Go" height=200px>

A Go lang tool designed to `shrink` AWS Autoscaling Clusters based on `CPU` and `Memory` Cloudwatch Alarm Metrics.

This tool is meant to be for people who want to scale their instances back in whenever **both** `CPU` **AND** `Memory` conditions are met.

## Prerequisites

-   A DynamoDB table
-   AWS credentials with the following permissions:
    -   DynamoDB Table Read/Write access
    -   Ability to Execute Autoscaling Policies

### DyanamoDB Table structure

In order to work, your DynamoDB table should look like the one that follows:

| ID  | isLowCPU | isLowMemory |
| :-: | :------: | :---------: |
|  0  |  false   |    false    |

**NB** it is important that the table is pre-populated with the values showed above, most importantly you must make sure that the `ID` = 0 - that value will never change.

## Installing

You can directly use the `go` tool to download and install the `sigmund` package into your `GOPATH`:

```bash
$ go get github.com/darkraiden/sigmund.go
```

You can also clone the repository yourself:

```bash
$ mkdir -p $GOPATH/src/github.com/darkraiden/
$ cd $GOPATH/src/github.com/darkraiden/
$ git clone git@gihub.com:darkraiden/sigmund
```

Note: This project uses dep so you'll need to run `dep ensure` to grab all the dependencies.

## How to use it

First things first, initialise a `Sigmund` from your application:

```go
  // Get a Sigmund
  s, err := sigmund.New("eu-west-1", "anASGName", "anASGPolicyName", "aDynamoTableName", "LowCPU") // LowCPU, LowMemory, OKCPU and OKMemory are accepted values
  if err != nil {
    panic(err)
  }
```

Now that you have your `Sigmund`, you're ready to update the DB _and eventually_ execute the autoscaling group policy:

```go
  err := s.Shrink()
  if err != nil {
    panic(err)
  }
```

## Running the tests

Every Package of this project comes with some unit tests which use the Go `testing` package. Run the tests, from the package folder, by typing:

```bash
$ go test -v ./...
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/darkraiden/sigmund/tags).

## To Do

-   [ ] Add meaningful tests 🙈

## Authors

-   [Davide Di Mauro](https://github.com/darkraiden)

See also the list of [contributors](contributors.md) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
