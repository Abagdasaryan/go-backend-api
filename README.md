# Go Backend API

A small, elegant Go REST API built with Gin framework, designed for easy deployment to Railway.

## Features

- 🚀 **Fast & Lightweight**: Built with Go and Gin framework
- 🔒 **CORS Enabled**: Ready for frontend integration
- 📊 **Health Check**: Built-in health monitoring endpoint
- 🎯 **RESTful**: Clean API design with proper HTTP status codes
- 🚢 **Railway Ready**: Optimized for Railway deployment
- 🐳 **Docker Support**: Multi-stage Dockerfile included

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/` | Root welcome message |
| `GET` | `/api/v1/` | API welcome with endpoint list |
| `GET` | `/api/v1/health` | Health check with uptime |
| `GET` | `/api/v1/echo/:message` | Echo a message back |
| `POST` | `/api/v1/data` | Create new data |
| `GET` | `/api/v1/data` | Retrieve all data |

## Quick Start

### Local Development

1. **Install Go** (version 1.21 or higher)
2. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd go-backend-api
   ```

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

5. **Test the API**
   ```bash
   curl http://localhost:8080/api/v1/health
   ```

### Environment Variables

Create a `.env` file (optional for local development):

```env
PORT=8080
GIN_MODE=debug
```

## Deployment to Railway

### Option 1: Using Railway CLI

1. **Install Railway CLI**
   ```bash
   npm install -g @railway/cli
   ```

2. **Login to Railway**
   ```bash
   railway login
   ```

3. **Initialize and deploy**
   ```bash
   railway init
   railway up
   ```

### Option 2: Using Railway Dashboard

1. **Push your code to GitHub**
2. **Connect your repository to Railway**
3. **Railway will automatically detect the Go project and deploy**

### Option 3: Using Docker

1. **Build the Docker image**
   ```bash
   docker build -t go-backend-api .
   ```

2. **Run locally with Docker**
   ```bash
   docker run -p 8080:8080 go-backend-api
   ```

## Testing the API

### Health Check
```bash
curl https://your-app.railway.app/api/v1/health
```

### Echo Message
```bash
curl https://your-app.railway.app/api/v1/echo/hello
```

### Create Data
```bash
curl -X POST https://your-app.railway.app/api/v1/data \
  -H "Content-Type: application/json" \
  -d '{"name": "test", "value": 123}'
```

### Get All Data
```bash
curl https://your-app.railway.app/api/v1/data
```

## Project Structure

```
go-backend-api/
├── main.go           # Main application file
├── go.mod            # Go module file
├── go.sum            # Go dependencies checksum
├── Dockerfile        # Multi-stage Docker build
├── .dockerignore     # Docker ignore file
├── railway.toml      # Railway configuration
├── .gitignore        # Git ignore file
└── README.md         # This file
```

## Dependencies

- **Gin**: HTTP web framework
- **godotenv**: Environment variable loading

## Railway Configuration

The `railway.toml` file configures:
- **Builder**: Uses nixpacks for automatic Go detection
- **Health Check**: Monitors `/api/v1/health` endpoint
- **Restart Policy**: Automatically restarts on failure
- **Port**: Automatically detects PORT environment variable

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test locally
5. Submit a pull request

## License

MIT License - feel free to use this project as a starting point for your own APIs!
