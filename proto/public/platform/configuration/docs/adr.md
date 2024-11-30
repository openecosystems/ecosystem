# System Design and Architecture
* Parent has no configuration defined
* Child is created
* Child sets config value
* Parent sets same config value
* Child


Configuration Group
* Workspace has configuration
* Org has configuration
* Workspace and org config is equivalent
* Bound to a group
* Anyone in a group can use this configuration

process
* create config group => empty config created automatically
* login and get workspace config
* give me the configuration for the user
* Check the requesting user belongs to a config group
    * if yes merge user config w/ workspace config
    * show overridden and that it was override by the config group


Configuration:
- id
- organization_slug
- workspace_slug
- created_at
- updated_at
- configuration_type
- configuration_status
- status_details
- parent_id
- Data
  - Catalog
  - Custom Configurations
  - Clinical
- Platform
  - System 1
  - System 2

- Connector
  - Connector 1
    - Configuration
      - URL
      - OAUTH
      - BASIC
      - API TOKEN

Configuration Versioning
- A config can be versioned
- A version can be live, locked, or in draft
- Publish a configuration
