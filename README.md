# Go-Contix

A tools for checking the target performance is or not available for sales

## Usage

Edit the config file first

    vim data/cron-task.yaml

Help

    contix --help

Fetch latest proxy list

    contix proxy fetch

Update the latest and success connected proxy to pool

    contix proxy update

Show the status of checking task list from the configuration file

    contix cron list

Start the cron

    contix cron run

Send the testing email

    contix mail send

Show the rendered email content base on dummy data

    contix mail render

## Develop

Install the tools

    make develop

Install the vendor

    make vendor

Clean the builds and assets

    make clean

Run the watcher

    make watch

Try to execute the application

    ./go-contix --help

Develop editor

    cd editor
    make watch

## Release

Build

    make release

Release file or Windows like cmder and fonts

    make windows
