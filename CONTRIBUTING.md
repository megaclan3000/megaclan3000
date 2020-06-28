# Set up development environment

Here is a minimal guide on how to set up the project. Some knowledge of the
command line and git is assumed.

## Get a steam API key

Go to https://steamcommunity.com/dev/apikey and generate a API Key. You will
need it to fetch any user data from the steam, otherwise the application will
not be able to run.

## Create configuration file

Copy `./config.json.example` to `./config.conf` and edit accordingly. You need
at least one steam id and the API key set for the application to work.

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

5. Push your changes to your branch. `git push` will give you an error message,
   the first time, saying there is no upstream branch (yet). Use the printed
   command to add one. This is only needed on the first push to a new branch.

```
# git push

fatal: The current branch CONTRIBUTING.md has no upstream branch.
To push the current branch and set the remote as upstream, use
    git push --set-upstream origin fix-weapons-chart

# git push --set-upstream origin fix-weapons-chart
```

Your changes should now have been pushed successfully. You can keep committing
and pushing to your branch with just `git push` until you are done. Split the
work in reasonable commits to make tracking changes easier for the reviewer.

6. After finishing your work and pushing all commits, submit a pull request from
   the repo's GitHub page.

7. The reviewer will be notifyied and review your changes. He might ask you to
   change things, you can keep committing to your branch. Every time you push to
   your branch, the checks will be run by Github. Only if the checks pass the
   brach can be merged into master.

8. After any corrections, the reviewer will merge your branch into master and
   delete it afterwards. You will see all your commits in the master branch.
   Before doing any more work you should change back to the master branch
   locally and pull the changes (i.e. the merge).

```
git checkout master
```

At this point you are ready to start working on your next issue (like discribed
at `2.`) While this process might seem complicated at first, it really is not as
bad and ensures that the master branch always contains a working version that
can be deployed to the server. Furthermore any changes in master will have to be
reviewed by at least one second person, making the margin for mistakes smaller.


# Available data (e.g. for frontend development)

All data on a player can be accessed through the `PlayerInfo`
[PlayerInfo](https://godoc.org/github.com/megaclan3000/megaclan3000/internal/steamclient#PlayerInfo)
struct.  It contains four structs with different classes of data obtained through
four API endpoint from the steam API.  Their respective fields are documented in
godoc.
- [PlayerSummary](https://godoc.org/github.com/megaclan3000/megaclan3000/internal/steamclient#PlayerSummary)
- [UserStatsForGame](https://godoc.org/github.com/megaclan3000/megaclan3000/internal/steamclient#UserStatsForGame)
- [RecentlyPlayedGames](https://godoc.org/github.com/megaclan3000/megaclan3000/internal/steamclient#RecentlyPlayedGames)
- [PlayerHistory](https://godoc.org/github.com/megaclan3000/megaclan3000/internal/steamclient#PlayerHistory)

The `PlayerInfo`'s fields can be accessed in the `{{ template markers }}` as the object is passed to them.
