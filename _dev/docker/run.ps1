docker stack down pulse-tools
$filebase = Join-Path $PSScriptRoot "docker-compose.tools.yaml"
docker stack deploy --compose-file $filebase pulse-tools