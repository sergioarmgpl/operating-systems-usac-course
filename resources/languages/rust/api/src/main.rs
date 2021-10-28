pub mod handlers;

extern crate chrono;
extern crate serde;
extern crate serde_json;
extern crate serde_derive;

extern crate actix;
extern crate actix_web;
extern crate futures;
use actix_web::{App, HttpServer, web};

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
