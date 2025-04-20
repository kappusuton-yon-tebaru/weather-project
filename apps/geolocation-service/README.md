# geolocation service

## APIs

### Error

Common format

```json
{
  "source": "internal|external"
  "message": "string",
}
```

> Note: external `source` indicates that error is derived from external API source

---

### **GET** `/api/search`

**Query params**

| key  | value                      |
| ---- | -------------------------- |
| text | a free-styled text address |

**Response**

- 200

```json
{
  "results": [
    {
      "latitude": "float",
      "longitude": "float",
      "country": "string",
      "city": "string",
      "address": "string"
    }
  ]
}
```

- non-200
  response with common error format

**Example**

```bash
curl localhost:3000/api/search?text=Bangkok, Thailand
```

---

### **GET** `/api/ip`

**Query params**

| key | value                   |
| --- | ----------------------- |
| ip  | an ipv4 or ipv6 address |

**Response**

- 200

```json
{
  "result": {
    "latitude": "float",
    "longitude": "float",
    "country": "string",
    "city": "string",
    "address": "string"
  }
}
```

- non-200
  response with common error format

**Example**

```bash
curl localhost:3000/api/ip?ip=8.8.8.8
```
