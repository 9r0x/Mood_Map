#!/bin/sh

curl -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -H "Postman-Token: e954cac3-d640-5e3e-25b2-3e41b9104be8" -d '{
    "content": "不开心，好想要妹子",
    "location": {
      "lat": 123.456,
      "lng": 789.012
    },
    "type": "text",
    "date": 1325556147.0,
    "source": "http://weibo.cn/some/weibo/path"
  }' "http://localhost:8088/points"
