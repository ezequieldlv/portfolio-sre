---
title: "My SRE & DevSecOps Roadmap"
draft: false
hideMeta: true
summary: "From Localhost to Cloud Native: My journey to SRE and Offensive Security."
---

> **Philosophy:** *"Reliability is not an accident. It is a feature we build."*

This roadmap documents my engineering journey from bare-metal infrastructure to cloud-native orchestration. It focuses on **infrastructure resilience**, **automation**, **Zero Trust security**, and **observability**.

---

## 🧱 Phase 1: The Foundation & Infrastructure ✅
*Focus: Linux Hardening, Containerization basics, and Hardware setup.*
- [x] **Hardware & OS:** Raspberry Pi 5 (8GB) with NVMe Boot, OS Hardening (Headless Debian/Raspberry Pi OS Lite) with strict SSH Key Auth, and Static IP.
- [x] **Containerization Core:** Docker Engine installation, Docker Compose (IaC basics), and Portainer.
- [x] **Networking V1:** Tailscale (Mesh VPN for basic remote access) and Pi-hole (Network-wide Ad Blocking & DNS).

## ⚙️ Phase 2: Automation & Scripting (Eliminating Toil) ✅
*Focus: Replacing manual maintenance with Python/Bash scripts and cron jobs.*
- [x] **Media Ops Automation:** Full *Arr Stack deployment, qBittorrent with VPN isolation, and Hardlinks setup.
- [x] **Scripting Dojo:** - *The Auditor:* Custom Python script using `psutil` to read Kernel sensors.
  - *The Alerting:* Telegram Bot API integration for critical alerts.
- [x] **Self-Healing:** Cron jobs for auto-updates (`apt`) and Docker cleanup.

## 🔒 Phase 3: Security Fortress (Zero Trust Networking) ✅
*Focus: Bypassing ISP CGNAT, Secret Management, and SSL/TLS.*
- [x] **Cloudflare Tunnels:** Deploy cloudflared container (Bypass CGNAT / No Open Ports) and Configure Inbound Rules (Zero Trust policy).
- [x] **Secret Management:** Migrate hardcoded credentials to `.env` files with Git `.gitignore` policy enforcement.
- [x] **Encryption:** Strict HTTPS/TLS enforcement (Cloudflare Edge Certificates) and Public Hostnames configuration.

## 💻 Phase 4: The Builder (DevOps & Coding) ✅
*Focus: Transitioning from "Configuring Software" to "Building Software".*
- [x] **Golang (Go) Basics:** Syntax, Goroutines, and HTTP Standard Library. Project "Hello SRE" (API returning server telemetry).
- [x] **Advanced Docker Build:** Create custom Dockerfiles utilizing Multi-Stage Builds (Go -> Distroless/Alpine).
- [x] **Web Server Implementation:** Deploy Personal Portfolio (Hugo) via Nginx Container, fetching data from Go API.

## 📊 Phase 5: Deep Observability ✅
*Focus: "If you can't measure it, you can't improve it." Moving beyond simple scripts.*
- [x] **The Stack:** Prometheus (Scraping metrics) and Grafana (Golden Signals Dashboard).
- [x] **Log Management:** Docker Logs aggregation (Loki) and Nginx Access/Error logs analysis (Geo-IP mapping).
- [x] **Health Checks:** Implement Docker Healthchecks and Uptime Kuma (External monitoring dashboard).

## 🥷 Phase 6: Defensive Hacking (Red Teaming) ✅
*Focus: Auditing the infrastructure from an attacker's perspective.*
- [x] **Vulnerability Scanning:** Audit the Raspberry Pi with Nmap and check for exposed headers.
- [x] **Hardening:** Fail2Ban (SSH brute-force protection) and Cloudflare WAF (Geo-Blocking and Bot Fight Mode).
- [x] **Training (TryHackMe):** Pre-Security (Networking basics) and Jr. Penetration Tester (Web Hacking).

## 🚢 Phase 7: CI/CD & GitOps ✅
*Focus: Automating the software delivery pipeline.*
- [x] **GitHub Actions:** CI to automate Go build and Linting on `git push`. CD to trigger deployment upon successful build.
- [x] **GitOps:** Watchtower to automatically update running containers when new images are pushed.
- [x] **Chaos Engineering:** Custom scripts to randomly restart containers to test resilience.

## 🌩️ Phase 8: AWS Cloud Foundation & IaC (MyssTic Warden) ✅
*Focus: Expanding beyond the Home Lab into a highly available, Zero-Trust Public Cloud environment.*
- [x] **Infrastructure as Code (IaC):** Terraform provisioning (Free Tier EC2). Enterprise State via S3 Remote Backend (AES-256) + DynamoDB State Locking.
- [x] **Zero-Trust Networking:** Custom VPC architecture (Public DMZ + Isolated Private Subnets) with strict Security Groups.
- [x] **Persistence & Security:** Multi-AZ Amazon RDS (PostgreSQL) deployment, AWS DLM for automated backups, and AWS Secrets Manager integration.
- [x] **Serverless Observability:** Amazon CloudWatch alarms triggering SNS Topics, invoking an AWS Lambda (Python) to push real-time alerts to a Telegram Bot.

## ☁️ Phase 9: Cloud Mastery & DevSecOps ✅
*Focus: Shift-Left Security, Zero-Trust IAM, and Configuration Management.*
- [x] **Zero-Trust IAM:** Replaced GitHub static secrets with secure AWS OIDC (OpenID Connect) federation.
- [x] **Shift-Left Security:** Implemented Trufflehog (Secret Scanning) and Checkov (IaC compliance) directly into the integration workflow.
- [x] **Ansible Automation:** Idempotent playbooks to standardize AWS EC2 base setups and manage system packages securely.

## 🔄 Phase 9.5: Enterprise GitOps & Continuous Deployment ✅
*Focus: Full deployment automation across multiple environments with Zero Human Intervention.*
- [x] **CI Engine:** GitHub Actions handles code auditing, multi-stage linter validation (Hadolint, Actionlint), and builds immutable multi-architecture (ARM64) Docker images via QEMU emulators.
- [x] **Staging (Edge/Raspberry Pi):** Automated pipeline pushes the `:stage` tag to GHCR. An autonomous Watchtower pull-agent automatically updates the local Pi environment without exposing public ports.
- [x] **Production (Cloud/AWS):** Automated pipeline pushes the `:prod` tag directly to AWS ECR. Real GitOps Pull-CD via Watchtower running natively on AWS Graviton architecture with zero human terminal intervention.

## 🏰 Phase 10: Advanced Cloud Architecture & AWS Hardening ✅
*Focus: FinOps, Container Orchestration Pivot, and Enterprise Identity Governance.*
- [x] **Domain & Go-Live:** Production environment verified and fully operational on native ARM64 architecture.
- [x] **Self-Hosting Security:** Deploy Vaultwarden password manager using the secure GitOps pipeline, connecting it to an isolated, single-AZ RDS PostgreSQL database with strict Security Group boundaries.
- [x] **Container Orchestration Pivot:** Evaluated AWS ECS. *Strategic Decision:* Discarded vendor-locked ECS in favor of industry-standard Kubernetes (K3s) on bare-metal Edge hardware to maximize FinOps and deep technical learning.
- [x] **Deep Cloud AWS Hardening:**
  - Exhaustively differentiated Identity-based policies vs. Resource-based policies within IAM.
  - Enabled IAM Access Analyzer at the regional level to audit and trim inactive permissions.
  - Applied strict S3 Bucket Policies to deny unencrypted HTTP traffic (force `aws:SecureTransport`).
  - **FinOps:** Implemented S3 Lifecycle Policies to automatically transition logs to S3 Glacier Deep Archive after 30 days ($0.0009/GB), minimizing cloud waste.
- [x] **"Deploy & Destroy" Proof of Works (Capa 7):**
  - **AWS WAF (Web Application Firewall):** Deployed via Terraform, simulated an SQLi attack, intercepted malicious payloads via AWS Managed Rules, and executed `terraform destroy` for zero-cost auditing.
  - **ALB & Auto Scaling:** Provisioned an Application Load Balancer to route multi-AZ traffic, handled custom Health Check matchers (`200-499`) for strict reverse proxies, and tore down the infrastructure cleanly.

## 🛡️ Phase 10.5: Edge SOC & Defensive CyberSec (`mysstic-sentinel`) ✅
*Focus: Building a custom AI-driven SIEM (Sentinel) and deploying Local Edge Telemetry.*

**🧠 Part 1: MyssTic Sentinel (The Cloud-Native SIEM Backend)**
*A proprietary log ingestion and threat analysis engine built from scratch.*
- [x] **Backend Architecture:** Developed a decoupled Producer-Consumer REST API using Python (Django REST Framework).
- [x] **Asynchronous Processing:** Implemented Redis as a message broker and Celery as a distributed worker to handle high-volume log ingestion without blocking the main thread.
- [x] **Generative AI Integration:** Integrated Google Gemini LLM via Python SDK to evaluate raw logs, detect anomaly patterns, and determine threat severity dynamically.
- [x] **Secure Ingestion & SOAR:** Secured endpoints with JWT (JSON Web Tokens) and implemented automated Telegram webhooks for critical [CRÍTICO] alerts.

**👁️ Part 2: MyssTic Edge (Local SOC & Telemetry)**
*The first line of defense deployed on the Raspberry Pi 5 to collect, clean, and visualize logs.*
- [x] **Zero-Trust Routing & TLS:** Configured Pi-Hole as an Internal Route 53 and Caddy Proxy for internal TLS termination (`.lan` domains) over the Tailscale mesh.
- [x] **Data Pipeline & Search:** Deployed Promtail to extract raw logs from Docker sockets, routing them to Grafana Loki for fast-indexed log querying.
- [x] **Active Observability:** Deployed Uptime Kuma for internal/external health checks and Prometheus + Node Exporter for deep hardware telemetry.
- [x] **The Hybrid Brain Integration:** Connect MyssTic Edge (Promtail/Alertmanager) to send authenticated HTTP POST Webhooks (JWT) containing critical local logs to the MyssTic Sentinel API for AI evaluation.

---

## ☸️ Phase 11: The Final Boss (`mysstic-edge` Kubernetes)
*Focus: Industry-standard container orchestration on Bare-Metal.*
- [ ] **K3s on Edge:** Migrate the Raspberry Pi infrastructure from Docker Compose to a lightweight K3s cluster.
- [ ] **Kubernetes Abstractions:** Master the design of Pods, Deployments, Services, ConfigMaps, and Ingress controllers.
- [ ] **Helm Charts & Kustomize:** Dynamic packaging and versioning of microservices for declarative GitOps deployments.
- [ ] **Cluster Hardening:** Implement Network Policies for strict Pod isolation and RBAC (Role-Based Access Control).
- [ ] **Native K8s Observability:** Deploy Kube-Prometheus-Stack (kube-state-metrics) for deep cluster-level telemetry.
- [ ] **GitOps Realization:** Deploy ArgoCD for declarative, auditable, and purely code-based cluster state synchronization.

## 🚀 Extras & Pro League (Horizon)
*Focus: Data Engineering, Cloud Architecture, and Career Growth.*
- [ ] Introduction to Data Engineering, APM architectures (Elastic/ELK stack), and data workflow automation (n8n / Apache Airflow).
- [ ] Hack The Box (HTB) - Retired Machines in "blind" mode (no guides or writeups).
- [ ] AWS Solutions Architect Associate official preparation and certification.

---
*Roadmap updated automatically via CI/CD.*