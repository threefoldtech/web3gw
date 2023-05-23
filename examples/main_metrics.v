module main
import threefoldtech.threebot.metrics

fn main(){
	args := metrics.MetricsURLArgs{
		network: "production", 
		farm_id:100, 
		node_key: "5ET2XwxP6EQ1aLFBtWwJP2EN9CwJexrxPuXzVNtQbKyi5R8q"}
	println(metrics.get_metrics_url(args))
}
