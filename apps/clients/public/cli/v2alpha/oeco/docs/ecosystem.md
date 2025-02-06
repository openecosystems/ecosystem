
## oeco ecosystem
Dynamically see economic system members and their capabilities

```bash
oeco ecosystem create
```

## Ecosystem Create Form
### Pre-conditions
- You must have the oeco cli installed

### Post-conditions
- When running the ecosystem, start a pushpin client to listen for discovery changes
- Tabs
    - Active Connectors
        - When it starts, pull the latest available connectors in the network (fetch all keys from the connector bucket),
        - then start pushpin for differentials
- Upon creating an ecosystem, create the following accounts:
    - api.{ecosystem-name}.mesh
    - edge.{ecosystem-name}.mesh
- To create more accounts, use oeco account create

### Form
- Pick an ecosystem name [no spaces, dashes, and only alphanumeric] . This will be your main domain oeco.mesh
  - (go get -u github.com/segmentio/go-slugify)
- Type
- Cidr
