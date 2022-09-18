# API

A microservice is published by this project that lists the collection of **rediOps**
definitions registered with the service

The services exposed are defined in the original **rediOps** OpenAPI specifcation
and are:

* /health : the health of the service
* /.well-known/devops.json : the **rediOps** definition for this code
* /devops/ : List of **rediOps** brief definitions registered with this service
* /devops/{id} : Detail definition 

**Note**: The listing of the **rediOps** defintions does not by default include
the hosting service code, which is accessed via the _/.well-known/deveops.json_
location.

## Further information

Either browse the code repository or use the **rediOps** specification
at _/.well-known/deveops.json_ to locate the standard 
information.
