# Hello

client (usually a web browser or mobile app). UI actions will be transformed
into spans and metrics by specialized interceptors in the OpenCensus agent.
Clients will typically communicate with the agent over the open internet,
and for web clients, the communication with be over HTTP. For web clients
particularly, it's also important that the amount of client code be small.
These considerations impact how the fields in this proto are specified.

