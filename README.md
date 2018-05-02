dp-developer-site
================

:warning: Work in progress :warning:

A Go application that generates the HTML pages for the developer site for ons.gov.uk. 

Should be available soon at [developer.ons.gov.uk](https://developer.ons.gov.uk).

### Getting started

To run this app you'll need [Golang](https://golang.org/) installed and setup. Then run:

```
make
```

or

```
go run main.go
```

#### Development *(Not available yet)* 

First run
```
make watch
```

The templates that make up the developer site are available in `/templates`.

Static pages, such as the introduction are in `/static`.

### Configuration

There's no configuration yet, although we probably want to add confguration for the version of Sixteens used.

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2017-2018, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
