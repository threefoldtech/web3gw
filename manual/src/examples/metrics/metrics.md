# Metrics

## Get metrics url of a node

- To get a metrics url for a specific node you first need to create MetricsURLArgs instance as follows

```
args := metrics.MetricsURLArgs{
    network: "development",
    farm_id:1,
    node_id: "11"
}

```

- Then call get_metrics_url function passing this struct to it

```
url := metrics.get_metrics_url(args) or {
    println("failed to construct metrics url for this node")
    return
}
```
