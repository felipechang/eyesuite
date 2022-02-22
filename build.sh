#!/bin/bash
(
  cd frontend
  npm install -g typescript
  npm install
  tsc
  npm run build
)
go install
go build -o eyesuite
chmod +x ./eyesuite
./eyesuite