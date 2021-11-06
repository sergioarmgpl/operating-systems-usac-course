# Rust PubSub publisher example

Ejemplo de un publisher de Google Pub/Sub implementado en Rust  
&nbsp;

## Variables de entorno necesarias
&nbsp;

### TOPIC_NAME 
Variable de entorno con el nombre del topic. Ejemplo:

```bash
> TOPIC_NAME=projects/<project-name>/topics/<topic-name>
```
### AUTH_FILENAME 
Variable de entorno con el nombre del archivo .json que contiene la llave privada de la cuenta de Google. Ejemplo:

```bash
> AUTH_FILENAME=pubsub.privatekey.json
```

## Compilar y ejecutar
```bash
> cargo build && cargo run
```





