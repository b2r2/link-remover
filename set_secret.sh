#!/bin/bash
set -euo pipefail
if [ -f /run/secrets/TOKEN ]; then
   # shellcheck disable=SC2155
   export TOKEN=$(cat /run/secrets/TOKEN)
fi

echo "$TOKEN"