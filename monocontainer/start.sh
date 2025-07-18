#!/bin/bash
echo "Starting Go API"
/usr/local/bin/goapi/main &
echo "Starting Node.js"
http-server ./dist -p 3000 &
echo "Starting Nginx"
nginx -g "daemon off;"