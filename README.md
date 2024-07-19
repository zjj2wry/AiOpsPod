# AiOpsPod
A Kubernetes-based framework that uses Large Language Models (LLMs) for automated operational efficiency

## Description
AiOpsPod is a cutting-edge operations monitoring and automation system that combines the power of Kubernetes (K8s) and Large Language Models (LLMs) to streamline and enhance operational efficiency. The primary focus of AiOpsPod is to provide intelligent, automated responses to operational queries, seamless integration with monitoring systems, and efficient fault handling mechanisms. Here's a detailed look at the project's core functionalities and features:

## Key Features
- Continuous Learning: AiOpsPod leverages LLMs that continuously learn from Standard Operating Procedures (SOPs) stored in Feishu documentation. This enables the system to provide accurate and up-to-date responses to operations-related questions.
- Monitoring Integration: The system seamlessly integrates with internal monitoring tools such as Prometheus for data queries and Loki for log searches. This allows for real-time monitoring and log analysis directly within the AiOpsPod framework.
- Automated Alert Handling: When an alert is received from the Alert Manager, AiOpsPod processes the alert, consults relevant SOPs, and interacts with users through Feishu to provide runbooks. Upon user approval, AiOpsPod executes the necessary automation tasks, adhering to the principle of minimal privileges.
- Agent Architecture: AiOpsPod is designed with a modular architecture, featuring separate agents for handling logs, monitoring information, and automated commands. Each agent operates as an independent K8s service or job, enabling efficient service discovery and task delegation.

## Architecture Overview
The system architecture is designed to ensure robust and scalable operations:

```PLAINTEXT
        +-------------------------+                 +-------------------------+
        |      Feishu API         |                 |      Alert Manager      |
        +-------------------------+                 +-------------------------+
                        |                                              |
                        v                                              v
                        +----------------------------------------------+
                                          |
                                          v
                          +--------------------------------+
                          |          LLM Service           |                    
                          |                                |             +-------------------+
                          |  - OPS QA                      |             | Vector Database   |
                          |  - SOP Update                  |<----------->| (SOP Storage)     |
                          |  - Prometheus/loki Integration |             +-------------------+
                          |  - Runbook Execution           |
                          |  - Extension Agent/Tool        |
                          +--------------------------------+
                                      |
   +----------------------+-----------+------------+------------------------+
   |                      |                        |                        |
   v                      v                        v                        v
+-----------+      +--------------+     +-----------------+      +------------------------+
|Prometheus |      | Loki Service |     | Extension       |      |  Feishu Notification   |
| Service   |      | (Log Query)  |     | Agent/Tool      |      |  Runbook  Execution    |
+-----------+      +--------------+     +-----------------+      +------------------------+
```

## Run locally
Start weaviate vector db
> If you cann't pull the image in China, add `"registry-mirrors": ["https://dockerhub.icu"]` to `~/.docker/daemon.json` and restart your docker daemon.
```bash
docker run -d -p 8080:8080 -p 50051:50051 semitechnologies/weaviate:1.25.8
```
