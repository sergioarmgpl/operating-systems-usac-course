extern crate base64;
extern crate hyper;
extern crate dotenv;
extern crate hyper_rustls;
extern crate google_pubsub1 as pubsub;
extern crate yup_oauth2 as oauth;
extern crate hyper_tls;

use std::ascii::AsciiExt;
use std::collections::HashMap;
use std::default::Default;
use std::env;
use hyper_rustls::HttpsConnector;
use hyper::client::HttpConnector;
use hyper::client::Client;
use hyper_tls::TlsStream;
use oauth::ServiceAccountAuthenticator;

//Settings env variables
const TOPIC_NAME = env::var("TOPIC_NAME").unwrap();
const AUTH_FILENAME = env::var("AUTH_FILENAME").unwrap();

#[tokio::main]
async fn main() {
    dotenv::dotenv().ok();

    //Creating and setting http connector to http client
    let https = hyper_rustls::HttpsConnector::with_native_roots();
    let cliente = hyper::Client::builder()
    .build::<_, hyper::Body>(https);

    //Setting google credentials
    let creds = yup_oauth2::read_service_account_key(&AUTH_FILENAME)
    .await
    .unwrap();

    //Setting OAUTH2 Token

    let sa: yup_oauth2::authenticator::Authenticator<hyper_rustls::HttpsConnector<hyper::client::HttpConnector>> = ServiceAccountAuthenticator::builder(creds)
    .build().await.unwrap();
   
    //Setting notification body

    let hub = pubsub::Pubsub::new(cliente, sa);
    let methods = hub.projects();
    let message = "PubSub desde Rust :D".to_string();
    
    // Adding our N attributes
    let mut atributos  = HashMap::new();
    atributos.insert("Key1".to_string(), "Value1".to_string());
    atributos.insert("Key2".to_string(), "Value2".to_string());
    atributos.insert("Key3".to_string(), "Value3".to_string());

    
    let message = pubsub::api::PubsubMessage {
        data: Some(base64::encode(message.as_bytes())),
        attributes: Some(atributos),
        ..Default::default()
    };

    // Sending notification

    let request = pubsub::api::PublishRequest { messages: Some(vec![message]) };
    let result = methods.topics_publish(request.clone(), &TOPIC_NAME).doit().await;

    match result {
        Err(e) => {
            println!("Publish error: {}", e);
        }
        Ok((_response, response)) => {
            for msg in response.message_ids.unwrap_or(Vec::new()) {
                println!("Published message #{}", msg);
            }
        }
    }

    println!("completed :D");
}
