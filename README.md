# prusk-sql

Repo for my sql cmd utility written in Go

This is 100% a work in progress and contains very bad code! It's just for me to play around with Go.

This utility is designed to automate the frequent queries I run when testing SQL Server running in containers/kubernetes

It will test the connection and then retrieve the SQL version. It also has an option to list the databases in the instance.

### Usage

To connect to SQL running in a container and retreive the SQL version: -

    prusk-sql -server 127.0.0.1 -port 15789 -user sa -password Testing1122

Add the databases flag to list all databases in the SQL instance: -

    prusk-sql -server 127.0.0.1 -port 15789 -user sa -password Testing1122 -databases