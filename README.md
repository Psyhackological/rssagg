<h1 align="center">GoFeed2DB</h1>
<p align="center">
  <img width="20%" align="center" src="media/rssagg_logo_plain.svg" alt="RSSaggLogo">
</p>

> Go lang meets RSS feeds and keeps it into PotsgreSQL.

## Features

## Dependencies
### Dev tools
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv) - a Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).
- [sqlc](https://sqlc.dev/) - generates fully type-safe idiomatic code from SQL. 
- [goose](https://pressly.github.io/goose/) - a database migration tool. Manage your database schema by creating incremental SQL changes and/or Go functions.
- [podman](https://podman.io/) - a tool for managing OCI containers and pods and in this example: PostgreSQL container.

Example usage:

#### SQLc
```bash
sqlc generate
```

#### Goose

##### Down

```bash
goose -dir ./sql/schema postgres postgres://postgres:your_secure_password@localhost:5432/rssagg down
```

Output:
```text
2025/07/07 19:06:06 OK   006_posts.sql (25.68ms)
```

##### Up

```bash
goose -dir ./sql/schema postgres postgres://postgres:your_secure_password@localhost:5432/rssagg up
```

Output:
```text
2025/07/07 19:06:08 OK   006_posts.sql (36.42ms)
2025/07/07 19:06:08 goose: successfully migrated database to version: 6
```

### Backend
- [PostgreSQL](https://www.postgresql.org/) - the database of choice that is powerful, open source object-relational database system with over 35 years of active development that has earned it a strong reputation for reliability, feature robustness, and performance.
- [chi](https://go-chi.io/#/README) - A lightweight, idiomatic and composable router for building Go HTTP services.

## Endpoints
| Middleware Auth? | HTTP Method | Endpoint Path                     | Description                                                                 |
|------------------|-------------|-----------------------------------|-----------------------------------------------------------------------------|
| ❌ No            | `GET`       | `/v1/healthz`                     | Service health check                                                        |
| ❌ No            | `GET`       | `/v1/err`                         | Error simulation endpoint                                                   |
| ❌ No            | `POST`      | `/v1/users`                       | Create new user                                                             |
| ✅ Yes           | `GET`       | `/v1/users`                       | Retrieve authenticated user's data                                          |
| ✅ Yes           | `POST`      | `/v1/feeds`                       | Create new RSS feed subscription                                            |
| ❌ No            | `GET`       | `/v1/feeds`                       | Get list of available RSS feeds                                             |
| ✅ Yes           | `GET`       | `/v1/posts`                       | Get posts from followed feeds                                               |
| ✅ Yes           | `POST`      | `/v1/feed_follows`                | Follow new RSS feed                                                         |
| ✅ Yes           | `GET`       | `/v1/feed_follows`                | Get list of followed feeds                                                  |
| ✅ Yes           | `DELETE`    | `/v1/feed_follows/{feedFollowID}` | Unfollow specific feed using ID parameter                                   |

## Installation

## Usage

## Ideas for extending the project

* Support [pagination](https://nordicapis.com/everything-you-need-to-know-about-api-pagination/) of the endpoints that can return many items
* Support different options for sorting and filtering posts using query parameters
* Classify different types of feeds and posts (e.g. blog, podcast, video, etc.)
* Add a CLI client that uses the API to fetch and display posts, maybe it even allows you to read them in your terminal
* Scrape lists of feeds themselves from a third-party site that aggregates feed URLs
* Add support for other types of feeds (e.g. Atom, JSON, etc.)
* Add integration tests that use the API to create, read, update, and delete feeds and posts
* Add bookmarking or "liking" to posts
* Create a simple web UI that uses your backend API
## License
[![GNU GPLv3 Image](https://www.gnu.org/graphics/gplv3-with-text-136x68.png)](https://choosealicense.com/licenses/gpl-3.0/)

Software licensed under the [GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/).
