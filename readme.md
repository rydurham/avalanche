# Avalanche

A solution to the Falling Rocks code kata, written in Go.

## Usage

Submit "arenas" to the program via json files:

```
$ ./avalanche examples/example1.json
```

There are some pre-defined json files in the `examples` folder.

## Build

Once you have cloned this project into the `github.com/rydurham/avalanche` folder in your gopath, you can build it from within that directory:

```
{gopath}/github.com/rydurham/avalanche $ go build
```

A dockerfile has also been provided that will do this for you:

```
docker build -t avalanche .
docker run avalanche ./avalanche examples/example1.json
```

## Tests

To run the tests:

```
{gopath}/github.com/rydurham/avalanche $ go test
```

or

```
docker run avalanche go test
```

You can optionally pass in a `-v` flag to see more details.

## To Do

The next steps for this project are:

- Update the code to use more idiomatic error handling
- Add an http server that will turn this service into a json api
