---
title: "My SRE & DevSecOps Roadmap"
draft: false
hideMeta: true
summary: "From Localhost to Cloud Native: My journey to SRE and Offensive Security."
---

> **Philosophy:** *"Reliability is not an accident. It is a feature we build."*

This roadmap documents my engineering journey from bare-metal infrastructure to cloud-native orchestration. It focuses on **infrastructure resilience**, **automation**, **Zero Trust security**, and **observability**.

---

## ✅ Phase 1: The Foundation & Infrastructure
*Focus: Linux Hardening & Containerization.*
- [x] **Hardware Setup:** Raspberry Pi 5 (4GB) with USB 3.0 Boot (SSD).
- [x] **OS Hardening:** Debian Bookworm (Headless) + SSH Key Auth.
- [x] **Containerization:** Docker Engine + Docker Compose (IaC) + Portainer.
- [x] **Networking:** Tailscale Mesh VPN + Pi-hole (DNS AdBlocking).
- [x] **Storage (NAS):** Samba configuration for cross-platform file sharing.

## ✅ Phase 2: Automation & Scripting
*Focus: Eliminating Toil with Code.*
- [x] **Media Ops:** Automated *Arr Stack deployment with isolated networking.
- [x] **Observability V1:** Centralized Dashboard implementation (Homepage).
- [x] **The Auditor:** Custom Python script (using `psutil`) to monitor Kernel sensors.
- [x] **The Alerting:** Telegram Bot API integration for critical alerts.
- [x] **Self-Healing:** Cron jobs for auto-updates (`apt`) and Docker cleanup.

## ✅ Phase 3: Security Fortress (Zero Trust)
*Focus: Hardening, Encryption & Access Control.*
- [x] **Connectivity:** Cloudflare Tunnels (Bypass CGNAT securely / No open ports).
- [x] **Active Defense:**
    - [x] **Fail2Ban:** SSH brute-force protection logic.
    - [x] **Cloudflare WAF:** Web Application Firewall rules (Geo-Blocking).
- [x] **Secrets:** Migration to `.env` files + `.gitignore` policy enforcement.
- [x] **Encryption:** Strict HTTPS/TLS enforcement for all services.

---

## 🚧 Phase 4: The Builder (Current Focus)
*Focus: Transitioning from "Configuring" to "Building Software".*
- [ ] **Golang Basics:** Syntax, Goroutines, HTTP Standard Library.
- [ ] **Project A:** "Hello SRE" - API returning server telemetry in JSON.
- [ ] **Advanced Docker:** Multi-Stage Builds (Go -> Distroless/Alpine).
- [ ] **Web Server:** Nginx Reverse Proxy for this Portfolio (Self-Hosted).

---

## 🔮 Future Phases (The Path Forward)

### Phase 5: Deep Observability
- **Stack:** Prometheus + Grafana (Golden Signals Dashboard).
- **Logs:** Centralized aggregation with Loki.
- **Health:** Uptime Kuma implementation.

### Phase 6: Defensive Hacking (Red Teaming)
- **Audit:** Vulnerability scanning with Nmap/OpenVAS.
- **Training:** TryHackMe (Pre-Security & Jr. Penetration Tester paths).
- **Hardening:** Auditing exposed headers and configurations.

### Phase 7: CI/CD & GitOps
- **Pipeline:** GitHub Actions for automated testing/building on `git push`.
- **GitOps:** Automated container updates (Watchtower) & IaC synchronization.
- **Chaos:** Resilience testing via random service restarts.

### Phase 8: Hybrid Cloud & IaC
- **Backup:** Encrypted backups to AWS S3 (Glacier).
- **IaC:** Terraform provisioning (AWS EC2).
- **Config Mgmt:** Ansible playbooks for automated server setup.

### Phase 9: Orchestration (The Migration)
- **K8s:** Migration to **K3s** (Lightweight Kubernetes).
- **ArgoCD:** Declarative GitOps for Kubernetes.
- **Cloud:** AWS Solutions Architect Associate preparation.

### Phase 10: Endgame (Pro League)
- **Offensive Security:** Hack The Box (HTB) - Retired Machines.
- **Architecture:** Hybrid Cluster design (On-Prem K3s + Cloud).

---
*Roadmap updated automatically via CI/CD.*