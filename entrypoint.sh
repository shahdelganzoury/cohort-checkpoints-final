#!/bin/sh
set -e

EXERCISE=$(echo "$EXERCISE" | sed 's/\([a-z]\)\([A-Z]\)/\1-\2/g' | tr 'A-Z' 'a-z')

if [ -z "$EXERCISE" ]; then
  echo "EXERCISE env var is not set"
  exit 1
fi

TEST_DIR="/app/tests/$EXERCISE"

if [ ! -d "$TEST_DIR" ]; then
  echo "No tests found for: $EXERCISE"
  exit 1
fi

WORKDIR="/jail/work"
mkdir -p "$WORKDIR"
cp -r /jail/student/* "$WORKDIR"/

cd "$WORKDIR"
go mod init student 2>/dev/null || true

if [ "${CODE_EDITOR_RUN_ONLY:-}" = "true" ]; then
  go run .
  exit 0
fi

# ============================================================
# Forbidden check
# ============================================================

FORBIDDEN_FILE="$TEST_DIR/forbidden.txt"

if [ -f "$FORBIDDEN_FILE" ]; then
  STUDENT_FILES=$(find . -name "*.go" ! -name "*_test.go")

  while IFS= read -r pattern || [ -n "$pattern" ]; do
    case "$pattern" in
      ""|\#*) continue ;;
    esac

    for file in $STUDENT_FILES; do
      if grep -q "$pattern" "$file"; then
        echo "Forbidden usage detected: $pattern"
        exit 1
      fi
    done
  done < "$FORBIDDEN_FILE"
fi

# ============================================================
# Copy tests
# ============================================================

cp -r "$TEST_DIR/." "$WORKDIR"/

# ============================================================
# Visible tests
# ============================================================

set +e
VISIBLE_OUTPUT=$(go test -v -count=1 -timeout 10s ./... 2>&1)
VISIBLE_EXIT=$?
set -e

if [ "$VISIBLE_EXIT" -ne 0 ]; then
  echo "$VISIBLE_OUTPUT"
  exit 1
fi

# ============================================================
# Hidden tests
# ============================================================

set +e
HIDDEN_OUTPUT=$(go test -tags hidden -v -count=1 -timeout 30s ./... 2>&1)
ALL_EXIT=$?
set -e

if [ "$ALL_EXIT" -ne 0 ]; then
  echo "$HIDDEN_OUTPUT"
  exit "$ALL_EXIT"
fi

echo "All tests passed"
exit 0