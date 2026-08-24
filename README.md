
# sharingvision-be-test

## Setup
```bash
# Load environment variables
cp .env.example .env
# Adjust .env based on your local database credentials

# Run migrations and seeders
go run main.go migrate
go run main.go seed

# Start server
go run main.go
```

## Endpoints

### Create
- `POST /article`
- `POST /article/`

**Request Body:**
  ```json
  {
    "title": "string (min 20 chars)",
    "content": "string (min 200 chars)",
    "category": "string (min 3 chars)",
    "status": "publish | draft | thrash"
  }
  ```

**Response:**
  ```json
{}
  ```

### Read Paged
- `GET /article/<limit>/<offset>`

**Example:**
- `GET /article/5/0`

### Read by ID
- `GET /article/<id>`

### Update
- `PUT /article/<id>`
- `PATCH /article/<id>`
- `POST /article/<id>/update`

**Request Body:** (Same as Create)

### Delete
- `DELETE /article/<id>`
- `POST /article/<id>?action=delete`