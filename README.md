I finally bought the Dream Router! We no longer need this Repository.

https://twitter.com/_jeyraof/status/1537756024106496002

---

# is-ui-available

### Build
```bash
$ go build -o is-ui-available
```

### Usage
```bash
$ ./is-ui-available -config=sample_config.yaml

[{"name":"Cloud Key Gen2 Plus","url":"https://store.ui.com/collections/unifi-network-unifi-os-consoles/products/unifi-cloudkey-plus","in_stock":true}]
```

### syntax for config.yaml
```yaml
products:
  - name: "name for identifying"
    url: "url for ubiquiti item"
  - name: "name for identifying (2)"
    url: "url for ubiquiti item (2)"
```

### Piping
use jq:
```bash
$ ./is-ui-available -config=sample_config.yaml | jq .   
```

```json
[
  {
    "name": "Dream Router",
    "url": "https://store.ui.com/collections/unifi-network-unifi-os-consoles/products/dream-router",
    "in_stock": false
  },
  {
    "name": "Cloud Key Gen2 Plus",
    "url": "https://store.ui.com/collections/unifi-network-unifi-os-consoles/products/unifi-cloudkey-plus",
    "in_stock": true
  }
]
```

use jq for choosing only in stock:
```bash
$ ./is-ui-available -config=config.yaml | jq '[.[] | select(.in_stock)]' 
```
```json
[
  {
    "name": "Cloud Key Gen2 Plus",
    "url": "https://store.ui.com/collections/unifi-network-unifi-os-consoles/products/unifi-cloudkey-plus",
    "in_stock": true
  }
]
```
