# QueueCraft

QueueCraft is a lightweight, modular job queue and worker system built in Go. It helps you run background jobs (like sending emails or processing data) using a pool of workers — perfect for building reliable async workflows in your Go applications.

---

## 🚀 Features

- ⚡️ Fast and efficient worker pool (goroutines + channels)
- 🧱 PostgreSQL-backed persistent job storage (using Neon DB)
- 🧩 Uses pgx as the database driver
- 🔁 Built-in support for retries, job delays, and scheduling
- 🛠️ Custom job handler support (`send_email`, `generate_report`, etc.)
- 📜 Database migrations managed with Goose
- 🧪 Testable design with clean abstractions
- 🧹 Graceful shutdown and in-progress job recovery

---

## 📦 Tech Stack

- [Go](https://golang.org/)
- [pgx](https://github.com/jackc/pgx) as the PostgreSQL driver
- [Neon (PostgreSQL)](https://neon.tech/)
- [Goose](https://github.com/pressly/goose) for migrations
- [SQLC (optional)](https://sqlc.dev/) for type-safe SQL
- Go standard library

---

## 🛠️ Getting Started

### Prerequisites

- Go 1.20+
- A Neon PostgreSQL instance
- [Goose CLI](https://github.com/pressly/goose#installing) installed
- (Optional) SQLC if you're using generated code

---

### Clone the repo

```bash
git clone https://github.com/matthewyoungjr/queuecraft.git
cd queuecraft
