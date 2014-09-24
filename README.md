# git-http-multi-backend

Git "smart HTTP" protocol server.

[git-http-backend](http://git-scm.com/docs/git-http-backend), a server side
implementation of Git over HTTP is a simple CGI program to serve git
repositories. CGI itself doesn't handle multiple simultaneous requests and
configuring Nginx + `git-http-backend` to serve multiple clients requires
setting up a pool of FastCGI processes, which requires multiple tools like
`fcgiwrap`, `spawn-fcgi` and `multiwatch` (to make sure `fcgiwrap` workers
respawn).

`git-http-multi-backend` solves this problem in a simple way by starting a
regular HTTP server which for each incoming request spawns a new
`git-http-backend` process and forwards its response to the client.

Due to the fact that Go's `net/http` server handles every request in each own
goroutine it's possible to handle large number of concurrent clones/pushes.

Limiting the number of simultaneous connections is not in the scope of this
tool and can be easily achieved by configuring Nginx or Haproxy to do this
task.

## Installation

Currently you need Go development environment to build git-http-multi-backend.

The following command will fetch the package and build the binary at
`$GOPATH/bin/git-http-multi-backend`:

    go get gitorious.org/gitorious/git-http-multi-backend

## Usage

### Starting

Usage:

    git-http-multi-backend [options]

Options:

* `-r <repos-dir>` - Directory containing git repositories, defaults to "."
* `-c <backend-command>` - CGI binary to execute, defaults to "git http-backend"
* `-l <[addr]:port>` -  Address/port to listen on, defaults to 127.0.0.1:80

Example:

    git-http-multi-backend -r /var/git/repositories

## License

git-http-multi-backend is free software licensed under the
[GNU Affero General Public License](http://www.gnu.org/licenses/agpl-3.0.html).
git-http-multi-backend is developed as part of the Gitorious project.
