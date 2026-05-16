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
*Focus: Linux Hardening, Containerization basics, and Hardware setup.*
- [x] **Hardware & OS:** Raspberry Pi 5 (8GB) with NVMe Boot, OS Hardening (Headless Debian/Raspberry Pi OS Lite) with strict SSH Key Auth, and Static IP.
- [x] **Containerization Core:** Docker Engine installation, Docker Compose (IaC basics), and Portainer.
- [x] **Networking V1:** Tailscale (Mesh VPN for basic remote access) and Pi-hole (Network-wide Ad Blocking & DNS).

## ✅ Phase 2: Automation & Scripting (Eliminating Toil)
*Focus: Replacing manual maintenance with Python/Bash scripts and cron jobs.*
- [x] **Media Ops Automation:** Full *Arr Stack deployment, qBittorrent with VPN isolation, and Hardlinks setup.
- [x] **Scripting Dojo:** - *The Auditor:* Custom Python script using `psutil` to read Kernel sensors.
  - *The Alerting:* Telegram Bot API integration for critical alerts.
- [x] **Self-Healing:** Cron jobs for auto-updates (`apt`) and Docker cleanup.

## ✅ Phase 3: Security Fortress (Zero Trust Networking)
*Focus: Bypassing ISP CGNAT, Secret Management, and SSL/TLS.*
- [x] **Cloudflare Tunnels:** Deploy cloudflared container (Bypass CGNAT / No Open Ports) and Configure Inbound Rules (Zero Trust policy).
- [x] **Secret Management:** Migrate hardcoded credentials to `.env` files with Git `.gitignore` policy enforcement.
- [x] **Encryption:** Strict HTTPS/TLS enforcement (Cloudflare Edge Certificates) and Public Hostnames configuration.

## ✅ Phase 4: The Builder (DevOps & Coding)
*Focus: Transitioning from "Configuring Software" to "Building Software".*
- [x] **Golang (Go) Basics:** Syntax, Goroutines, and HTTP Standard Library. Project "Hello SRE" (API returning server telemetry).
- [x] **Advanced Docker Build:** Create custom Dockerfiles utilizing Multi-Stage Builds (Go -> Distroless/Alpine).
- [x] **Web Server Implementation:** Deploy Personal Portfolio (Hugo) via Nginx Container, fetching data from Go API.

## ✅ Phase 5: Deep Observability
*Focus: "If you can't measure it, you can't improve it." Moving beyond simple scripts.*
- [x] **The Stack:** Prometheus (Scraping metrics) and Grafana (Golden Signals Dashboard).
- [x] **Log Management:** Docker Logs aggregation (Loki) and Nginx Access/Error logs analysis (Geo-IP mapping).
- [x] **Health Checks:** Implement Docker Healthchecks and Uptime Kuma (External monitoring dashboard).

## ✅ Phase 6: Defensive Hacking (Red Teaming)
*Focus: Auditing the infrastructure from an attacker's perspective.*
- [x] **Vulnerability Scanning:** Audit the Raspberry Pi with Nmap and check for exposed headers.
- [x] **Hardening:** Fail2Ban (SSH brute-force protection) and Cloudflare WAF (Geo-Blocking and Bot Fight Mode).
- [x] **Training (TryHackMe):** Pre-Security (Networking basics) and Jr. Penetration Tester (Web Hacking).

## ✅ Phase 7: CI/CD & GitOps
*Focus: Automating the software delivery pipeline.*
- [x] **GitHub Actions:** CI to automate Go build and Linting on `git push`. CD to trigger deployment upon successful build.
- [x] **GitOps:** Watchtower to automatically update running containers when new images are pushed.
- [x] **Chaos Engineering:** Custom scripts to randomly restart containers to test resilience.

## ✅ Phase 8: AWS Cloud Foundation & IaC (MyssTic Warden)
*Focus: Expanding beyond the Home Lab into a highly available, Zero-Trust Public Cloud environment.*
- [x] **Infrastructure as Code (IaC):** Terraform provisioning (Free Tier EC2). Enterprise State via S3 Remote Backend (AES-256) + DynamoDB State Locking.
- [x] **Zero-Trust Networking:** Custom VPC architecture (Public DMZ + Isolated Private Subnets) with strict Security Groups.
- [x] **Persistence & Security:** Multi-AZ Amazon RDS (PostgreSQL) deployment, AWS DLM for automated backups, and AWS Secrets Manager integration.
- [x] **Serverless Observability:** Amazon CloudWatch alarms triggering SNS Topics, invoking an AWS Lambda (Python) to push real-time alerts to a Telegram Bot.

---

## 🚧 Phase 9: Configuration Management & Hybrid Cloud (Current Focus)
*Focus: Idempotent server configuration, enterprise GitOps, and Edge-to-Cloud connectivity.*
- [ ] **Ansible Automation:** Playbooks to setup the EC2 instance automatically (Docker, UFW) and standardize the Raspberry Pi (Edge) setup.
- [ ] **Hybrid Network & DNS:** Integrate EC2 and Raspberry Pi into a single Mesh VPN (Tailscale), migrate to AWS Route 53 DNS, and deploy Reverse Proxy (Nginx Proxy Manager / Traefik) on AWS.
- [ ] **Cloud Deployments & FinOps:** Migrate Portfolio to AWS for HA. Self-Host Vaultwarden in EC2, connected to the isolated RDS. Implement S3 Lifecycle Policies for backups via Duplicati.

---

## 🔮 Future Phases (The Path Forward)

### Phase 10: Orchestration (The Final Boss)
*Focus: Industry standard container orchestration.*
- [ ] **Kubernetes (K3s):** Migrate Docker Compose services to Lightweight Kubernetes.
- [ ] **ArgoCD:** Implement declarative GitOps for Kubernetes.
- [ ] **Cloud Scaling (Experiment):** Connect On-Premises K3s with AWS EKS (Hybrid Cluster).

### 🚀 Extras & Pro League (Horizon)
*Focus: Mastery, Advanced Cloud & Offensive Security.*
- [ ] **AWS Training:** AWS Solutions Architect Associate Course.
- [ ] **Hack The Box (HTB):** Tackle "Retired Machines" without guides.
- [ ] **Cloud Architecture:** Design Application Load Balancers (ALB) and Auto Scaling.
- [ ] **Enterprise Automation:** Deploy n8n or Apache Airflow.

---
*Roadmap updated automatically via CI/CD.*