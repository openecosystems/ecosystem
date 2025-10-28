## oeco ecosystem join

Join an existing ecosystem

```bash
oeco ecosystem join
```

## Ecosystem Join

### Pre-conditions

-   You must have the oeco cli installed
-   An ecosystem must already exists

### Pre-condition wire validations

### Pre-condition logic validations

### Post-conditions

### Form

-   Pick an ecosystem domain name. This will be your main domain {domain-name}.mesh
    -   (go get -u github.com/segmentio/go-slugify)
-   Type
-   Cidr

### Business Logic

```mermaid
sequenceDiagram
    autonumber
    participant S as System
    participant U as User
    participant CLI as CLI
    participant E as Ecosystem
    links E: {"Details": "https://docs.openecosystems.com/ecosystem"}
    S->>U: System presents form to User
    U->>S: User completes form
    S->>CLI: System downloads metadata file from the ecosystem edge.{ecosystem-name}.cloud
    S->>CLI: System creates a new context file: {ecosystem-name}.yaml
    S->>CLI: System calls create account <br/>internally to create a new Local Machine Service Account credential: <br/>{sanitized.os.hostname}.{ecosystem-name}.mesh
    S->>CLI: System calls provision ecosystem <br/>internally to configure ecosystem: <br/>configurations/api.{ecosystem-name}.mesh
    S->>CLI: System calls documentation client <br/>to pull markdown instructions for next steps
    U-->>E: User can now connect to connectors on the ecosystem or host their own
```
