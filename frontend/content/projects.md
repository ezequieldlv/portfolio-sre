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

<br>
<div style="text-align: center; margin-top: 20px;">
  <p style="font-weight: bold; color: #5cb85c; margin-bottom: 5px;">📸 Live Infrastructure Status (Portainer)</p>
  <img src="/images/portainer_dashboard.png" alt="Portainer Dashboard Production" style="width: 100%; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); border: 1px solid #ddd;">
</div>
<br>

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

<br>
<div style="display: flex; gap: 20px; flex-wrap: wrap; justify-content: center; margin-top: 20px;">
  
  <div style="flex: 1; min-width: 300px; text-align: center;">
    <p style="font-weight: bold; color: #d9534f; margin-bottom: 5px;">❌ Before: Legacy Static Site</p>
    <img src="/images/legacy_google_site.png" alt="Legacy Documentation" style="width: 100%; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); border: 1px solid #ddd;">
  </div>

  <div style="flex: 1; min-width: 300px; text-align: center;">
    <p style="font-weight: bold; color: #5cb85c; margin-bottom: 5px;">✅ After: Docs-as-Code (MkDocs)</p>
    <img src="/images/modern_mkdocs.png" alt="Modern MkDocs System" style="width: 100%; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); border: 1px solid #ddd;">
  </div>

</div>
<br>

---

## 🌐 SRE Portfolio & Live Telemetry
*Edge Computing / Microservices / CI/CD*

More than just a website, this portfolio is a live microservices demo. The frontend (static) is currently fetching real-time telemetry from a **Golang API** running on a separate Docker container within my Raspberry Pi cluster, exposed via Zero Trust.

* **Frontend:** Hugo + Nginx (Static Edge)
* **Backend:** Golang 1.22 REST API (Distroless Container)
* **Security:** Cloudflare Tunnels (Zero Trust) + Strict CORS Policies.

<div style="background-color: #0d1117; color: #58a6ff; padding: 20px; border-radius: 8px; font-family: 'Courier New', Courier, monospace; border: 1px solid #30363d; margin: 20px 0;">
    <div style="display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #30363d; padding-bottom: 10px; margin-bottom: 15px;">
        <span style="font-weight: bold; color: #c9d1d9;">📡 RPi-5 Cluster Telemetry</span>
        <span id="api-status-indicator" style="color: #f85149;">🔴 Offline</span>
    </div>
    <ul id="telemetry-data" style="list-style: none; padding: 0; margin: 0; line-height: 1.8; color: #c9d1d9; font-size: 0.9em;">
        <li>> Initializing handshake with api.ez-lab.site... ⏳</li>
    </ul>
</div>

<script>
    async function fetchTelemetry() {
        try {
            const response = await fetch('https://api.ez-lab.site/api/health');
            if (!response.ok) throw new Error('Network response was not ok');
            
            const data = await response.json();
            
            document.getElementById('api-status-indicator').innerHTML = '🟢 Live';
            document.getElementById('api-status-indicator').style.color = '#3fb950';
            
            const container = document.getElementById('telemetry-data');
            container.innerHTML = `
                <li><span style="color: #79c0ff;">$</span> <strong>OS/Arch:</strong> ${data.os}_${data.architecture}</li>
                <li><span style="color: #79c0ff;">$</span> <strong>Engine:</strong> ${data.go_version}</li>
                <li><span style="color: #79c0ff;">$</span> <strong>Server Time:</strong> ${data.server_time}</li>
                <li><span style="color: #79c0ff;">$</span> <strong>Response:</strong> <em>"${data.message}"</em></li>
            `;
        } catch (error) {
            document.getElementById('telemetry-data').innerHTML = 
                '<li style="color: #f85149;">> [ERR] API connection refused. Backend might be down or blocked by CORS.</li>';
            console.error("Telemetry Error:", error);
        }
    }
    document.addEventListener("DOMContentLoaded", fetchTelemetry);
</script>

> **Tech Stack:** `Golang` `Hugo` `Nginx` `Cloudflare` `Docker Networks`

---