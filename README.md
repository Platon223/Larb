<p align="center">
  <img style="width: 120px; margin: 0 auto 40px; display: block; border-radius: 14px;" src="https://i.imgur.com/Gu7EcLR.jpeg" alt="Logo" referrerpolicy="no-referrer"/>
</p>

# Larb: LogArbor CLI

Larb is a CLI tool that is used in order to interact with <a href="https://github.com/Platon223/LogArbor">LogArbor</a>. Larb allows developers to observe their applications' logs, alerts, and metrics in the terminal. Larb also has an extra feature that doesn't exist on the LogArbor platform itself: Logby, it is an AI chat assistant that helps developers get started with LogArbor and Larb. This tool is built for developers who like to do everything in their terminal even if it is Log Managment.

---

# ⚙️ Setup

Follow these steps to get started with Larb

---

## 💻 Installation

### Mac
```bash
brew install --cask Platon223/tap/larb
```

### Linux
```bash
curl -sSL https://github.com/Platon223/Larb/releases/latest/download/larb_linux_amd64.tar.gz | tar -xz
sudo mv larb /usr/local/bin/
```

**For ARM Linux:**
```bash
curl -sSL https://github.com/Platon223/Larb/releases/latest/download/larb_linux_arm64.tar.gz | tar -xz
sudo mv larb /usr/local/bin/
```

### Windows
Download the `.tar.gz` from the [GitHub releases page](https://github.com/Platon223/Larb/releases/latest), then extract and add `larb.exe` to your PATH manually.

**Using PowerShell:**
```powershell
Invoke-WebRequest -Uri "https://github.com/Platon223/Larb/releases/latest/download/larb_windows_amd64.tar.gz" -OutFile "larb.tar.gz"
tar -xf larb.tar.gz
Move-Item larb.exe C:\Windows\System32\larb.exe
```

---

## 📖 Basic Usage

Once installed, you can start using Larb with the following commands:

```bash
# Initialize Larb configuration
larb init <your_api_key>

# View logs for a specific service
larb get --id <service-id>

# Tail logs in real-time
larb tail logs <service-id>

# Chat with Logby (AI Assistant)
larb chat

# List all services
larb all
```

For more commands, please visit our <a href="https://logarbor.com/docs">docs</a>


--- 


## Lisence

Copyright © 2026 Platon Tikhnenko. All rights reserved.

This project is proprietary. The source code is made public for the sole purpose 
of portfolio review by potential employers. It may not be copied, redistributed, 
or used for any other purpose without explicit written permission.


