module main

import threefoldtech.threebot.tfgrid

fn deploy_zdb(mut client tfgrid.TFGridClient, zdbs []tfgrid.ZDB) ![]tfgrid.ZDBResult{
	mut zdb_results := []tfgrid.ZDBResult{}
	for zdb in zdbs{
		res := client.zdb_deploy(zdb)!
		zdb_results << res
	}

	return zdb_results
}

fn get_zdbs(mut client tfgrid.TFGridClient, zdb_names []string) ![]tfgrid.ZDBResult{
	mut zdb_results := []tfgrid.ZDBResult{}
	for zdb_name in zdb_names{
		res := client.zdb_get(zdb_name)!
		zdb_results << res
	}

	return zdb_results
}

fn delete_zdbs(mut client tfgrid.TFGridClient, zdb_names []string) !{
	for zdb_name in zdb_names{
		client.zdb_delete(zdb_name)!
	}
}