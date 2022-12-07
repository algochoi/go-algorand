#!/bin/bash

# The test confirms algod v1 REST API returns the expected post-sunset status code (410).

filename=$(basename "$0")
scriptname="${filename%.*}"
date "+${scriptname} start %Y%m%d_%H%M%S"

my_dir="$(dirname "$0")"
source "$my_dir/rest.sh" "$@"

# Function is inspired by prior tests defining similar functions.
function rest() {
    local endpoint=$1
    local method=$2
    curl --include -X "$method" -q -s -H "Authorization: Bearer $PUB_TOKEN" "$NET$endpoint"
}

set -e
set -x
set -o pipefail

export SHELLOPTS

declare -a gets=(
  "/v1/status"
  "/v1/status/wait-for-block-after/1"
  "/v1/account/1"
  "/v1/account/1/transaction/1"
  "/v1/transactions/params"
  "/v1/account/1/transactions"
  "/v1/block/1"
  "/v1/ledger/supply"
  "/v1/transactions/pending"
  "/v1/transactions/pending/1"
  "/v1/account/1/transactions/pending"
  "/v1/asset/1"
  "/v1/assets"
  "/v1/transaction/1"
)

declare -a posts=(
  "/v1/transactions"
)

for endpoint in "${gets[@]}"; do
  response=$(rest "$endpoint" 'GET')
  if [ "$(echo "$response" | grep -c 'HTTP/1.1 410 Gone')" -ne 1 ]; then
      date "+${scriptname} status code != 410 for endpoint = $endpoint with response = $response"
      exit 1
    fi
done

for endpoint in "${posts[@]}"; do
 response=$(rest "$endpoint" 'POST')
  if [ "$(echo "$response" | grep -c 'HTTP/1.1 410 Gone')" -ne 1 ]; then
      date "+${scriptname} status code != 410 for endpoint = $endpoint with response = $response"
      exit 1
    fi
done

date "+${scriptname} OK %Y%m%d_%H%M%S"
