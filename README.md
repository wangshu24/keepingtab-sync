# KeepingTab Sync Service

A lightweight Go-based background service that handles cross-device synchronization of priority tabs for the KeepingTab Chrome extension. In the MVP architecture, this service works alongside the Node.js API using a shared SQLite database for simple, cost-effective synchronization.

## 🎯 Purpose

The KeepingTab Sync Service is designed to:

- **Synchronize priority tabs** across multiple devices and browser instances
- **Merge tab collections** intelligently when conflicts arise
- **Maintain session state** for seamless cross-device experiences
- **Process sync operations** via shared SQLite database
- **Provide lightweight background processing** with minimal resource usage

## 🏗️ Architecture Overview

### MVP System Components (Lean Architecture)

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Chrome Ext     │    │  keepingtab-api │    │ keepingtab-sync │
│  (Frontend)     │◄──►│  (Node.js API)  │◄──►│  (Go Service)   │
└─────────────────┘    │ ┌─────────────┐ │    └─────────────────┘                 
                       │ │ In-Memory   │ │              |       
                       │ │ Cache (Map) │ │              |       
                       │ └─────────────┘ │              |       
                       └─────────────────┘              |    
                                │                       │
                                ▼                       │
                       ┌─────────────────┐              │
                       │     SQLite      │◄─────────────┘
                       │ (Fly.io Volume) │
                       └─────────────────┘
```

### MVP Data Flow

1. **Chrome Extension** → Manages local priority tabs (max 3 tabs)
2. **keepingtab-api** → REST API with SQLite persistence and in-memory cache
3. **keepingtab-sync** → Background service reading from shared SQLite
4. **SQLite** → Single source of truth on Fly.io volume
5. **In-Memory Cache** → Hot data cache in Node.js API process

## 🔧 Technical Stack

- **Language**: Go 1.21
- **Database**: PostgreSQL (via `DATABASE_URL`)
- **Queue**: Redis (via `REDIS_URL`)
- **Deployment**: Fly.io with Docker
- **Architecture**: Microservice pattern

## 📁 Project Structure

```
keepingtab-sync/
├── main.go          # Entry point and service initialization
├── go.mod           # Go module dependencies
├── Dockerfile       # Container build configuration
├── fly.toml         # Fly.io deployment configuration
└── README.md        # This documentation
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Redis instance (local or remote)
- PostgreSQL database
- Environment variables configured

### Environment Variables

```bash
# Required environment variables
REDIS_URL=redis://localhost:6379
DATABASE_URL=postgres://user:password@localhost:5432/keepingtab
```

### Local Development

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd keepingtab-sync
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set environment variables:**
   ```bash
   export REDIS_URL="redis://localhost:6379"
   export DATABASE_URL="postgres://user:password@localhost:5432/keepingtab"
   ```

4. **Run the service:**
   ```bash
   go run main.go
   ```

### Docker Development

1. **Build the container:**
   ```bash
   docker build -t keepingtab-sync .
   ```

2. **Run with environment variables:**
   ```bash
   docker run -e REDIS_URL="redis://host:6379" \
              -e DATABASE_URL="postgres://user:pass@host:5432/db" \
              keepingtab-sync
   ```

## 🔄 Sync Logic (Planned Implementation)

### Tab Synchronization Flow

1. **Change Detection**: Chrome extension detects priority tab changes
2. **API Push**: Extension pushes changes to `keepingtab-api` via `/v1/sync/push`
3. **Queue Processing**: API enqueues sync events to Redis
4. **Background Processing**: `keepingtab-sync` consumes queue messages
5. **Conflict Resolution**: Service merges changes using intelligent algorithms
6. **State Persistence**: Final state saved to PostgreSQL
7. **Real-time Updates**: Other devices receive updates via polling or WebSocket

### Merge Strategies

- **Last-Write-Wins**: Simple timestamp-based conflict resolution
- **Priority-Based**: Preserve higher priority tabs during conflicts
- **Device-Aware**: Consider device context for intelligent merging
- **User Preferences**: Allow user-defined merge behavior

## 📊 Data Models

### Priority Tab Structure
```go
type PriorityTab struct {
    ID          int       `json:"id"`
    UserID      string    `json:"user_id"`
    TabID       string    `json:"tab_id"`
    Title       string    `json:"title"`
    URL         string    `json:"url"`
    FavIconURL  string    `json:"favicon_url"`
    Key         int       `json:"key"`         // 0, 1, 2, 3
    DeviceID    string    `json:"device_id"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### Sync Event Structure
```go
type SyncEvent struct {
    EventID     string    `json:"event_id"`
    UserID      string    `json:"user_id"`
    DeviceID    string    `json:"device_id"`
    EventType   string    `json:"event_type"`  // "add", "remove", "update"
    TabData     PriorityTab `json:"tab_data"`
    Timestamp   time.Time `json:"timestamp"`
}
```

## 🚀 Deployment

### Fly.io Deployment

The service is configured for deployment on Fly.io:

```bash
# Deploy to Fly.io
fly deploy

# Set environment variables
fly secrets set REDIS_URL="redis://your-redis-url"
fly secrets set DATABASE_URL="postgres://your-db-url"
```

### Configuration

- **App Name**: `keepingtab-sync`
- **Primary Region**: `iad` (US East)
- **Process**: Single `sync` process running the main binary

## 🔮 Roadmap

### Phase 1: Core Sync (Current)
- [ ] Redis queue consumer implementation
- [ ] PostgreSQL database schema and connections
- [ ] Basic tab merge logic
- [ ] Error handling and logging

### Phase 2: Advanced Features
- [ ] Real-time WebSocket updates
- [ ] Intelligent conflict resolution
- [ ] Device management and identification
- [ ] Sync history and rollback capabilities

### Phase 3: Optimization
- [ ] Performance monitoring and metrics
- [ ] Horizontal scaling support
- [ ] Advanced caching strategies
- [ ] Backup and disaster recovery

## 🤝 Integration with KeepingTab Ecosystem

### Chrome Extension Integration
- Extension stores priority tabs locally using Chrome Storage API
- Syncs changes via REST API calls to `keepingtab-api`
- Maintains 3-tab limit with keyboard shortcuts (1-3, 0)

### API Service Integration
- Receives sync requests from Chrome extension
- Handles user authentication and authorization
- Enqueues sync events to Redis for background processing
- Provides endpoints for pulling latest tab state

## 📝 Development Notes

- Service currently in early development stage
- Main functionality is placeholder implementation
- Requires Redis and PostgreSQL setup for full functionality
- Designed for horizontal scaling and high availability

## 🔒 Security Considerations

- All sync operations require user authentication
- Tab data encrypted in transit and at rest
- Device identification for security and conflict resolution
- Rate limiting to prevent abuse

---

**Status**: 🚧 Under Development  
**Last Updated**: 2025-09-07  
**Go Version**: 1.21  
**Deployment**: Fly.io Ready
