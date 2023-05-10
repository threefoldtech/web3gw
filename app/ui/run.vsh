#!/usr/bin/env -S v

// print command then execute it
fn sh(cmd string) {
	println('❯ ${cmd}')
	print(execute(cmd).output)
}

sh('sh install.sh')
sh('sh build.sh')
file_str := read_file('templates/playground.html') or {panic(err)}
formatted := file_str.replace('@', '@@')
write_file('templates/playground.html', formatted) or {panic(err)}
sh('v watch run .')