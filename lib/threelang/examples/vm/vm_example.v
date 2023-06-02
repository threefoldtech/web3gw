module main

import os
import threefoldtech.threebot.threelang

const testpath = os.dir(@FILE)

fn do() ! {
	mut tl:=threelang.new(path:testpath)!
}

fn main() {
	do() or { panic(err) }
}
