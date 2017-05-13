package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Client clientInfo
	Server serverInfo
}

type clientInfo struct {
	Private_key         string
	Public_key          string
	Backup_dirs         []string
	Exclude_dirs        []string
	Threads             int
	Sql_file            string
	Server              string
	Purge_deleted_files int
	Queue_blobs         string
	Notify_email        string
	Server_port         int
}

type serverInfo struct {
	Private_key                  string
	Public_key                   string
	Sql_file                     string
	Local_store                  string
	Notify_email                 string
	Threads                      int
	Contract_grace_period        int
	Keep_files_for               int
	Keep_local_copy_of_all_blobs bool
	Server_port                  int
}

func main() {
	var config tomlConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Client:\n")
	fmt.Printf("private = %s\npublic = %s\nthreads = %d server_port = %d\n",
		config.Client.Private_key, config.Client.Public_key,
		config.Client.Threads, config.Client.Server_port)
	fmt.Printf("backup dirs = %v\n", config.Client.Backup_dirs)
	fmt.Printf("exclude dirs = %v\n", config.Client.Exclude_dirs)
	fmt.Printf("\nServer:\n")
	fmt.Printf("private = %s\npublic = %s\nthreads = %d, server_port = %d\n",
		config.Server.Private_key, config.Server.Public_key,
		config.Server.Threads, config.Server.Server_port)
	fmt.Printf("Blob store = %s\nSQL store = %s\n",
		config.Server.Local_store, config.Server.Sql_file)
}
