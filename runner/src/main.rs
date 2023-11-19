use std::{fs, thread, time, path};
use std::fs::read_to_string;
use std::process::{Command, Child};
use std::time::{Duration, Instant};
use reqwest;
use serde_json::json;
use serde::{Deserialize, Serialize};
use toml;


#[derive(Deserialize, Serialize)]
struct Deploy {
    name: String,
    runtime: String,
    entry: String,
    instructions: Instructuions
}

#[derive(Deserialize, Serialize)]
struct Instructuions {
    ent_file: String,
}

fn read_deployment(config_path: &str) -> Deploy {

	let instruct: Deploy = {
        let config_text = fs::read_to_string(config_path).expect("error reading file");
        toml::from_str(&config_text).expect("error reading stream")
    };
    
   return instruct;
}


fn spawn_child( config: Deploy, elapsed: Instant) -> Child {

	if config.runtime == "Deno"{
		let child = Command::new("/root/.deno/bin/deno")
        .arg("run")
		.arg("--allow-all")
        .arg(config.instructions.ent_file)
        .spawn()
        .expect("failed to execute process");
		println!("Code started in: {} seconds", elapsed.elapsed().as_secs());
	    return child;
	} else {
		let child = Command::new("/usr/bin/node")
        .arg(config.instructions.ent_file)
        .spawn()
        .expect("failed to execute process");
		println!("Code started in: {} seconds", elapsed.elapsed().as_secs());
	    return child;
	}
}


fn main() {

	loop {

		let exists = path::Path::new("/app/deploy.toml").is_file();

		if exists{
			println!("File found");
			break;
		}
		println!("File not found");
		thread::sleep(time::Duration::from_secs(1));
	}

	let elapsed = Instant::now();

	let config = read_deployment("/app/deploy.toml");
	let mut child = spawn_child(config, elapsed);

	loop {
		let time_elapsed = elapsed.elapsed();

		if time_elapsed.as_secs() > 100{
			break
		}
		println!("elapsed time: {} seconds", time_elapsed.as_secs());
		std::thread::sleep(Duration::from_secs(1));
    }

	child.kill().expect("failed to kill process");
	send_kill();
	println!("done!")
}


fn send_kill() {
	

	let container_id = read_to_string("/etc/hostname").unwrap();
	

	let client = reqwest::blocking::Client::new();

	let data = json!({
		"ContainerId": container_id.trim(),
		"msg": "kill",
	});
	let json_string = serde_json::to_string(&data).unwrap();


    let res = client
        .post("http://10.0.0.13:3000/app/container/kill")
        .body(json_string)
        .send();

	match res {
		Ok(response) => {
			println!("Status: {}", response.status());
			println!("Headers:\n{:#?}", response.headers());
			println!("Body:\n{}", response.text().unwrap());
		},
		Err(err) => {
			println!("Error sending request: {}", err);
		},
	}
	
	println!("sent");
        
}