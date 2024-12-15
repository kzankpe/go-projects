# TMDB CLI Tool Solutions

Project Url : https://roadmap.sh/projects/tmdb-cli

Golang CLI solution for the TMDB CLI Tool [challenge](https://roadmap.sh/projects/tmdb-cli) from [roadmap.sh](https://roadmap.sh/).
A simple command line interface (CLI) to fetch data from The Movie Database (TMDB) and display it in the terminal.

## Features

Tool to fetch data from The Movie Database (TMDB) and display it in the terminal. The user filters the category of data to fetch.

## Installation

To run this application, follow these steps:

```bash
git clone https://github.com/kzankpe/go-projects.git
```
Run the following command to build and run the project:

```bash
cd tmdb-cli
go build -o tmdb-app
./tmdb-app --help # To see the list of available commands
```

## Usage

```bash
tmdb-app --type "playing"
tmdb-app --type "popular"
tmdb-app --type "top"
tmdb-app --type "upcoming"
```