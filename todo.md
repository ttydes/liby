# TODO List for ticket making

## Gateway Integration
- [ ] Set up gateway.

## Frontend Development
- [ ] Create a basic HTMX page.
  - [ ] Patch the HTMX library as it has not been merged yet.
- [ ] Demonstrate asynchronous and synchronous rendering by streaming data to the frontend.
  - [ ] (literally just write the word stream on the protofile service. trivial. we want this for media.)

####*example of grpc + htmx. simples:*
#####*https://github.com/vildan-valeev/golang-htmx-grpc/blob/main/index.html#L13

####*example of alternate. htmx bindings: (it'll produce normal html in browser anyway so can revert.)*
#####https://github.com/rajasegar/ccl-demo-raja/blob/main/views/home.lisp

## Monitoring & Metrics
- [ ] Implement Prometheus for metrics collection and notifications via push.
  - [ ] Ensure the configuration is straightforward and efficient.

## Scalability & Deployment
- [ ] Configure Kubernetes clusterized deployment.

## Server Deployment
- [ ] Deploy the application on my server.
  - [ ] Add add a few more servers.
  - [ ] Set up oracle cloud as a last resort backup (cloud bursting).

## Accessibility & Communication
- [ ] Set up accessibility hotline addresses.
  - [ ] Configure custom email addresses for communication.

## Code Examples & Documentation
- [ ] Write a C# example.
  - [ ] Document the differences.

## Code Examples & Documentation
- [ ] Keep a single protogen to make imports the same for every service.
  - [ ] Maintain each service as a standalone project where feasible though.

## Additional Features
- [ ] Implement non-essential middleware (e.g., auto-documentation generation).
