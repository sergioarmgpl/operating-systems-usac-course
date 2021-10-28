### Start API REST whit RUST
Initialice the project 
``` CMD
$ cargo init --bin api
```

Import the dependencies on the Cargo.toml file
``` toml
[dependencies]
actix = "0.8"
actix-web = "1.0.0-beta"
chrono = { version = "0.4.19", default-features= false, features = ["clock", "std"] }
lazy_static = "1.4.0"

# Serialization/Deserialization
serde_json = "1.0"
serde = "1.0"
serde_derive = "1.0"
futures = "0.1"
```


Import the neccesaries crates to main.rs file
``` rs
pub mod handlers;

extern crate chrono;
extern crate serde;
extern crate serde_json;
extern crate serde_derive;

extern crate actix;
extern crate actix_web;
extern crate futures;
use actix_web::{App, HttpServer, web};
```

Write yor main function, creating the sys variable, creating a new server Http, whit the service index, 
from the handlers in the function index, set your IP Address or 0.0.0.0 or 'localhost' and select a port.
start the service and run the sys variable.
``` rs
fn main() {
    let sys = actix::System::new("api-rust");

    HttpServer::new(
        || App::new()
            .service(
                web::resource("/helloWorld")
                    .route(web::get().to_async(handlers::hello::index))
            )
    )
    .bind("0.0.0.0:8088").unwrap()
    .start();

    println!("Started http server: 0.0.0.0:8088");
    let _ = sys.run();
}
```
In the hello.rs file define the acctions to do your index function
``` rs
use actix_web::{HttpRequest, HttpResponse };

pub fn index(_req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok().json(r#"{ data: "Hello World!" }"#)
}
```

Build and run your api
``` CMD
$ cargo run
```