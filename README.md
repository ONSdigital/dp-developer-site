# dp-developer-site

A Go application that generates the HTML pages for the developer hub for ons.gov.uk.

Available at [developer.ons.gov.uk](https://developer.ons.gov.uk).

## Getting started

To run this app you'll need certain versions installed:

| Language | version   |
|----------|-----------|
| Go       | >= 1.21.7 |
| Node     | = 20.1.3  |
| Python   | >= 3.9.2  |

These versions are what the pipeline will build with.

You can use `nvm use` to set your Node version to the current compatible one.

You'll also need `github.com/fogleman/serve@latest` which you can install using:

```sh
    make install-prereqs
```

## Development

The templates that make up the developer hub are available in `/templates`.

Static pages, such as the introduction are in `/static`.

The most common use case will be to update files and watch them change on the local site:

to watch the files for changes and host them run:

```bash
make watch-serve
```

or, with a custom port:

```bash
make watch-serve PORT=8000
```

### Using Local Paths

You can also load API specs from your local machine. 
Just provide the file path instead of a URL, and the app will figure out if it's local or remote.

Example:

```go
sources := spec.APIs{
    {"dataset-api", "https://raw.githubusercontent.com/ONSdigital/dp-dataset-api/master/swagger.yaml", nil, nil},
    {"local-api", "/path/to/your/local/swagger.yaml", nil, nil},
}
```

### Other notable commands are as follows

To watch and automatically rebuild the site on any changes to any `.go`, `.md` or `.tmpl` file run:

```bash
make watch
```

To only serve the built assets (on default port of `23600`) run:

```bash
make serve
```

...which has an optional `PORT` environment variable, for example:

```bash
make server PORT=8000
```

### Configuration

| Environment variable | Default | Description                    |
| -------------------- | ------- | ------------------------------ |
| PORT                 | 23600   | The port to serve the files on |

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2017-2024, Office for National Statistics (<https://www.ons.gov.uk>)

Released under MIT license, see [LICENSE](LICENSE.md) for details. 
