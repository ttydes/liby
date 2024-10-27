# Sequence Diagrams
## Inventory Management
Inventory Management Sequence Diagram with Additional Services
Below is a sequence diagram illustrating the inventory management process;
Including the primary and alternate scenarios, with additional services:

```mermaid
sequenceDiagram
    participant BranchManager
    participant Client
    participant MediaService
    participant InventoryService
    participant BranchService
    participant DataLayer
    participant Database

    Note over BranchManager,Database: Pre-Condition: Branch Manager is authenticated and navigated to the Inventory Management interface.

    alt Primary Scenario
        BranchManager->>Client: Navigate to Inventory Management page
        Client->>BranchManager: Display Inventory Management UI
        BranchManager->>Client: Search/Filter media by criteria
        Client->>MediaService: Request media search
        MediaService->>DataLayer: Query media data
        DataLayer->>Database: Retrieve media data
        Database-->>DataLayer: Return media data
        DataLayer-->>MediaService: Return media data
        MediaService-->>Client: Return filtered media
        Client->>BranchManager: Display filtered media

        BranchManager->>Client: Select transfer button for media with given id
        Client->>InventoryService: Request transfer page
        InventoryService->>BranchService: Request branch data for intra-city branches
        BranchService->>DataLayer: Query branch data
        DataLayer->>Database: Retrieve branch data
        Database-->>DataLayer: Return branch data
        DataLayer-->>BranchService: Return branch data
        BranchService-->>InventoryService: Return branch data

        InventoryService->>DataLayer: Request media counts for all matched branches
        DataLayer->>Database: Retrieve media counts from respective inventory database
        Database-->>DataLayer: Return media counts
        DataLayer-->>InventoryService: Return steam of counts
        InventoryService-->>Client: Return stream of counts
        Client-->>BranchManager: Render counts for respective branches in real-time

        BranchManager->>Client: Enter transfer details (amount, direction)
        Client->>InventoryService: Initiate transfer
        InventoryService->>DataLayer: Initiate ACID transaction for transfer
        DataLayer->>Database: Execute transaction

        Database-->>DataLayer: Transaction successful
        DataLayer-->>InventoryService: Transaction successful
        InventoryService->>Client: Notify transfer success
        Client->>BranchManager: Update UI to reflect new inventory state

        else Alternate Scenario: Transaction Failure
        Database-->>DataLayer: Transaction failed (media out of stock)
        DataLayer-->>InventoryService: Transaction failed
        InventoryService->>Client: Notify transfer failure
        Client-->>BranchManager: Suggest other branches with high availability

    else Alternate Scenario: No branches intra-city with availability
        Client-->>BranchManager: Redirect back to media selection view


    Note over BranchManager,Database: Post-Condition: ACID Transactionality. Data-consistency is guaranteed, and reflected, not eventually consistent, bu at all times.
    end
```

System Scalability:
```mermaid
sequenceDiagram
        actor SysAdmin
        actor Librarian
        participant System
        participant Data-Layer
        participant Cluster-Node
        participant Service1
        participant Service2
        participant Prometheus

        Note over SysAdmin,Prometheus: Pre-Condition: System is configured as a distributable cluster
        Note over SysAdmin,Prometheus: Admin is aware of resource usage via metrics/notification

        alt Primary Scenario
            Prometheus->>SysAdmin: Notify resource usage approaching saturation
            SysAdmin->>System: Declare additional server node
            System->>Cluster-Node: Provision new node
            Cluster-Node-->>System: Node provisioned
            System-->>SysAdmin: Node added to cluster
            Note over System: Additional node expands bandwidth and resources
        end

        alt Alternative Scenario: Containerized-service instance failure
            Service1->>System: Service instance failed
            System->>Service1: Replace failed instance
            Service1-->>System: Instance replaced
            Prometheus->>SysAdmin: Notify instance replacement
        end

        alt Alternative Scenario: New feature added to dynamic front-end
            Librarian->>Service1: Use new feature
            Service1->>System: Increased service-1 usage
            System->>Service1: Exponential backoff on retry requests
            Service1-->>System: Retry requests minimized
            System->>Service2: Increased service-2 usage
            Service2-->>System: Efficiency improved
            System->>System: Detect over/under utilization
            System->>Service1: Provision/trim instances accordingly
            Service1-->>System: Instances adjusted
            Prometheus->>SysAdmin: Notify resource adjustments
        end

        Note over SysAdmin,Prometheus: Post-Condition: System resources expanded, no downtime
```

Logging-in; Authentication Workflow -
Registration Sequence via SSL/TLS or Oauth:
```mermaid
sequenceDiagram
    participant Client
    participant AccountsService
    participant AuthService

    alt Registration via SSL/TLS
        Client->>AccountsService: gRPC Request: Register(username, password, email) [SSL/TLS]
        AccountsService->>AccountsService: Validate registration data
        AccountsService->>AuthService: gRPC Request: CreateUser(username, password, email) [SSL/TLS]
        AuthService->>AuthService: Create user account
        AuthService-->>AccountsService: Return user_id
        AccountsService-->>Client: Return success message [SSL/TLS]
    else Registration via OAuth
        Client->>AccountsService: gRPC Request: RegisterOAuth(OAuth token) [SSL/TLS]
        AccountsService->>AuthService: gRPC Request: CreateUserOAuth(OAuth token) [SSL/TLS]
        AuthService->>AuthService: Create user account using OAuth token
        AuthService-->>AccountsService: Return user_id
        AccountsService-->>Client: Return success message [SSL/TLS]
    end
```

Login Sequence via SSL/TLS or Oauth:
```mermaid
sequenceDiagram
    participant Client
    participant AccountsService
    participant AuthService

    Note over Client,AuthService: Precondition: User must have completed a registration sequence.

    alt Login via SSL/TLS
        Client->>AccountsService: gRPC Request: Login(username, password) [SSL/TLS]
        AccountsService->>AccountsService: Validate username and password
        AccountsService->>AuthService: gRPC Request: GenerateToken(user_id) [SSL/TLS]
        AuthService->>AuthService: Generate JWT token
        AuthService-->>AccountsService: Return JWT token
        AccountsService-->>Client: Return JWT token [SSL/TLS]
    else Login via OAuth
        Client->>AccountsService: gRPC Request: LoginOAuth(OAuth token) [SSL/TLS]
        AccountsService->>AuthService: gRPC Request: GenerateTokenOAuth(OAuth token) [SSL/TLS]
        AuthService->>AuthService: Generate JWT token using OAuth token
        AuthService-->>AccountsService: Return JWT token
        AccountsService-->>Client: Return JWT token [SSL/TLS]
    end
```
