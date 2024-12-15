# GitHub User Activity Solutions

Project Url : https://roadmap.sh/projects/github-user-activity

Golang CLI solution for the github-activities [challenge](https://roadmap.sh/projects/github-user-activity) from [roadmap.sh](https://roadmap.sh/).

The application should run from the command line, accept the GitHub username as an argument, fetch the user’s recent activity using the GitHub API, and display it in the terminal.

## Features

List recent activities of a user.

## Installation

To run this application, follow these steps:

```bash
git clone https://github.com/kzankpe/go-projects.git
```
Run the following command to build and run the project:

```bash
cd github-activities
go build -o gh-activities
./gh-activities --help # To see the list of available commands
```

## Usage

```bash
gh-activities [username]
```
