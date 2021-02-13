[![ContainerSSH - Launch Containers on Demand](https://containerssh.github.io/images/logo-for-embedding.svg)](https://containerssh.github.io/)

<!--suppress HtmlDeprecatedAttribute -->
<h1 align="center">ContainerSSH Struct Manipulation Library</h1>

[![Go Report Card](https://goreportcard.com/badge/github.com/containerssh/structutils?style=for-the-badge)](https://goreportcard.com/report/github.com/containerssh/library-template)
[![LGTM Alerts](https://img.shields.io/lgtm/alerts/github/ContainerSSH/structutils?style=for-the-badge)](https://lgtm.com/projects/g/ContainerSSH/library-template/)

This library provides methods for manipulating structs.

<p align="center"><strong>⚠⚠⚠ Warning: This is a developer documentation. ⚠⚠⚠</strong><br />The user documentation for ContainerSSH is located at <a href="https://containerssh.io">containerssh.io</a>.</p>

## Copying structures

The `structutils.Copy()` method provides a deep-copy mechanism for structs:

```go
oldData := yourStruct{}
newData := yourStruct{}
err := structutils.Copy(&newData, &oldData)
```

The `newData` will be completely independent from `oldData`, including the copying of any pointers.

## Merging structures

The `structutils.Merge()` method provides a deep merge of two structs:

```go
oldData := yourStruct{}
newData := yourStruct{}
err := structutils.Merge(newData, oldData)
```

The `Merge` method will copy non-default values from `oldData` to `newData`.

## Adding default values

The `structutils.Defaults()` method loads the default values from the `default` tag in a struct:

```go
type yourStruct struct {
	Text string `default:"Hello world!"`
	Decision bool `default:"true"`
	Number int `default:"42"`
}

//...

func main() {
    data := yourStruct{}
    structutils.Defaults(&data)
    // data will now contain the default values
}
```

Default values can be provided as follows:

- Scalars can be provided directly.
- Maps, structs, etc. can be provided in JSON format.
- `time.Duration` can be provided in text format (e.g. 60s).
- Structs may implement a receiver with the method called SetDefaults() as described in [defaults.go](defaults.go).
