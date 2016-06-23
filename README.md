## httpserver

Convert any directory into a fast http(s) website.

![httpserver screenshot](https://raw.githubusercontent.com/kataras/httpserver/master/screenshot.png)

## Install

`$ go get -u github.com/kataras/httpserver`


Make sure that you have appended your system's environment variable `$PATH` the `$GOPATH/bin`

## Usage

All flags are **OPTIONAL**

```
httpserver [global arguments...]

GLOBAL ARGUMENTS:

   -d        the current working directory (default './')

   -g        enable serving with gzip compression & use file cache (default 'false')

   -h        the server addr to listen for (default '0.0.0.0:8080')

   -l        enable requests logging (default 'false')

   -c        cert file for https (default '')

   -k        key file for https (default '')

```

Example:

`$ httpserver -d ./ -g -h 0.0.0.0:8080 -l -c C:\\secret\\mykeyfile.cert -k C:\\secret\\mykeyfile.key`

## Versioning

Current: **v0.0.2**


Read more about Semantic Versioning 2.0.0

 - http://semver.org/
 - https://en.wikipedia.org/wiki/Software_versioning
 - https://wiki.debian.org/UpstreamGuide#Releases_and_Versions


## Contributors

Thanks goes to the people who have contributed code to this package, see the

- [Contributors page](https://github.com/kataras/httpserver/graphs/contributors).


## License

This project is licensed under the [BSD 3-Clause License](https://opensource.org/licenses/BSD-3-Clause).
License can be found [here](https://github.com/kataras/httpserver/blob/master/LICENSE).
