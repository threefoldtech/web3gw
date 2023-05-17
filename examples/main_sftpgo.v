module main

import threefoldtech.threebot.sftpgo

import log


fn users_crud(mut cl sftpgo.SFTPGoClient, mut logger log.Logger){
	mut user := sftpgo.User {
		username: "test_user"
		email: "test_email@test.com"
		password: "test_password"
		public_keys: ["ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDTwULSsUubOq3VPWL6cdrDvexDmjfznGydFPyaNcn7gAL9lRxwFbCDPMj7MbhNSpxxHV2+/iJPQOTVJu4oc1N7bPP3gBCnF51rPrhTpGCt5pBbTzeyNweanhedkKDsCO2mIEh/92Od5Hg512dX4j7Zw6ipRWYSaepapfyoRnNSriW/s3DH/uewezVtL5EuypMdfNngV/u2KZYWoeiwhrY/yEUykQVUwDysW/xUJNP5o+KSTAvNSJatr3FbuCFuCjBSvageOLHePTeUwu6qjqe+Xs4piF1ByO/6cOJ8bt5Vcx0bAtI8/MPApplUU/JWevsPNApvnA/ntffI+u8DCwgP"]
		permissions: {"/": ["*"]}
		status: 1
	}

	// add user 
	created_user := cl.add_user(user)  or {
		logger.error("Failed to add user: $err")
		exit(1)
	}
	logger.info("user created: $created_user")

	returned_user := cl.get_user("test_user") or {
		logger.error("failed to get user: $err")
		exit(1)
	}
	logger.info("got user: $returned_user")

	// update user 
	user.email = "test_email@modified.com"
	cl.update_user(user)  or {
		logger.error("failed to update user: $err")
		exit(1)
	}

	updated_user := cl.get_user("test_user") or {
		logger.error("failed to get updated user: $err")
		exit(1)
	}
	logger.info("updated user: $updated_user")

	deleted := cl.delete_user("test_user") or {
		logger.error("failed to update user: $err")
		exit(1)
	}
	logger.info("user deleted: $deleted")

}
fn folders_crud(mut cl sftpgo.SFTPGoClient, mut logger log.Logger){
	// create folder struct 
	mut folder := sftpgo.Folder{
		name: "folder2"
		mapped_path: "/folder2"
		description: "folder 2 description"
	}

	//list all folders
	folders := cl.list_folders() or { 
		logger.error("failed to list folder: $err")
		exit(1)
	}
	logger.info("folders: $folders")

	//add folder 
	created_folder := cl.add_folder(folder)  or {
		logger.error("failed to add folder: $err")
		exit(1)
	}
	logger.info("folder created: $created_folder")

	//get folder
	returned_folder := cl.get_folder(folder.name) or { 
		logger.error("failed to get folder: $err")
		exit(1)
	}
	logger.info("folder: $returned_folder")

	//update folder 
	folder.description = "folder2 description modified"
	cl.update_folder(folder)  or {
		logger.error("failed to update folder: $err")
		exit(1)
	}
	//get updated folder
	updated_folder := cl.get_folder(folder.name) or { 
		logger.error("failed to get updated folder: $err")
		exit(1)
	}
	logger.info("updated folder: $updated_folder")

	deleted := cl.delete_folder(folder.name) or {
		logger.error("failed to update user: $err")
		exit(1)
	}
	logger.info("folder deleted: $deleted")
}

fn roles_crud(mut cl sftpgo.SFTPGoClient, mut logger log.Logger){
	// create folder struct 
	mut role := sftpgo.Role{
		name: "role1"
		description: "role 1 description"
		users: []
		admins: []
	}

	// list existing roles
	roles := cl.list_roles() or { 
		logger.error("failed to list roles: $err")
		exit(1)
	}
	logger.info("roles: $roles")

	//add Role 
	created_role := cl.add_role(role)  or {
		logger.error("failed to add role: $err")
		exit(1)
	}
	logger.info("role created: $created_role")

	//get role
	returned_role := cl.get_role(role.name) or { 
		logger.error("failed to get folder: $err")
		exit(1)
	}
	logger.info("role: $returned_role")

	//update role
	role.description = "role1 description modified"
	cl.update_role(role)  or {
		logger.error("failed to update role: $err")
		exit(1)
	}
	
	//get updated role
	updated_role := cl.get_role(role.name) or { 
		logger.error("failed to get updated role: $err")
		exit(1)
	}
	logger.info("updated_role: $updated_role")


	//delete role
	deleted := cl.delete_role(role.name) or {
		logger.error("failed to update role: $err")
		exit(1)
	}
	logger.info("role deleted: $deleted")
}


fn get_events(mut cl sftpgo.SFTPGoClient, mut logger log.Logger){
	fs_events := cl.get_fs_events(0, 0, 100, "DESC") or {
		logger.error("failed to list fs events: $err")
		exit(1)
	}

	logger.info("fs_events: $fs_events")

	provider_events := cl.get_provider_events(0, 0, 100, "DESC") or {
		logger.error("failed to list provider events: $err")
		exit(1)
	}
	logger.info("provider_events: $provider_events")

}

fn main(){
	args := sftpgo.SFTPGOClientArgs{
		url: "http://localhost:8080/api/v2",
		jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOlsiQVBJIiwiMTcyLjE3LjAuMSJdLCJleHAiOjE2ODQzMzI0OTMsImp0aSI6ImNoaWRtN2V0bHYwYzczOTA0ODFnIiwibmJmIjoxNjg0MzMxMjYzLCJwZXJtaXNzaW9ucyI6WyIqIl0sInN1YiI6IjE2ODQxNDI1NTkwNDMiLCJ1c2VybmFtZSI6ImFzaHJhZiJ9.oJwpLEiHDaBfTxH38qVIVmje1JBjVMzFV8TdRC-2gug"
	}
	mut cl := sftpgo.new(args)
	cl.header = cl.construct_header()
	mut logger := log.Logger(&log.Log{
		level: .debug
	})
	users_crud(mut cl, mut logger)
	folders_crud(mut cl, mut logger)
	roles_crud(mut cl, mut logger)
	get_events(mut cl, mut logger)

}
