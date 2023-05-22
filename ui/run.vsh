#!/usr/bin/env -S v

// print command then execute it
fn sh(cmd string) {
	println('❯ ${cmd}')
	print(execute(cmd).output)
}

sh('sh install.sh')
sh('sh build.sh')
sh('v watch run .')
