---
schema_version: 2
kind: crd
name: AIService
id: crd.aiservice
source: hub
traefik_version: v3.20.2
extracted_from:
  - pkg/apis/hub/v1alpha1/crd/hub.traefik.io_aiservices.yaml
summary: AIService is a Kubernetes-like Service to interact with a text-based LLM provider. It defines the parameters and credentials required to interact with various LLM providers.
fields:
  - name: anthropic
    type: object
  - name: azureOpenai
    type: object
  - name: bedrock
    type: object
  - name: cohere
    type: object
  - name: deepSeek
    type: object
  - name: gemini
    type: object
  - name: mistral
    type: object
  - name: ollama
    type: object
  - name: openai
    type: object
  - name: qWen
    type: object
representations:
  yaml_path: spec
  crd:
    apiVersion: hub.traefik.io/v1alpha1
    kind: AIService
    spec_path: .spec
---

# AIService

AIService is a Kubernetes-like Service to interact with a text-based LLM provider. It defines the parameters and credentials required to interact with various LLM providers.
