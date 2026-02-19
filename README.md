# 🚀 SRE Portfolio (Monorepo)

A minimalist, high-performance portfolio built with a **Docs-as-Code** philosophy. Designed to showcase skills in **Site Reliability Engineering**, **Linux Administration**, and **Microservices Architecture**.

## 🛠 Tech Stack

### Frontend (Web)
* **Engine:** Hugo (Static Site Generator)
* **Theme:** PaperMod (Customized "Invictus" SRE Theme)
* **Server:** Nginx (Alpine)

### Backend (API)
* **Language:** Golang 1.22
* **Architecture:** REST API exposing system telemetry
* **Container:** Distroless/Alpine (Multi-stage build)

### Infrastructure
* **Host:** Raspberry Pi 5 (Home Lab Cluster)
* **Orchestration:** Docker Compose
* **Networking:** Cloudflare Tunnel (Zero Trust)
* **CI/CD:** GitHub Actions (Source Control)

## ⚡ Quick Start (Local)

1.  **Clone the monorepo:**
    ```bash
    git clone [https://github.com/ezequieldlv/portfolio-sre](https://github.com/ezequieldlv/portfolio-sre)
    cd portfolio-sre
    ```

2.  **Run with Docker Compose (Full Stack):**
    ```bash
    docker compose up -d
    ```
    *Access Frontend at `http://localhost:8095`*
    *Access Backend at `http://localhost:8585`*

## 🗺 Roadmap
- [x] Base Design & Content (Hugo)
- [x] Backend API Development (Golang)
- [x] Docker Containerization (Multi-Stage)
- [x] Deploy on Raspberry Pi 5 (Home Lab)
- [x] Network Segmentation & Nginx Reverse Proxy
- [ ] Deep Observability (Prometheus/Grafana)
- [ ] Kubernetes Migration (K3s)

---
*Created by Ez | 2026*
