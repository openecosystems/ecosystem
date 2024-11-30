use fastly::config_store::{ConfigStore};
use std::fmt::{Write};
use std::string::ToString;

pub(crate) struct RoutingRule {
    pub organization_slug: String,
    pub workspace_slug: String,
    pub jan: String,
}

/// Extract Routing Rules
pub(crate) fn extract_workspace_routing_rules(api_key: &str, service_id: &str) -> Option<RoutingRule> {

    let mut store = String::new();
    write!(&mut store, "routes_{}", service_id).unwrap();

    let _routing_rules_config_store = ConfigStore::try_open(&store);
    if _routing_rules_config_store.is_err() {
        println!("Error {}", _routing_rules_config_store.err().unwrap().to_string());
        println!("Routing Rules Store could not be opened: {}", store);
        return None
    }

    let routing_rules_config_store = _routing_rules_config_store.unwrap();

    let _routing_rule = routing_rules_config_store.try_get(api_key);
    if _routing_rule.is_err() {
        println!("Error {}", _routing_rule.err().unwrap().to_string());
        println!("Could not find {} in the routing rules store", api_key);
        return None
    }

    let routing_rule = _routing_rule.unwrap();
    if routing_rule.is_none() {
        println!("Routing Rules Store opened and found rule, however rule is empty");
        return None
    }

    let _routing_rule: Vec<&str> = routing_rule.as_ref().unwrap()
        .split(";")
        .collect();

    if _routing_rule.len() != 3 {
        return None
    }

    let jan = _routing_rule[0];
    let organization_slug = _routing_rule[1];
    let workspace_slug = _routing_rule[2];

    Some(RoutingRule {
        organization_slug: organization_slug.to_string(),
        workspace_slug: workspace_slug.to_string(),
        jan: jan.to_string()
    })
}
