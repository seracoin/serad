#!/bin/bash

APPDIR=/tmp/serad-temp
SERAD_RPC_PORT=29587

rm -rf "${APPDIR}"

serad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${SERAD_RPC_PORT}" --profile=6061 &
SERAD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${SERAD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $SERAD_PID

wait $SERAD_PID
SERAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "serad exit code: $SERAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SERAD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
