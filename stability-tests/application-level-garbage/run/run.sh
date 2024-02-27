#!/bin/bash
rm -rf /tmp/serad-temp

serad --devnet --appdir=/tmp/serad-temp --profile=6061 --loglevel=debug &
SERAD_PID=$!
SERAD_KILLED=0
function killSeradIfNotKilled() {
    if [ $SERAD_KILLED -eq 0 ]; then
      kill $SERAD_PID
    fi
}
trap "killSeradIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:22611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SERAD_PID

wait $SERAD_PID
SERAD_KILLED=1
SERAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "serad exit code: $SERAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SERAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
