# -*- mode: Python -*-

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

k8s_yaml(kustomize('./k8s/overlays/local'))

# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')

# Records the current time, then kicks off a server update.
# Normally, you would let Tilt do deploys automatically, but this
# shows you how to set up a custom workflow that measures it.
local_resource(
  'deploy',
  'python ./scripts/record-start-time.py',
)

compile_cmd = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/discount ./cmd/discount'

local_resource(
  'discount-go-compile',
  compile_cmd,
  deps=['./cmd', './internal/'],
  resource_deps=['deploy'])

docker_build_with_restart(
  'phenrigomes/blue-discount',
  '.',
  entrypoint=['/app/build/discount'],
  dockerfile='Dockerfile.dev',
  only=[
    './build',
    './config',
    './migrations',
    './proto'
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./config', '/app/config'),
    sync('./migrations', '/app/migrations'),
  ],
)

k8s_resource('discount-service', port_forwards=8000)
k8s_resource('discount-db', port_forwards=5432)
