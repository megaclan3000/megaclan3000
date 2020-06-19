



# Set up development environment

Here is a minimal guide on how to set up the project. Some knowledge of the
command line and git is assumed.

## Install requirements

As a bare minimum, you will need the following software to be installed:

- [Go programming language](https://golang.org/doc/install#install)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- A decent editor of you choice, if possible with go support

## Basic setup and commands

With the tools above installed, clone the repo to a directory of your choice
(`$GOPATH` is recommended) and `cd` into it.

### Run the server

To start the application, use `go build` to build it and launch the compiled
binary with `./megaclan3000`

```
go build
./megaclan3000
```

You should see the application  running now in your browser if you go to
http://localhost:8080

## Run the tests

Before publishing changes to github, make sure the tests still pass. Run this
command from the project root, if everything went well you should see no errors

```
go test -v ./...
```

# Workflow, code review and permissions

Pushing directly to the `master` branch is forbidden to everyone. Every feature
or bugfix should be developed in a new branch. When ready, you can submit a
pull-request for your branch and it will be reviewed and merged into master by
the maintainer (@pinpox).

In summary, the basic workflow for a developer is as follows:

0. Make sure you have been added as contributor by the maintainer
1. Clone repo and cd into it.
2. Create a new branch, e.g. if you want to work on fixing a chart create a
   branch like `fix-weapons-chart`.

```
git checkout master
git checkout -b "fix-weapons-chart'
```

3. Work on your feature/fix
4. Add and commit your changes as needed, use meaningful commit messages

```
git add templates/chart-weapons.html
git commit -m'Fixed the weapons, now show all info'
```
5
