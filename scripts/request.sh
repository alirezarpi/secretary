#!/bin/bash

export AUTH_COOKIE=$(curl -sS -XPOST -d '{"username": "admin", "password": "$fkMtFaY($vm&XLqK9Qedzp7,9u1%Gpt"}' -D - localhost:6080/api/user/login | grep sc_session_id | cut -d " " -f 2 | tr -d ';')


# Create User
curl -s -XPOST --cookie $AUTH_COOKIE -d '{"username": "pg_sc_user", "password": "admin123", "active": true}' localhost:6080/api/user

# Create Resource
curl -s -XPOST --cookie $AUTH_COOKIE -d '{"name": "Database", "active": true}' localhost:6080/api/resource

# Create ResourceUser
curl -s -XPOST --cookie $AUTH_COOKIE -d '{"user_id": "84a6de0e-4778-4eef-9bd2-789e9c25da79", "resource_id": "9b2175df-854d-4a3a-9370-021216e45060", "active": true}' localhost:6080/api/resource/user

# Create ResourceDatabase
curl -s -XPOST --cookie $AUTH_COOKIE -d '{"resource_user_id": "e4f363a1-a1ff-4c4c-97ba-b073cf7eecf4", "db_host": "127.0.0.1", "db_port": "5432", "name": "test_db", "active": true}' localhost:6080/api/resource/database
