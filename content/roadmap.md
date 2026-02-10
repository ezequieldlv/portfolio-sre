---
title: "My SRE Roadmap"
draft: false
hideMeta: true
summary: "From Localhost to Cloud Native: My 9-Phase journey to SRE."
---

> **Philosophy:** *"SRE is what happens when you ask a software engineer to design an operations team."*

This roadmap documents my engineering journey. It focuses on **infrastructure resilience**, **automation**, **security**, and **observability**.

---

## ✅ Phase 1: The Foundation & Infrastructure
*Focus: Linux Hardening & Containerization.*
- [x] **Hardware Setup:** Raspberry Pi 5 (8GB) with SSD Boot.
- [x] **OS Hardening:** Debian (Headless) + SSH Key Auth (No root login).
- [x] **Containerization:** Docker Engine + Docker Compose (IaC) + Portainer.
- [x] **Networking:** Tailscale Mesh VPN + Pi-hole (DNS AdBlocking).

## ✅ Phase 2: Automation & Scripting
*Focus: Eliminating Toil with Code.*
- [x] **Media Ops:** Automated *Arr Stack deployment with isolated networking.
- [x] **The Auditor:** Custom Python script (using `psutil`) to monitor Kernel sensors.
- [x] **The Alerting:** Telegram Bot API integration for critical alerts.
- [x] **Self-Healing:** Cron jobs for auto-updates and Docker cleanup.

## ✅ Phase 3: Security Fortress (Zero Trust)
*Focus: Hardening, Encryption & Access Control.*
- [x] **Connectivity:** Cloudflare Tunnels (Bypass CGNAT securely).
- [x] **Active Defense:**
    - [x] **Fail2Ban:** SSH brute-force protection logic.
    - [x] **Cloudflare WAF:** Web Application Firewall rules configured.
- [x] **Secrets:** Migration to `.env` files + `.gitignore` policy enforcement.
- [x] **Encryption:** Strict HTTPS/TLS enforcement for all services.

---

## 🚧 Phase 4: The Builder (Current)
*Focus: Transitioning from "Configuring" to "Building".*
- [ ] **Golang Basics:** Syntax, Goroutines, HTTP Library.
- [ ] **Project:** "Hello SRE" - Lightweight Web Server.
- [ ] **Advanced Docker:** Multi-Stage Builds (Go -> Alpine/Scratch).
- [ ] **Web Server:** Nginx Reverse Proxy configuration.

---

## 🔮 Future Phases (Backlog)

### Phase 5: Deep Observability
- **Stack:** Prometheus + Grafana (Golden Signals).
- **Logs:** Centralized aggregation (Loki).

### Phase 6: Defensive Hacking (Red Team)
- **Audit:** Vulnerability scanning with Nmap/OpenVAS.
- **Penetration Testing:** Audit the infrastructure from an attacker's perspective.

### Phase 7: CI/CD & GitOps
- **Pipeline:** GitHub Actions for automated testing/building.
- **GitOps:** Auto-update containers on `git push`.

### Phase 8: Hybrid Cloud (AWS)
- **Backup:** Encrypted backups to AWS S3 (Glacier).
- **IaC:** Terraform provisioning (EC2 Free Tier).

### Phase 9: Orchestration (Final Boss)
- **K8s:** Migration to **K3s** (Lightweight Kubernetes).
- **ArgoCD:** GitOps for Kubernetes.