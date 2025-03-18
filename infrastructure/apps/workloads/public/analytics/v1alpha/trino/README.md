# Trino

Build container

nx run apps-workloads-public-analytics-v1alpha-trino:build
nx run apps-workloads-public-analytics-v1alpha-trino:container
docker run -d -e JAVA_TOOL_OPTIONS="-XX:UseSVE=0" --name trino -p 8081:8080 -v ./etc:/etc 'workloads/public/analytics/v1alpha/trino'
docker run -d -e JAVA_TOOL_OPTIONS="-XX:UseSVE=0" --name trino -p 8081:8080 -v ./etc:/etc -v ./data:/data 'workloads/public/analytics/v1alpha/trino'