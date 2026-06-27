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

## ✅ Phase 9: Cloud Mastery & DevSecOps (Completed)
*Focus: Shift-Left Security, Zero-Trust IAM, and Configuration Management.*
- [x] **Zero-Trust IAM:** Replaced GitHub static secrets with secure AWS OIDC (OpenID Connect) federation.
- [x] **Shift-Left Security:** Implemented Trufflehog (Secret Scanning) and Checkov (IaC compliance) directly into the integration workflow.
- [x] **Ansible Automation:** Idempotent playbooks to standardize AWS EC2 base setups and manage system packages securely.

## ✅ Phase 9.5: Enterprise GitOps & Continuous Deployment (Completed)
*Focus: Full deployment automation across multiple environments with Zero Human Intervention.*
- [x] **CI Engine:** GitHub Actions handles code auditing, multi-stage linter validation (Hadolint, Actionlint), and builds immutable multi-architecture (ARM64) Docker images via QEMU emulators.
- [x] **Staging (Edge/Raspberry Pi):** Automated pipeline pushes the `:stage` tag to GHCR. An autonomous Watchtower pull-agent automatically updates the local Pi environment without exposing public ports.
- [x] **Production (Cloud/AWS):** Automated pipeline pushes the `:prod` tag directly to AWS ECR. Real GitOps Pull-CD via Watchtower running natively on AWS Graviton architecture with zero human terminal intervention.

---

## 🔒 Phase 10: Advanced Cloud Security & Edge SOC (Next Immediate Focus)
*Focus: FinOps, Least Privilege, and transforming the Home Lab into a Security Operations Center.*
- [x] **Domain & Go-Live:** Live production deployment verified on native ARM64 architecture.
- [ ] **Self-Hosting Security:** Deploy Vaultwarden password manager using the secure GitOps pipeline, connecting it to an isolated, multi-AZ RDS PostgreSQL database with strict Security Group boundaries.
- [ ] **Deep Cloud AWS Hardening:**
  - Exhaustively differentiate Identity-based policies vs. Resource-based policies.
  - Enable IAM Access Analyzer at the regional level to audit and trim inactive permissions.
  - Apply strict S3 Bucket Policies to deny unencrypted HTTP traffic (force `aws:SecureTransport`).
  - FinOps: Implement S3 Lifecycle Policies to automatically transition Traefik logs to S3 Glacier Deep Archive after 30 days ($0.0009/GB).
- [ ] **"Deploy & Destroy" Experiments (Proof of Work):**
  - **AWS WAF (Web Application Firewall):** Deploy via Terraform, simulate an SQLi attack, document the mitigation for the portfolio/LinkedIn, and execute `terraform destroy` to optimize costs.
  - **ALB & Auto Scaling:** Provision an Application Load Balancer, emulate massive traffic, document horizontal instance auto-scaling, and tear down.
  - **Event-Driven Security:** Design advanced Lambda functions triggered by critical CloudWatch events or unauthorized S3 writes.
- [ ] **Hybrid SOC (Raspberry Pi):**
  - Deploy Uptime Kuma for system monitoring (Public Status Pages, 60s pings, and critical Telegram alerts).
  - Configure Grafana, Prometheus, and Node Exporter for deep resource monitoring (CPU, RAM, Network), securely scraping AWS metrics from the local environment via the Tailscale VPN tunnel.
  - Configure Pi-Hole as an Internal Route 53 for forced local domain resolution (e.g., `grafana.lan`, `vaultwarden.lan`).
  - Install Nginx or Caddy as a local Reverse Proxy to manage internal SSL certificates and TLS termination at the Edge.
- [ ] **Red Team & SOC Level 1 Training:**
  - Study defensive pathways on TryHackMe (SOC L1, Cyber Defense).
  - Practice web vulnerability remediation (XSS, CSRF, SQL Injections) using the PortSwigger Web Security Academy.
  - Perform local network port scanning and pentesting to ensure the isolation of home services (Samba, Jellyfin).

---

## ☸️ Phase 11: The Final Boss (Kubernetes)
*Focus: Industry-standard container orchestration.*
- [ ] **K3s on Edge:** Migrate the Raspberry Pi infrastructure from Docker Compose to a lightweight K3s cluster.
- [ ] **Kubernetes Abstractions:** Master the design of Pods, Deployments, Services, ConfigMaps, and Ingress controllers.
- [ ] **Cluster Hardening:** Implement Network Policies for strict Pod isolation and RBAC (Role-Based Access Control).
- [ ] **GitOps Realization:** Deploy ArgoCD for declarative, auditable, and purely code-based cluster state synchronization.

## 🚀 Extras & Pro League (Horizon)
*Focus: Data Engineering, Cloud Architecture, and Career Growth.*
- [ ] Introduction to Data Engineering, APM architectures (Elastic/ELK stack), and data workflow automation (n8n / Apache Airflow).
- [ ] Hack The Box (HTB) - Retired Machines in "blind" mode (no guides or writeups).
- [ ] AWS Solutions Architect Associate official preparation and certification.

---
*Roadmap updated automatically via CI/CD.*