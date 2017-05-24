# Docker Secret Viewer

A simple container that will output all secrets that are mounted into the container.  **Obviously**, be careful where you use this. :P

## Running

```
mkdir tmp
echo "secret-password" > tmp/db-password
docker run -v $(pwd)/tmp:/run/secrets mikesir87/secrets-viewer
```

