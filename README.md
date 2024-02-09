# Animals Application

This is a simple application written in Go that uses a graph database to store and retrieve data, a REST API to serve the data, and a simple web interface to interact with the data.

Technologies used:

- Go
  - [https://github.com/labstack/echo](echo)
  - [https://github.com/joeshaw/envdecode](envdecode)
- Neo4j
- HTMX

Project structure is from [https://github.com/learning-cloud-native-go/myapp](here).

## Data Model

The data model used for this project is fairly simple and consists of three nodes: `Animal`, `Person` and `Food`. The relationships between these nodes are also simple: a `Person` `CARES_FOR` `Animal` and `Animal` `EATS` `Food`.
