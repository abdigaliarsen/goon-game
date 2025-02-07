# Project Setup Guide

## 1. Configure Environment Variables

Run the following commands to set up environment variables:

```bash
cp config/discord_bot/.env.example config/discord_bot/.env && cp config/wikipedia/.env.example config/wikipedia/.env
```

Next, update these fields in the respective `.env` files:

- DISCORD_API_TOKEN
- DISCORD_APPLICATION_ID
- DISCORD_PUBLIC_KEY
- DISCORD_DEFAULT_CHAT_IDS

All other configurations are pre-configured and should work out of the box.

---

## 2. Running the Application

If Kafka is already running on your system, restart all services for a clean start:

```bash
sudo -i && make hard-restart
```

If Kafka is not running, start the application with:

```bash
make run
```

---

## 3. Architecture Overview

- **Kafka**: Handles high-throughput data streaming from the Wikipedia Service, ensuring efficient processing of large data volumes.

- **gRPC**: Provides fast, reliable communication between services, optimized for handling language-specific parameters.

- **Redis**: Used for caching the history of language changes and storing the current language parameter. Redis supports up to 10,000,000 entries with minimal memory usage (~10MB of RAM), ensuring quick data retrieval.

![Screenshot from 2025-02-07 07-11-45](https://github.com/user-attachments/assets/f2b7dcee-923a-454b-91ee-d2ffc24c0b99)

## 4. Further Scaling

- **Multithreaded Operations**: Implementing a dedicated `TaskService` to manage complex operations using Go channels and goroutines will enhance concurrency and improve performance for high-load scenarios.

- **Kibana Integration for Logging**: Since Kafka is already integrated, incorporating the ELK (Elasticsearch, Logstash, Kibana) stack will provide robust monitoring and centralized logging, enabling better visualization and analysis of logs.

- **Horizontal Scaling**: Deploy multiple instances of the microservices and use container orchestration tools like Kubernetes or Docker Swarm to manage scaling automatically based on traffic and resource utilization.

- **Database Optimization**: For handling larger datasets, consider sharding or partitioning in Redis and optimize Kafka topic partitions for parallel processing.

- **CI/CD Pipeline Integration**: Automate testing, building, and deployment processes using CI/CD tools like GitHub Actions, Jenkins, or GitLab CI to ensure rapid and reliable delivery of updates.
