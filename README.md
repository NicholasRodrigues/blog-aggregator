# Blogator

Blogator is a modern blogging platform designed to allow users to manage and share blogs efficiently. Built with Golang, it leverages PostgreSQL for data persistence and integrates with `goose` for database migrations and `sqlc` for type-safe SQL query management.

## Features

- **User Management**: Securely create and manage user profiles.
- **Feed Management**: Users can add and manage RSS feeds, representing blogs or other content sources.
- **Post Scraping**: Automatically scrapes new posts from added feeds.
- **Feed Following**: Users can follow specific feeds and manage their subscriptions.
- **API Key Authentication**: Each user is assigned a unique API key for secure interactions with the API.

## Getting Started

### Prerequisites

Ensure you have the following installed:
- Go (version 1.16 or later)
- PostgreSQL
- Goose (for database migrations)
- SQLC (for generating Go code from SQL)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourgithubusername/blogator.git
2. **Navigate to the project directory:**
   ```bash
   cd blogator
3. **Set up the environment variables:**
   Create a `.env` file in the project root and update the following variables:
   ```plaintext
   DB_URL="postgres://username:password@localhost:5432/blogator"
4. **Run Database Migrations:**
   ```bash
   goose postgres $DB_URL up
5. **Generate Go code from SQL:**
   ```bash
   sqlc generate
6. **Build the project:**
   ```bash
   go build -o blogator
7. **Run the server:**
   ```bash
   ./blogator

## API Endpoints

All API requests are prefixed with `/app/v1`.

### Users

- **Create User**: `POST /users`
  - **Body**: `{"name": "username"}`
  - **Response**: Returns the created user object with API key.
  
- **Get User by API Key**: `GET /users`
  - **Headers**: `Authorization: ApiKey <key>`
  - **Response**: Returns the user object.

### Feeds

- **Create Feed**: `POST /feeds`
  - **Body**: `{"name": "Feed Name", "url": "https://feedurl.com/rss"}`
  - **Auth Required**
  
- **Get Feeds**: `GET /feeds`
  - **Auth Not Required**
  
- **Follow Feed**: `POST /feeds/follow`
  - **Body**: `{"feed_id": "feed_uuid"}`
  - **Auth Required**

- **Unfollow Feed**: `DELETE /feed_follows/{feedFollowID}`
  - **Auth Required**

### Posts

- **Get Posts by User**: `GET /posts`
  - **Query Parameters**: `limit` (optional)
  - **Auth Required**

## Database Schema

- **Users**: Stores user data including a generated API key.
- **Feeds**: Contains details of the RSS feeds added by users.
- **Posts**: Captures posts fetched from the RSS feeds.

### Goose Migration(Inside Schemas)
goose postgres postgres://postgres:postgres@localhost:5432/blogator up

### Start Psql Server
psql postgres://postgres:postgres@localhost:5432/blogator

### SQLC 
sqlc generate

## Contributing

Contributions to improve Blogator are welcome. Please fork the repository and submit a pull request with your changes.

## License

Distributed under the MIT License. See `LICENSE` for more information.
