#!/bin/bash

export KC_SERVERSCOM_V2_TOKEN=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6Mzc4NDIsInR5cGUiOiJVc2VyIiwiYWNjZXNzX2dyYW50X2lkIjoxMzM5OSwiZXhwIjoyMjAwNDg4NzM2fQ.6eo5AjSwj9mp7tG6btg9il9oBxEdvTszHxOD65EgQ9s

curl -X POST -d @create.json -H "Authorization: Bearer $KC_SERVERSCOM_V2_TOKEN" -H 'Content-Type: application/json' 'https://api.servers.com/v1/hosts/dedicated_servers'
