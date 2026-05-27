#!/usr/bin/env bash
# Generate a self-signed cert that expires soon, so cert_expiry has something
# to flag in the security-audit scenario. Default validity: 7 days.
set -euo pipefail

cd "$(dirname "$0")/.."

days="${1:-7}"
host="${2:-admin.localhost}"

mkdir -p certs

openssl req -x509 -newkey rsa:2048 -nodes \
  -keyout certs/demo.key \
  -out certs/demo.crt \
  -days "$days" \
  -subj "/CN=$host" \
  -addext "subjectAltName=DNS:$host"

echo "Wrote certs/demo.crt (CN=$host) valid for $days days."
