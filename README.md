<div align="center">


**Rabbit-client**:
https://github.com/nitingoyal0996/reddit-client



# **Project Report: Rabbit Forum**

**COP5615: Distributed Operating Systems Principles**  
**Fall 2024**

---

**Nitin Goyal**  
Email: [nitin.goyal@ufl.edu](mailto:nitin.goyal@ufl.edu)  
Electrical and Computer Engineering  
University of Florida  

**November 24, 2024**

</div>

---

## Overview

The **Rebbit Engine** project simulates a distributed Reddit-like system. The system comprises two primary components:

1. **Engine**
2. **Client**



### The Engine

The **engine** serves as the backend server, hosting actors that facilitate core functionalities of the Reddit system. The **client** interacts with the engine through an actor called `consumer`, which can simulate thousands of users. 

The engine is designed with multiple abstraction layers to ensure extensibility and performance:

#### **Data Layer**
- Provides an abstraction over the database to handle storage and retrieval operations.
- Implements a repository pattern using SQLite.

#### **Service Layer**
- Implements business logic.
- Interfaces with the data layer to perform operations on stored data.

#### **Actors Layer**
- Manages actors that handle incoming client requests and interact with the service layer. 
- Built using the **Proto.actor** framework to ensure scalability and fault tolerance.

  Key actors implemented:
  1. **AuthActor**: Manages authentication data.  
  2. **UserActor**: Handles user data operations.  
  3. **SubredditActor**: Manages subreddit data.  
  4. **PostActor**: Handles posts within subreddits.  
  5. **KarmaActor**: Manages user subscriptions and interactions.  
  6. **CommentActor**: Handles comments on posts.

#### **Network Layer**
- Uses **protocol buffers** for efficient inter-node communication.  
- Exposes engine functionalities via Go's `http` package for handling API requests.

### Proto.actor Framework

The **Proto.actor** framework enables the system to handle a high volume of concurrent users by leveraging the **distributed actor model**. Key features utilized include:

- **Clustering**: Self-managed actors were deployed using the automanaged clustering package.  
- **Communication**: Protocol buffers were employed for communication between the **engine** and **client**, which run as separate processes.  
- **Fault Tolerance**: Built-in supervision ensures reliable actor lifecycle management.

### The Client

The **client** simulates user interactions and consists of two main components:

1. **Consumer**:
   - An actor responsible for performing actions such as user registration, login, subreddit creation, subscriptions, and posting.
   - Generates test data and simulates user behavior.

2. **Simulator**:
   - Spawns multiple consumers and executes predefined scenarios.  
   - Maintains a pool of active consumer actors.

## Usage

### Running the Engine
1. Navigate to the `./engine` directory.  
2. Install dependencies:  
   ```bash
   go mod tidy
   ```
3. Build the project:  
   ```bash
   go build .
   ```
4. Start the engine:  
   ```bash
   ./reddit-clone
   ```

### Running the Client
1. Navigate to the `./client` directory.
2. Install dependencies:  
   ```bash
   go mod tidy
   ```
3. Build the project:  
   ```bash
   go build .
   ```

4. Run the client simulation:
   ```bash
   ./client <simulation_number> <number_of_users> <number_of_subreddits>
   ```
   **Note**: Ensure the database is reset before each run to avoid errors from test data conflicts.



## Simulations

1. **Simulation 1: Many Users - Registration and Login**  
   - Tests the system’s ability to handle large-scale user registrations and login functionality.

2. **Simulation 2: Many Subreddits - Zipf Member Subscription**  
   - Generates subreddits and users based on Zipf distribution.  
   - Each subscribed user posts twice in their respective subreddits.

3. **Simulation 3: User Connection and Disconnection**  
   - Simulates random user logins and logouts at periodic intervals.  
   - Leverages JWT tokens for stateless authentication, generating a new token for each login.


## Simulation Results

### Scenario 1:

Largest number of users tests - 1000 - registration and login. I am confident it would work for more users as well. But have limited the number of users to 1000 due to time constraints for now.

```bash
./client 1 1000
```
![sim-1](image-2.png)

### Scenario 2:
![sim 2](image.png)

```bash
./client 2 500 10
```
The subscriber count of the reddit here is in Zipf distribution. The number of posts in each subreddit is double the subscriber count.

### Scenario 3:
the default connection duration is set to 5 seconds, it could be changed and increased in the main.go file of client module.

```bash
./client 3 1000
```

![sim 3](image-1.png)

## Conclusion

The **Rabbit Engine** project effectively demonstrates the design and implementation of a distributed system using the actor model. By leveraging the **Proto.actor** framework, the system achieves scalability, fault tolerance, and high concurrency, making it suitable for handling real-world workloads. The modular architecture—comprising data, service, actor, and network layers—ensures extensibility and maintainability. Through comprehensive simulations, the project validates core functionalities such as user registration, subreddit management, and user interactions, showcasing the potential of distributed actor-based systems in building robust, scalable applications. This project serves as a foundational model for developing distributed platforms with high performance and reliability.