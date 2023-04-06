#! /usr/bin/env sh

URL="http://localhost:9093/api/v1/alerts"

curl -si -X POST -H "Content-Type: application/json" "$URL" -d '
[
    {
        "labels": {
            "alertname": "InstanceDown",
            "instance": "node-test1:9100",
            "job": "node",
            "severity": "critical"
        },
        "annotations": {
            "summary": "Instance is down"
        },
        "generatorURL": "http://node-test1:9100"
    }
]
'