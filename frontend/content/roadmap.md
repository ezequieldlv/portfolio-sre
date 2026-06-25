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

## 🚧 Phase 9: Cloud Mastery & DevSecOps (Current Focus)
*Focus: Shift-Left Security, Zero-Trust IAM, and Configuration Management.*
- [ ] **Zero-Trust IAM:** Replace GitHub static secrets with AWS OIDC (OpenID Connect).
- [ ] **Shift-Left Security:** Implement Trufflehog (Secret Scanning) and Checkov (IaC compliance) before deployment.
- [ ] **Ansible Automation:** Idempotent playbooks to standardize AWS EC2 setup securely.
- [ ] **Self-Hosting (Vaultwarden):** Deploy Vaultwarden password manager using the new secure pipeline, connecting to the isolated RDS.

## 🚧 Phase 9.5: Enterprise GitOps & Continuous Deployment
*Focus: Full deployment automation across multiple environments with Zero Human Intervention.*
- [ ] **CI Engine:** GitHub Actions compiles Go/Python, runs linters, and builds immutable Docker images.
- [ ] **Staging (Edge/Raspberry Pi):** Pipeline pushes `:stage` tag to GHCR. Watchtower (Pull agent) automatically updates the local Pi without exposing ports.
- [ ] **Production (Cloud/AWS):** Pipeline pushes `:latest` or version tags to AWS ECR. Secure Push-CD via Tailscale VPN updates the EC2 Compose stack.

## 🔒 Phase 10: Advanced Cloud Security & Edge SOC
*Focus: FinOps, Least Privilege, and transforming the Home Lab into a Security Operations Center.*
- [ ] **Deep Cloud AWS:**
  - Implement IAM Access Analyzer (Regional level).
  - Apply strict S3 Bucket Policies (force `aws:SecureTransport`).
  - FinOps: S3 Lifecycle Policies (Transition Traefik logs to Glacier Deep Archive).
  - "Deploy & Destroy" Proof of Works: Terraform WAF (simulate SQLi) and ALB Auto Scaling.
- [ ] **Hybrid SOC (Raspberry Pi):**
  - Uptime Kuma for Status Pages & 60s pings (Telegram alerts).
  - Monitor AWS from Edge: Grafana & Prometheus scraping AWS metrics via Tailscale.
  - Pi-Hole as Internal Route 53 (e.g., `grafana.lan`).
  - Red Team/SOC L1 Training: TryHackMe Cyber Defense and PortSwigger Web Security.

## 🐧 Phase 10.5: The Metal (Arch Linux From Scratch)
*Focus: Terminal Mastery and Low-Level OS tuning.*
- [ ] **Project:** Install Arch Linux from scratch on personal notebook.
- [ ] **Customization:** Master dotfiles, Hyprland (Wayland compositor).
- [ ] **Tryhard IDE:** Setup Neovim as the primary development environment.

## ☸️ Phase 11: The Final Boss (Kubernetes)
*Focus: Industry-standard container orchestration.*
- [ ] **K3s on Edge:** Migrate the Raspberry Pi from Docker Compose to a lightweight K3s cluster.
- [ ] **Kubernetes Abstractions:** Master Pods, Deployments, Services, and Ingress controllers.
- [ ] **Security:** Implement Network Policies (Pod isolation) and RBAC.
- [ ] **GitOps Realization:** Deploy ArgoCD for declarative, pull-based cluster synchronization.

## 🚀 Extras & Pro League (Horizon)
*Focus: Data Engineering, Cloud Architecture, and Career Growth.*
- [ ] Learn Data Engineering, APM (Elastic/ELK stack), and Automation (n8n/Airflow).
- [ ] Hack The Box (HTB) - Retired Machines (No guides).
- [ ] AWS Solutions Architect Associate preparation.

---
*Roadmap updated automatically via CI/CD.*