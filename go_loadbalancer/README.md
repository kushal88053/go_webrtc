# Go Load Balancer

This is a simple round-robin load balancer written in Go. It routes incoming HTTP requests to a pool of backend servers using the round-robin strategy.

## Features

- Round-robin load balancing
- Reverse proxy to forward requests to backend servers
- Server health check placeholder (`IsAlive`)
- Easy to extend for more sophisticated logic (e.g., real health checks, weighted load balancing)

## Project Structure

go_loadbalancer/
├── main.go
└── README.md

## Requirements

- Go 1.18 or higher

## Getting Started

### 1. Clone or create the project

If you haven't yet:

```bash
mkdir go_loadbalancer
cd go_loadbalancer
```
