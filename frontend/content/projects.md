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
  <img src="../images/portainer_dashboard.png" alt="Portainer Dashboard Production" style="width: 100%; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); border: 1px solid #ddd;">
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
    <img src="../images/legacy_google_site.png" alt="Legacy Documentation" style="width: 100%; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); border: 1px solid #ddd;">
  </div>

  <div style="flex: 1; min-width: 300px; text-align: center;">
    <p style="font-weight: bold; color: #5cb85c; margin-bottom: 5px;">✅ After: Docs-as-Code (MkDocs)</p>
    <img src="../images/modern_mkdocs.png" alt="Modern MkDocs System" style="width: 100%; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); border: 1px solid #ddd;">
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

### 🏗️ Live Architecture

<style>
  pre.mermaid {
    display: flex;
    justify-content: center;
    width: 100%;
    margin: 20px 0;
  }
</style>

<pre class="mermaid">
graph TD
    %% Definición de colores (Estilo Tokyo Night + Cloudflare)
    classDef user fill:#16161e,stroke:#7aa2f7,stroke-width:2px,color:#c0caf5
    classDef cf fill:#f6821f,stroke:#fff,stroke-width:2px,color:#fff,font-weight:bold
    classDef daemon fill:#1a1b26,stroke:#7aa2f7,stroke-width:2px,color:#c0caf5
    classDef app fill:#292e42,stroke:#9ece6a,stroke-width:2px,color:#c0caf5
    classDef hardware fill:#16161e,stroke:#f7768e,stroke-width:2px,color:#c0caf5

    User((🌐 Internet)):::user -->|HTTPS / SSL| Cloudflare{☁️ Cloudflare Edge}:::cf
    Cloudflare -->|Zero Trust Tunnel| Cloudflared[🛡️ Cloudflared Daemon]:::daemon
    
    subgraph "Ez-Lab (Raspberry Pi 5)"
        Cloudflared -->|web-net| Frontend["🖥️ Portfolio (Hugo)"]:::app
        Cloudflared -->|web-net| Backend["⚙️ API Telemetry (Go)"]:::app
        
        Frontend -.->|Internal Fetch| Backend
        
        Auditor["🐍 Python Auditor"]:::hardware -->|Hardware Metrics| OS[("🌡️ Kernel/Sensors")]:::hardware
    end
</pre>

<script type="module">
  import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
  mermaid.initialize({ startOnLoad: true, theme: 'dark' });
</script>

<style>
.terminal-box {
    background-color: #1a1b26;
    border-radius: 8px;
    box-shadow: 0 10px 15px -3px rgba(0,0,0,0.4);
    overflow: hidden;
    margin: 25px 0;
    border: 1px solid #292e42;
}
.terminal-header {
    background-color: #16161e;
    padding: 10px 15px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #292e42;
    font-family: system-ui, -apple-system, sans-serif;
    font-size: 0.85rem;
}
.terminal-title { color: #a9b1d6; font-weight: 600; letter-spacing: 0.5px; }
.status-indicator { display: flex; align-items: center; gap: 8px; font-weight: bold; }
.status-dot { width: 8px; height: 8px; border-radius: 50%; }
.terminal-body {
    padding: 20px;
    font-family: 'Fira Code', 'Courier New', monospace;
    color: #c0caf5;
    font-size: 0.95rem;
    line-height: 1.8;
}
.prompt { color: #7aa2f7; font-weight: bold; margin-right: 10px; user-select: none; }
.val { color: #9ece6a; }
.err { color: #f7768e; }
</style>

<div class="terminal-box">
    <div class="terminal-header">
        <span class="terminal-title">📡 root@ez-lab:~# ./telemetry.sh</span>
        <span id="api-status" class="status-indicator" style="color: #565f89;">
            <div id="api-dot" class="status-dot" style="background-color: #565f89;"></div> ⏳ Wait
        </span>
    </div>
    <div class="terminal-body">
        <ul id="telemetry-data" style="list-style: none; padding: 0; margin: 0;">
            <li><span class="prompt">$</span> Establishing Zero-Trust handshake...</li>
        </ul>
    </div>
</div>

<script>
    async function fetchTelemetry() {
        try {
            const response = await fetch('https://api.ez-lab.site/api/health');
            if (!response.ok) throw new Error('Network error');
            const data = await response.json();
            
            // Actualizar el header a verde brillante
            document.getElementById('api-status').innerHTML = '<div class="status-dot" style="background-color: #9ece6a; box-shadow: 0 0 10px #9ece6a;"></div> Live';
            document.getElementById('api-status').style.color = '#9ece6a';
            
            // Inyectar datos con formato
            const container = document.getElementById('telemetry-data');
            container.innerHTML = `
                <li><span class="prompt">$</span> OS_ARCH  <span class="val">=></span> ${data.os}_${data.architecture}</li>
                <li><span class="prompt">$</span> ENGINE   <span class="val">=></span> ${data.go_version}</li>
                <li><span class="prompt">$</span> SYSTIME  <span class="val">=></span> ${data.server_time}</li>
                <li><span class="prompt">$</span> RESPONSE <span class="val">=></span> "${data.message}"</li>
            `;
        } catch (error) {
            document.getElementById('api-status').innerHTML = '<div class="status-dot" style="background-color: #f7768e; box-shadow: 0 0 10px #f7768e;"></div> Offline';
            document.getElementById('api-status').style.color = '#f7768e';
            document.getElementById('telemetry-data').innerHTML = `<li><span class="prompt">$</span> <span class="err">FATAL: Connection refused by edge gateway.</span></li>`;
        }
    }
    document.addEventListener("DOMContentLoaded", fetchTelemetry);
</script>

> **Tech Stack:** `Golang` `Hugo` `Nginx` `Cloudflare` `Docker Networks`

---