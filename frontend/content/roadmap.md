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
- [x] **Dominio y Sitio Web Operativo:** Go-Live verificado en producción bajo arquitectura nativa ARM64.
- [ ] **Self-Hosting Security:** Deploy Vaultwarden password manager using the secure GitOps pipeline, connecting it to an isolated, multi-AZ RDS PostgreSQL database with strict Security Group boundaries.
- [ ] **Deep Cloud AWS Hardening:**
  - Diferenciar exhaustivamente Identity-based policies vs Resource-based policies.
  - Activar IAM Access Analyzer a nivel regional para auditar y recortar permisos inactivos.
  - Aplicar S3 Bucket Policies estrictas que denieguen tráfico HTTP no cifrado (forzar `aws:SecureTransport`).
  - FinOps: Implementar S3 Lifecycle Policies para transicionar automáticamente logs de Traefik a S3 Glacier Deep Archive a los 30 días ($0.0009/GB).
- [ ] **Experimentos "Deploy & Destroy" (Proof of Work):**
  - **AWS WAF (Web Application Firewall):** Desplegar mediante Terraform, simular e inyectar un ataque SQLi de prueba, documentar el bloqueo para el portfolio/LinkedIn y ejecutar `terraform destroy` para mitigar costos.
  - **ALB & Auto Scaling:** Levantar balanceador de aplicación, emular tráfico masivo, documentar el escalado automático horizontal de instancias y destruir.
  - **Event-Driven Security:** Diseñar funciones Lambda avanzadas que se disparen por eventos críticos de CloudWatch o escrituras no autorizadas en S3.
- [ ] **Hybrid SOC (Raspberry Pi):**
  - Desplegar Uptime Kuma como vigilante del sistema (Status Pages públicos, pings cada 60s y alertas críticas integradas a Telegram).
  - Configurar Grafana, Prometheus y Node Exporter para monitorización profunda de recursos (CPU, RAM, Red), escrapeando las métricas de AWS desde el entorno local de forma segura a través del túnel de Tailscale VPN.
  - Configurar Pi-Hole como Route 53 interno para la resolución forzada de dominios locales (ej. `grafana.lan`, `vaultwarden.lan`).
  - Instalar Nginx o Caddy como Reverse Proxy local para gestionar certificados SSL internos y terminación TLS en el Edge.
- [ ] **Red Team & SOC Level 1 Training:**
  - Estudiar rutas defensivas en TryHackMe (SOC L1, Cyber Defense).
  - Practicar remediación de vulnerabilidades web (XSS, CSRF, Inyecciones SQL) usando la academia gratuita de PortSwigger.
  - Realizar escaneos de puertos y pentesting a nivel red local para asegurar el aislamiento de servicios domésticos (Samba, Jellyfin).

---

## 🐧 Phase 10.5: The Metal (Arch Linux From Scratch)
*Focus: Terminal Mastery and Low-Level OS tuning.*
- [ ] **Project:** Instalar Arch Linux desde cero en la notebook personal mediante CLI puro.
- [ ] **Customization:** Diseñar y versionar dotfiles propios, configurando Hyprland como compositor de ventanas bajo Wayland.
- [ ] **Tryhard IDE:** Configurar Neovim desde cero como el entorno de desarrollo y edición principal.

## ☸️ Phase 11: The Final Boss (Kubernetes)
*Focus: Industry-standard container orchestration.*
- [ ] **K3s on Edge:** Migrar la infraestructura de la Raspberry Pi desde Docker Compose hacia un clúster ligero de K3s.
- [ ] **Kubernetes Abstractions:** Dominar el diseño de Pods, Deployments, Services, ConfigMaps y controladores de Ingress.
- [ ] **Cluster Hardening:** Implementar Network Policies para aislamiento estricto de Pods y control de accesos mediante RBAC (Role-Based Access Control).
- [ ] **GitOps Realization:** Desplegar ArgoCD para la sincronización declarativa, auditable y puramente basada en código del estado del clúster.

## 🚀 Extras & Pro League (Horizon)
*Focus: Data Engineering, Cloud Architecture, and Career Growth.*
- [ ] Introducción a Data Engineering, arquitecturas APM (Elastic/ELK stack) y motores de automatización de flujos de datos (n8n / Apache Airflow).
- [ ] Resolver máquinas retiradas en Hack The Box (HTB) en modo "blind" (sin guías ni writeups).
- [ ] Preparación y certificación oficial de AWS Solutions Architect Associate.

---
*Roadmap updated automatically via CI/CD.*