---
title: "Engineering Projects"
draft: false
hideMeta: true
---

A showcase of my journey building resilient infrastructure and automating workflows.

## 🛡️ Ez-Lab: Hybrid Cloud Infrastructure
*Infrastructure as Code / DevSecOps / Self-Hosted*

Designed and deployed a resilient, self-hosted microservices architecture on a **Raspberry Pi 5**. This project simulates a production environment protected by Zero Trust principles.

* **The Challenge:** Securely exposing local services behind an ISP CGNAT without opening dangerous ports.
* **The Solution:** Implemented a **Cloudflare Tunnel** (Zero Trust) to bypass NAT restrictions while maintaining strict firewall rules.
* **Architecture:**
    * **Orchestration:** Docker Compose managing the *Arr Stack (Media Ops) and Pi-hole.
    * **Networking:** Tailscale Mesh VPN for secure remote management.
    * **Automation:** Python scripts auditing kernel sensors (Temp/RAM) and triggering Telegram alerts.
* **Status:** `Production 🟢`

> **Tech Stack:** `Docker` `Python` `Cloudflare Zero Trust` `Linux Hardening` `Bash`
> [🔗 View Repository](https://github.com/ezequieldlv/ez-lab)

---

## 🗺️ The SRE Roadmap (Engineering Log)
*Continuous Learning / Documentation*

I am publicly documenting my transition to Site Reliability Engineering through a structured **9-Phase Roadmap**.

* **Phase 1-3 (Completed):** Infrastructure Hardening, Containerization, and Security Fortress.
* **Phase 4 (Current):** "The Builder" - Learning Golang and Advanced Docker Builds.
* **Future Phases:** Deep Observability (Prometheus/Grafana), Kubernetes Orchestration, and AWS Cloud Migration.

> [👉 **View Full Roadmap & Progress**]({{< ref "roadmap.md" >}})

---

## 📄 Knowledge Base Modernization (Konecta)
*Docs-as-Code / Process Optimization*

Led an initiative to modernize internal technical documentation for Tier 2 Support agents.
* **Problem:** Critical troubleshooting info was scattered across legacy static sites and PDFs.
* **Solution:** Migrated content to a **Docs-as-Code** system using **MkDocs**, enabling version control and faster search.

> **Tech Stack:** `MkDocs` `Markdown` `Git`

---

## 🌐 SRE Portfolio (This Website)
*Edge Computing / Self-Hosted / CI/CD*

More than just a website, this portfolio is a live demo of my infrastructure skills. Hosted directly on my home cluster, exposed securely to the world.

* **Infrastructure:** Raspberry Pi 5 (8GB) running Arch Linux/Debian.
* **Web Server:** Nginx configured as a Reverse Proxy with strict caching policies.
* **Security:** Exposed via **Cloudflare Tunnels** (Zero Trust) to bypass CGNAT and enforce HTTPS without opening firewall ports.
* **CI/CD:** Automated deployment pipeline (Git push -> Live).

> **Tech Stack:** `Hugo` `Nginx` `Cloudflare` `Raspberry Pi`

---