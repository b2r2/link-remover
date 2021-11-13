#!/bin/bash
set -euo pipefail
if [ -f /run/secrets/TOKEN ]; then
   export TOKEN=$(cat /run/secrets/TOKEN)
fi

echo "Secret is: $TOKEN"