module main
import threefoldtech.threebot.metrics

fn main(){
	args := metrics.MetricsURLArgs{
		network: "development", 
		farm_id:1, 
		node_id: "11"}
	println(metrics.get_metrics_url(args) or {
		println("failed")
		return
	}
	)
}
