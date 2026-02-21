#!/usr/bin/env bash
set -euo pipefail

VERSION="${1:-dev}"
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
DIST_DIR="$ROOT_DIR/dist"
TEMPLATE="$ROOT_DIR/mcpb/manifest.json.tmpl"

if [ ! -f "$TEMPLATE" ]; then
  echo "Error: manifest.json.tmpl not found at $TEMPLATE" >&2
  exit 1
fi

if [ ! -d "$DIST_DIR" ]; then
  echo "Error: dist/ directory not found. Run goreleaser first." >&2
  exit 1
fi

PLATFORMS="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64"

count=0

for combo in $PLATFORMS; do
  os="${combo%/*}"
  arch="${combo#*/}"

  # goreleaser uses the format: borsdata-mcp_<os>_<arch>/borsdata-mcp[.exe]
  bin_name="borsdata-mcp"
  if [ "$os" = "windows" ]; then
    bin_name="borsdata-mcp.exe"
  fi

  # goreleaser appends microarch suffixes (e.g. _v1 for amd64, _v8.0 for arm64)
  # Try common patterns to locate the binary
  bin_path=""
  for dir in \
    "$DIST_DIR/borsdata-mcp_${os}_${arch}" \
    "$DIST_DIR/borsdata-mcp_${os}_${arch}_v1" \
    "$DIST_DIR/borsdata-mcp_${os}_${arch}_v8.0"; do
    if [ -f "$dir/${bin_name}" ]; then
      bin_path="$dir/${bin_name}"
      break
    fi
  done

  if [ -z "$bin_path" ]; then
    echo "Skipping ${os}/${arch}: binary not found"
    continue
  fi

  # Map OS name to mcpb platform name
  platform_os="$os"
  if [ "$os" = "darwin" ]; then
    platform_os="darwin"
  elif [ "$os" = "windows" ]; then
    platform_os="win32"
  fi

  mcpb_name="borsdata-mcp_${VERSION}_${platform_os}_${arch}.mcpb"

  # Create temp staging directory
  tmp_dir="$(mktemp -d)"
  trap "rm -rf '$tmp_dir'" EXIT

  # Generate manifest.json from template
  sed "s/%%VERSION%%/${VERSION}/g" "$TEMPLATE" > "$tmp_dir/manifest.json"

  # Copy binary into server/ subdirectory
  mkdir -p "$tmp_dir/server"
  cp "$bin_path" "$tmp_dir/server/${bin_name}"

  # Create .mcpb (ZIP archive)
  mcpb_path="$DIST_DIR/$mcpb_name"
  (cd "$tmp_dir" && zip -qr "$mcpb_path" manifest.json server/)

  # Clean up temp dir (reset trap)
  rm -rf "$tmp_dir"
  trap - EXIT

  echo "Created $mcpb_name"
  count=$((count + 1))
done

echo ""
echo "Done: $count .mcpb bundle(s) created in dist/"
