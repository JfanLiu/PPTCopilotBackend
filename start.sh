#!/bin/bash
python3 env.py $SERVER_IP
git config --global --add safe.directory /home/app

bee run
