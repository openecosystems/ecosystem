# Configuration Architecture and Design

- All Configuration across the entire system must live within the Configuration service

# Configurable Entities

- Organization
- Workspace
- Configuration Group

# Configuration Hierarchy

- Organization
  - Workspace
    - Workspace
      - Workspace
  - Configuration Group

# Organization Entity

- An organization is the top most configurable entity

# Workspace Entity

- A workspace can only have one parent
- A Workspace can either have an organization as a parent or another workspace as a parent

# Configuration Group Entity

- A configuration group can only have the organization as a parent
- A Configuration Group has a priority
- When a user belongs to multiple configuration groups, the configuration group with the highest priority will determine the user's configuration.

# Configuration Set

- A configuration set is either a fully contained configuration for an organizational unit or for a workspace
- Configuration sets can be moved from one environment to another
- Configuration sets can only be moved from within the scope of access: Lower Environments or Upper Environments
  - For example, development configuration sets can move to quality but not to sandbox
  - Sandbox configuration sets can move to production
  - This shouldn't be limited by code but by external process

## Promoting a Configuration set

- We should be able to copy all settings from one organization to the same organization in another environment

## Configuration Set Structure

- Organization
- Entity
  - Configurations

### Example Configuration Set Structure

- Organization

  - Workspace Configurations
    - Configuration 1
    - Configuration 2
    - Configuration 3
  - Billing
    - BillingPlan 1
  - Connected Test
    - Configuration 1
    - Configuration 2
    - Configuration 3
  - Connected Pass
    - Configuration 1
    - Configuration 2
    - Configuration 3
  - Health Service
    - Configuration 1
    - Configuration 2
    - Configuration 3
  - Connectors
    - Configuration 1
    - Configuration 2
    - Configuration 3
  - Rate Limits
    - Configuration 1
    - Configuration 2
    - Configuration 3
  - IAM
  - Consent
  - API
    - Enable Disable APIs
  - Audit
  - Connectors
  - Fraud
  - Media
  - Partnerships
  - Email
    - Custom SMTP
  - Solutions

- Organization
  - Billing
    - BillingPlan 1
  - Connected Test
    - Configuration 1
    - Configuration 2
    - Configuration 3

# Configuration Type Primitives

- Boolean
- List
- Map
- Int
- Float
- Double
- String

# Configuration Creation

- When a user creates an organization or workspace, the system creates a configuration

# Configuration Creation

- When a user creates a workspace, the system creates a configuration

# User Configuration Creation

- When a new user registers, the system creates a user configuration
