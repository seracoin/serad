#!/bin/bash
rm -rf /tmp/serad-temp

serad --devnet --appdir=/tmp/serad-temp --profile=6061 &
SERAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:22611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SERAD_PID

wait $SERAD_PID
SERAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "serad exit code: $SERAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SERAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
