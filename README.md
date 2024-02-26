# Running

To launch a web service you need to build and run a docker container:
```bash
docker build -f Dockerfile -t local/serverless-golang:latest ./
docker run -p 3000:3000 --rm local/serverless-golang:latest python3 /agent/launch.py
```
