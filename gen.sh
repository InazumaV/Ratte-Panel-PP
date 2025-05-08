#!/bin/bash
cd api/ || exit 1
swagger generate client -f https://raw.githubusercontent.com/perfect-panel/ppanel-docs/refs/heads/main/public/swagger/node.json
cd ../