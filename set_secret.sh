#!/bin/bash
set -euo pipefail
if [ -f /run/secrets/thepassword ]; then
   export TOKEN=$(cat /run/secrets/TOKEN)
fi