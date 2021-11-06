use actix_web::{HttpRequest, HttpResponse };

pub fn index(_req: HttpRequest) -> HttpResponse {
    HttpResponse::Ok().json(r#"{ data: "Hello World!" }"#)
}
