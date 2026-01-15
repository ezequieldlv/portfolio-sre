---
title: "Projects"
draft: false
hideMeta: true
---

Here is a collection of my work in **Infrastructure**, **Software Development**, and **Automation**.

## 🛡️ Ez-Lab: Zero Trust Private Cloud
*Infrastructure as Code / DevSecOps*

Designed and deployed a resilient, self-hosted infrastructure on a **Raspberry Pi 5**, simulating a production environment protected by a Zero Trust network.
* **The Challenge:** Accessing local services securely behind an ISP CGNAT without exposing public ports.
* **The Solution:** Implemented a **Tailscale Mesh VPN** as an overlay network, bypassing NAT restrictions while maintaining strict firewall rules (UFW).
* **Architecture:** Docker Compose microservices (Media Server, NAS, VPN) with persistent storage and automated recovery policies.
* **Key Achievement:** Solved Docker kernel routing issues to enable "Exit Node" functionality, routing traffic through the home network from anywhere.

> **Tech Stack:** `Docker Compose` `Tailscale` `Linux Hardening` `Jellyfin` `Samba` `Bash`
> [🔗 View Repository](https://github.com/ezequieldlv/ez-lab)

<div class="gallery-container">
    <div class="gallery-item">
        <img src="/portfolio-sre/images/portainer.png" alt="Ez-Lab Architecture">
        <span class="caption">Fig 1. Zero Trust Architecture Diagram</span>
    </div>
</div>

## 📄 Internal Knowledge Base Modernization
*Docs-as-Code / Automation (at Konecta)*

Led the migration of technical documentation from legacy Google Sites to a modern workflow.
Implemented Google Analytics to track internal search patterns, allowing us to identify missing documentation and reduce ticket resolution time.
* **Impact:** Improved information retrieval speed for Tier 2 agents.
* **Result:** Standardized troubleshooting guides facilitating collaborative maintenance.
**The Challenge:** Information was scattered across legacy sites.
**The Solution:** Migrated to a Docs-as-Code workflow.

<div class="gallery-container">
    <div class="gallery-item">
        <img src="/portfolio-sre/images/legacy_google_site.png" alt="Legacy System">
        <span class="caption">Fig 1. Legacy System (Old)</span>
    </div>
    <div class="gallery-item">
        <img src="/portfolio-sre/images/modern_mkdocs.png" alt="MkDocs System">
        <span class="caption">Fig 2. New MkDocs Workflow</span>
    </div>
</div>

> **Tech Stack:** `MkDocs` `Python` `GitLab` `Markdown` `Google Analytics`

## 📦 Logistics Microservices Backend
*Java / Microservices (University Project)*

Co-developed a backend solution for a terrestrial logistics system (shipping containers).
* **Architecture:** Microservices-based architecture ensuring modularity.
* **Tech:** **Java Spring Boot**, **Docker**, Keycloak (Identity Management), RESTful APIs.

## 🌐 SRE Portfolio (This Website)
*Web Performance / CI/CD*

A minimalist, high-performance portfolio built with a **Docs-as-Code** philosophy.
* **Engine:** Hugo (Static Site Generator) + PaperMod Theme.
* **Deployment:** Dockerized container served via Nginx (Upcoming).

## 🧬 Biohacking Data Analysis (Planned)
*Python / Data Observability*

Scripts to correlate physiological data (sleep, nutrition) with cognitive performance, applying observability principles to human biology.