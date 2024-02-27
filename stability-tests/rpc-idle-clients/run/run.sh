#!/bin/bash
rm -rf /tmp/serad-temp

NUM_CLIENTS=128
serad --devnet --appdir=/tmp/serad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
SERAD_PID=$!
SERAD_KILLED=0
function killSeradIfNotKilled() {
  if [ $SERAD_KILLED -eq 0 ]; then
    kill $SERAD_PID
  fi
}
trap "killSeradIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $SERAD_PID

wait $SERAD_PID
SERAD_EXIT_CODE=$?
SERAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "serad exit code: $SERAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SERAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
