#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

# Save the original config so we can restore it after the test.
cp main.tf main.tf.bak
trap 'cp main.tf.bak main.tf && rm -f main.tf.bak' EXIT

echo "=== terraform init ==="
terraform init -input=false

echo "=== Step 1: create (expect 'hello, world!') ==="
terraform apply -input=false -auto-approve 2>&1 | tee /tmp/apply1.log
if ! grep -q "hello, world!" /tmp/apply1.log; then
  echo "::error::expected 'hello, world!' in create output"
  exit 1
fi

echo "=== Step 2: change name to 'earth' (expect 'earth changed!') ==="
sed -i 's/name = "world"/name = "earth"/' main.tf
terraform apply -input=false -auto-approve 2>&1 | tee /tmp/apply2.log
if ! grep -q "earth changed!" /tmp/apply2.log; then
  echo "::error::expected 'earth changed!' in update output"
  exit 1
fi

echo "=== Step 3: destroy (expect 'bye, earth!') ==="
terraform destroy -input=false -auto-approve 2>&1 | tee /tmp/destroy.log
if ! grep -q "bye, earth!" /tmp/destroy.log; then
  echo "::error::expected 'bye, earth!' in destroy output"
  exit 1
fi

echo "ALL TESTS PASSED"
