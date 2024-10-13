# Earn for Glance
This project was born from the collaboration of several companies, combining their expertise and knowledge to create a new product-as-a-service. By integrating **neuroscience** and **artificial intelligence**, we aim to transform the way media content is consumed, bridging the gap between consumers and producers


## The principal goals of the project are: ##

-- Enable consumers to earn money by watching videos.

-- Allow media producers to upload their creations and share them with a community of reviewers.

-- Collect information about consumers, including their facial expressions, opinions, and assessments.

-- Process the collected information using breakthrough technologies.

-- Provide producers with powerful insights to make informed decisions before launching their products.


This document aims to clearly explain how the product works, both technically and functionally, as detailed in the following sections.


## Table of contents

[1. Introduction](#1-introduction)

&nbsp;&nbsp;[1.1 Purpose of this Repository](#11-purpose-of-this-repository)

&nbsp;&nbsp;[1.2 Out of Scope](#12-out-of-scope)

&nbsp;&nbsp;[1.3 Reason](#13-reason)

&nbsp;&nbsp;[1.4 Disclaimer](#14-disclaimer)

&nbsp;&nbsp;[1.5 Give a Star](#15-give-a-star)

&nbsp;&nbsp;[1.6 Share It](#16-share-it)

[2. Domain](#2-domain)

&nbsp;&nbsp;[2.1 Description](#21-description)

&nbsp;&nbsp;[2.2 Conceptual Model](#22-conceptual-model)

&nbsp;&nbsp;[2.3 Event Storming](#23-event-storming)

[3. Architecture](#3-architecture)

&nbsp;&nbsp;[3.0 C4 Model](#30-c4-model)

&nbsp;&nbsp;[3.1 High Level View](#31-high-level-view)

&nbsp;&nbsp;[3.2 Module Level View](#32-module-level-view)

&nbsp;&nbsp;[3.3 API and Module Communication](#33-api-and-module-communication)

&nbsp;&nbsp;[3.4 Module Requests Processing via CQRS](#34-module-requests-processing-via-cqrs)

&nbsp;&nbsp;[3.5 Domain Model Principles and Attributes](#35-domain-model-principles-and-attributes)

&nbsp;&nbsp;[3.6 Cross-Cutting Concerns](#36-cross-cutting-concerns)

&nbsp;&nbsp;[3.7 Modules Integration](#37-modules-integration)

&nbsp;&nbsp;[3.8 Internal Processing](#38-internal-processing)

&nbsp;&nbsp;[3.9 Security](#39-security)

&nbsp;&nbsp;[3.10 Unit Tests](#310-unit-tests)

&nbsp;&nbsp;[3.11 Architecture Decision Log](#311-architecture-decision-log)

&nbsp;&nbsp;[3.12 Architecture Unit Tests](#312-architecture-unit-tests)

&nbsp;&nbsp;[3.13 Integration Tests](#313-integration-tests)

&nbsp;&nbsp;[3.14 System Integration Testing](#314-system-integration-testing)

&nbsp;&nbsp;[3.15 Event Sourcing](#315-event-sourcing)

&nbsp;&nbsp;[3.16 Database change management](#316-database-change-management)

&nbsp;&nbsp;[3.17 Continuous Integration](#317-continuous-integration)

&nbsp;&nbsp;[3.18 Static code analysis](#318-static-code-analysis)

&nbsp;&nbsp;[3.19 System Under Test SUT](#319-system-under-test-sut)

&nbsp;&nbsp;[3.20 Mutation Testing](#320-mutation-testing)

[4. Technology](#4-technology)

[5. How to Run](#5-how-to-run)

[6. Contribution](#6-contribution)

[7. Roadmap](#7-roadmap)

[8. Authors](#8-authors)

[9. License](#9-license)

[10. Inspirations and Recommendations](#10-inspirations-and-recommendations)



## 1. Introduction

### 1.1 Purpose of this Repository
### 1.2 Out of Scope
### 1.3 Reason
### 1.4 Disclaimer
### 1.5 Give a Star
### 1.6 Share It

## 2. Domain

### 2.1 Description

**Definition:**

> Domain - A sphere of knowledge, influence, or activity. The subject area to which the user applies a program is the domain of the software. [Domain-Driven Design Reference](http://domainlanguage.com/ddd/reference/), Eric Evans

### 2.2 Conceptual Model
**Definition:**

> Conceptual Model - A conceptual model is a representation of a system, made of the composition of concepts that are used to help people know, understand, or simulate a subject the model represents. [Wikipedia - Conceptual model](https://en.wikipedia.org/wiki/Conceptual_model)

**Conceptual Model**

PlantUML version:

VisualParadigm version (not maintained, only for demonstration):

**Conceptual Model of commenting feature**


### 2.3 Event Storming
While a Conceptual Model focuses on structures and relationships between them, **behavior** and **events** that occur in our domain are more important.

There are many ways to show behavior and events. One of them is a light technique called [Event Storming](https://www.eventstorming.com/) which is becoming more popular. Below are presented 3 main business processes using this technique: user registration, meeting group creation and meeting organization.

Note: Event Storming is a light, live workshop. One of the possible outputs of this workshop is presented here. Even if you are not doing Event Storming workshops, this type of process presentation can be very valuable to you and your stakeholders.

## 3. Architecture

### 3.0 C4 Model

[C4 model](https://c4model.com/) is a lean graphical notation technique for modelling the architecture of software systems. <br>

As can be found on the website of the author of this model ([Simon Brown](https://simonbrown.je/)): *The C4 model was created as a way to help software development teams describe and communicate software architecture, both during up-front design sessions and when retrospectively documenting an existing codebase* <br>

*Model C4* defines 4 levels (views) of the system architecture: *System Context*, *Container*, *Component* and *Code*. Below are examples of each of these levels that describe the architecture of this system. <br>

*Note: The [PlantUML](https://plantuml.com/) (diagram as text) component was used to describe all C4 model levels. Additionally, for levels C1-C3, a [C4-PlantUML](https://github.com/plantuml-stdlib/C4-PlantUML) plug-in connecting PlantUML with the C4 model was used*.


#### 3.0.1 C1 System Context


#### 3.0.2 C2 Container


#### 3.0.3 C3 Component (high-level)


#### 3.0.4 C3 Component (module-level)


#### 3.0.5 C4 Code (meeting group aggregate)


### 3.1 High Level View

**Module descriptions:**

### 3.2 Module Level View

Each Module has [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and consists of the following submodules (assemblies):

- **Application** - the application logic submodule which is responsible for requests processing: use cases, domain events, integration events, internal commands.
- **Domain** - Domain Model in Domain-Driven Design terms implements the applicable [Bounded Context](https://martinfowler.com/bliki/BoundedContext.html)
- **Infrastructure** - infrastructural code responsible for module initialization, background processing, data access, communication with Events Bus and other external components or systems
- **IntegrationEvents** - **Contracts** published to the Events Bus; only this assembly can be called by other modules

**Note:** Application, Domain and Infrastructure assemblies could be merged into one assembly. Some people like horizontal layering or more decomposition, some don't. Implementing the Domain Model or Infrastructure in separate assembly allows encapsulation using the [`internal`](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/internal) keyword. Sometimes Bounded Context logic is not worth it because it is too simple. As always, be pragmatic and take whatever approach you like.

### 3.3 API and Module Communication

The API only communicates with Modules in two ways: during module initialization and request processing.

### 3.4 Module Requests Processing via CQRS
Processing of Commands and Queries is separated by applying the architectural style/pattern [Command Query Responsibility Segregation (CQRS)](https://docs.microsoft.com/en-us/azure/architecture/patterns/cqrs).

Commands are processed using *Write Model* which is implemented using DDD tactical patterns:

Queries are processed using *Read Model* which is implemented by executing raw SQL statements on database views:

**Key advantages:**

- Solution is appropriate to the problem - reading and writing needs are usually different
- Supports [Single Responsibility Principle](https://en.wikipedia.org/wiki/Single_responsibility_principle) (SRP) - one handler does one thing
- Supports [Interface Segregation Principle](https://en.wikipedia.org/wiki/Interface_segregation_principle) (ISP) - each handler implements interface with exactly one method
- Supports [Parameter Object pattern](https://refactoring.com/catalog/introduceParameterObject.html) - Commands and Queries are objects which are easy to serialize/deserialize
- Easy way to apply [Decorator pattern](https://en.wikipedia.org/wiki/Decorator_pattern) to handle cross-cutting concerns
- Supports Loose Coupling by use of the [Mediator pattern](https://en.wikipedia.org/wiki/Mediator_pattern) - separates invoker of request from handler of request

**Disadvantage:**

- Mediator pattern introduces extra indirection and is harder to reason about which class handles the request

For more information: [Simple CQRS implementation with raw SQL and DDD](https://www.kamilgrzybek.com/design/simple-cqrs-implementation-with-raw-sql-and-ddd/)

### 3.5 Domain Model Principles and Attributes

### 3.6 Cross-Cutting Concerns

### 3.7 Modules Integration

Integration between modules is strictly **asynchronous** using Integration Events and the In Memory Event Bus as broker. In this way coupling between modules is minimal and exists only on the structure of Integration Events.

**Modules don't share data** so it is not possible nor desirable to create a transaction which spans more than one module. To ensure maximum reliability, the [Outbox / Inbox pattern](http://www.kamilgrzybek.com/design/the-outbox-pattern/) is used. This pattern provides accordingly *"At-Least-Once delivery"* and *"At-Least-Once processing"*.

The Outbox and Inbox is implemented using two SQL tables and a background worker for each module. The background worker is implemented using the Quartz.NET library.

**Saving to Outbox:**

**Processing Outbox:**


### 3.8 Internal Processing
The main principle of this system is that you can change its state only by calling a specific Command.

Commands can be called not only by the API, but by the processing module itself. The main use case which implements this mechanism is data processing in eventual consistency mode when we want to process something in a different process and transaction. This applies, for example, to Inbox processing because we want to do something (calling a Command) based on an Integration Event from the Inbox.

This idea is taken from Alberto's Brandolini's Event Storming picture called "The picture that explains “almost” everything" which shows that every side effect (domain event) is created by invoking a Command on Aggregate. See [EventStorming cheat sheet](https://xebia.com/blog/eventstorming-cheat-sheet/) article for more details.

Implementation of internal processing is very similar to implementation of the Outbox and Inbox. One SQL table and one background worker for processing. Each internally processing Command must inherit from `InternalCommandBase` class:


### 3.9 Security

**Authentication**

Authentication is implemented using JWT Token and Bearer scheme using IdentityServer. For now, only one authentication method is implemented: forms style authentication (username and password) via the OAuth2 [Resource Owner Password Grant Type](https://www.oauth.com/oauth2-servers/access-tokens/password-grant/). It requires implementation of the `IResourceOwnerPasswordValidator` interface:


**Authorization**

Authorization is achieved by implementing [RBAC (Role Based Access Control)](https://en.wikipedia.org/wiki/Role-based_access_control) using Permissions. Permissions are more granular and a much better way to secure your application than Roles alone. Each User has a set of Roles and each Role contains one or more Permission. The User's set of Permissions is extracted from all Roles the User belongs to. Permissions are always checked on `Controller` level - never Roles:


### 3.10 Unit Tests

**Definition:**

>A unit test is an automated piece of code that invokes the unit of work being tested, and then checks some assumptions about a single end result of that unit. A unit test is almost always written using a unit testing framework. It can be written easily and runs quickly. It’s trustworthy, readable, and maintainable. It’s consistent in its results as long as production code hasn’t changed. [Art of Unit Testing 2nd Edition](https://www.manning.com/books/the-art-of-unit-testing-second-edition) Roy Osherove

**Attributes of good unit test**

- Automated
- Maintainable
- Runs very fast (in ms)
- Consistent, Deterministic (always the same result)
- Isolated from other tests
- Readable
- Can be executed by anyone
- Testing public API, not internal behavior (overspecification)
- Looks like production code
- Treated as production code

**Implementation**

Unit tests should mainly test business logic (domain model): </br>

Each unit test has 3 standard sections: Arrange, Act and Assert:

**1\. Arrange**

The Arrange section is responsible for preparing the Aggregate for testing the public method that we want to test. This public method is often called (from the unit tests perspective) the SUT (system under test).

Creating an Aggregate ready for testing involves **calling one or more other public constructors/methods** on the Domain Model. At first it may seem that we are testing too many things at the same time, but this is not true. We need to be one hundred percent sure that the Aggregate is in a state exactly as it will be in production. This can only be ensured when we:

- **Use only public API of Domain Model**
- Don't use [InternalsVisibleToAttribute](https://docs.microsoft.com/en-us/dotnet/api/system.runtime.compilerservices.internalsvisibletoattribute?view=netframework-4.8) class
  - This exposes the Domain Model to the Unit Tests library, removing encapsulation so our tests and production code are treated differently and it is a very bad thing
- Don't use [ConditionalAttribute](https://docs.microsoft.com/en-us/dotnet/api/system.diagnostics.conditionalattribute?view=netframework-4.8) classes - it reduces readability and increases complexity
- Don't create any special constructors/factory methods for tests (even with conditional compilation symbols)
  - Special constructor/factory method only for unit tests causes duplication of business logic in the test itself and focuses on state - this kind of approach causes the test to be very sensitive to changes and hard to maintain
- Don't remove encapsulation from Domain Model (for example: change keywords from `internal`/`private` to `public`)
- Don't make methods `protected` to inherit from tested class and in this way provide access to internal methods/properties

**Isolation of external dependencies**

There are 2 main concepts - stubs and mocks:

> A stub is a controllable replacement for an existing dependency (or collaborator) in the system. By using a stub, you can test your code without dealing with the dependency directly.

>A mock object is a fake object in the system that decides whether the unit test has passed or failed. It does so by verifying whether the object under test called the fake object as expected. There’s usually no more than one mock per test.
>[Art of Unit Testing 2nd Edition](https://www.manning.com/books/the-art-of-unit-testing-second-edition) Roy Osherove

Good advice: use stubs if you need to, but try to avoid mocks. Mocking causes us to test too many internal things and leads to overspecification.

**2\. Act**

This section is very easy - we execute **exactly one** public method on aggregate (SUT).

**3\. Assert**

In this section we check expectations. There are only 2 possible outcomes:

- Method completed and Domain Event(s) published
- Business rule was broken

  ### 3.11 Architecture Decision Log

All Architectural Decisions (AD) are documented in the [Architecture Decision Log (ADL)](docs/architecture-decision-log).

More information about documenting architecture-related decisions in this way : [https://github.com/joelparkerhenderson/architecture_decision_record](https://github.com/joelparkerhenderson/architecture_decision_record)


### 3.12 Architecture Unit Tests

In some cases it is not possible to enforce the application architecture, design or established conventions using compiler (compile-time). For this reason, code implementations can diverge from the original design and architecture. We want to minimize this behavior, not only by code review.</br>

To do this, unit tests of system architecture, design, major conventions and assumptions  have been written. In .NET there is special library for this task: [NetArchTest](https://github.com/BenMorris/NetArchTest). This library has been written based on the very popular JAVA architecture unit tests library - [ArchUnit](https://www.archunit.org/).</br>

Using this kind of tests we can test proper layering of our application, dependencies, encapsulation, immutability, DDD correct implementation, naming, conventions and so on - everything what we need to test. Example:</br>

More information about architecture unit tests here: [https://blogs.oracle.com/javamagazine/unit-test-your-architecture-with-archunit](https://blogs.oracle.com/javamagazine/unit-test-your-architecture-with-archunit)

### 3.13 Integration Tests

#### Definition

"Integration Test" term is blurred. It can mean test between classes, modules, services, even systems - see [this](https://martinfowler.com/bliki/IntegrationTest.html) article (by Martin Fowler). </br>

For this reason, the definition of integration test in this project is as follows:</br>

- it verifies how system works in integration with "out-of-process" dependencies - database, messaging system, file system or external API
- it tests particular use case
- it can be slow (as opposed to Unit Test)

#### Approach

- **Do not mock dependencies over which you have full control** (like database). Full control dependency means you can always revert all changes (remove side-effects) and no one can notice it. They are not visible to others. See next point, please.
- **Use "production", normal, real database version**. Some use e.g. in memory repository, some use light databases instead "production" version. This is still mocking. Testing makes sense if we have full confidence in testing. You can't trust the test if you know that the infrastructure in the production environment will vary. Be always as close to production environment as possible.
- **Mock dependencies over which you don't have control**. No control dependency means you can't remove side-effects after interaction with this dependency (external API, messaging system, SMTP server etc.). They can be visible to others.

#### Implementation

Integration test should test exactly one use case. One use case is represented by one Command/Query processing so CommandHandler/QueryHandler in Application layer is perfect starting point for running the Integration Test:</br>

For each test, the following preparation steps must be performed:</br>

1. Clear database
2. Prepare mocks
3. Initialize testing module

After preparation, test is performed on clear database. Usually, it is the execution of some (or many) Commands and: </br>
a) running a Query or/and  </br>
b) verifying mocks </br>
to check the result.

Each Command/Query processing is a separate execution (with different object graph resolution, context, database connection etc.) thanks to Composition Root of each module. This behavior is important and desirable.

### 3.14 System Integration Testing

#### Definition

[System Integration Testing (SIT)](https://en.wikipedia.org/wiki/System_integration_testing) is performed to verify the interactions between the modules of a software system. It involves the overall testing of a complete system of many subsystem components or elements.

#### Implementation

Implementation of system integration tests is based on approach of integration testing of modules in isolation (invoking commands and queries) described in the previous section.

The problem is that in this case we are dealing with **asynchronous communication**. Due to asynchrony, our **test must wait for the result** at certain times.

To correctly implement such tests, the **Sampling** technique and implementation described in the [Growing Object-Oriented Software, Guided by Tests](https://www.amazon.com/Growing-Object-Oriented-Software-Guided-Tests/dp/0321503627) book was used:

>An asynchronous test must wait for success and use timeouts to detect failure. This implies that every tested activity must have an observable effect: a test must affect the system so that its observable state becomes different. This sounds obvious but it drives how we think about writing asynchronous tests. If an activity has no observable effect, there is nothing the test can wait for, and therefore no way for the test to synchronize with the system it is testing. There are two ways a test can observe the system: by sampling its observable state or by listening for events that it sends out.

Test below:

1. Creates Meeting Group Proposal in Meetings module
2. Waits until Meeting Group Proposal to verification will be available in Administration module with 10 seconds timeout
3. Accepts Meeting Group Proposal in Administration module
4. Waits until Meeting Group is created in Meetings module with 15 seconds timeout


### 3.15 Event Sourcing

#### Theory

During the implementation of the Payment module, *Event Sourcing* was used. *Event Sourcing* is a way of preserving the state of our system by recording a sequence of events. No less, no more.

It is important here to really restore the state of our application from events. If we collect events only for auditing purposes, it is an [Audit Log/Trail](https://en.wikipedia.org/wiki/Audit_trail) - not the *Event Sourcing*.

The main elements of *Event Sourcing* are as follows:

- Events Stream
- Objects that are restored based on events. There are 2 types of such objects depending on the purpose:
-- Objects responsible for the change of state. In Domain-Driven Design they will be *Aggregates*.
-- *Projections*: read models prepared for a specific purpose
- *Subscriptions* : a way to receive information about new events
- *Snapshots*: from time to time, objects saved in the traditional way for performance purposes. Mainly used if there are many events to restore the object from the entire event history. (Note: there is currently no snapshot implementation in the project)

#### Tool

In order not to reinvent the wheel, the *SQL Stream Store* library was used. As the [documentation](https://sqlstreamstore.readthedocs.io/en/latest/) says:

*SQL Stream Store is a .NET library to assist with developing applications that use event sourcing or wish to use stream based patterns over a relational database and existing operational infrastructure.*

Like every library, it has its limitations and assumptions (I recommend the linked documentation chapter "Things you need to know before adopting"). For me, the most important 2 points from this chapter are:

1. *"Subscriptions (and thus projections) are **eventually consistent** and always will be."* This means that there will always be an inconsistency time from saving the event to the stream and processing the event by the projector(s).
2. *"No support for ambient System.Transaction scopes enforcing the concept of the stream as the consistency and transactional boundary."* This means that if we save the event to a events stream and want to save something **in the same transaction**, we must use [TransactionScope](https://learn.microsoft.com/en-us/dotnet/api/system.transactions.transactionscope?view=net-8.0). If we cannot use *TransactionScope* for some reason, we must accept the Eventual Consistency also in this case.

Other popular tools:

- [EventStore](https://eventstore.com/) *"An industrial-strength database solution built from the ground up for event sourcing."*
- [Marten](https://martendb.io/) *".NET Transactional Document DB and Event Store on PostgreSQL"*

#### Implementation

There are 2 main "flows" to handle:

- Command handling: change of state - adding new events to stream (writing)
- Projection of events to create read models

##### Command Handling

The whole process looks like this:

1. We create / update an aggregate by creating an event
2. We add changes to the Aggregate Store. This is the class responsible for writing / loading our aggregates. We are not saving changes yet.
3. As part of Unit Of Work  a) Aggregate Store adds events to the stream b) messages are added to the Outbox

##### Events Projection

The whole process looks like this:

1. Special class `Subscriptions Manager` subscribes to Events Store (using SQL Store Stream library)
2. Events Store raises `StreamMessageRecievedEvent`
3. `Subscriptions Manager` invokes all projectors
4. If projector know how to handle given event, it updates particular read model. In current implementation it updates special table in SQL database.

#### Sample view of Event Store

Sample *Event Store* view after execution of SubscriptionLifecycleTests Integration Test which includes following steps:

1. Creating Price List
2. Buying Subscription
3. Renewing Subscription
4. Expiring Subscription

looks like this (*SQL Stream Store* table - *payments.Messages*):

### 3.16 Database Change Management

Database change management is accomplished by *migrations/transitions* versioning. Additionally, the current state of the database structure is also versioned.

Migrations are applied using a simple [DatabaseMigrator](src/Database/DatabaseMigrator) console application that uses the [DbUp](https://dbup.readthedocs.io/en/latest/) library. The current state of the database structure is kept in the [SSDT Database Project](https://docs.microsoft.com/en-us/sql/ssdt/how-to-create-a-new-database-project).

The database update is performed by running the following command:

```shell
dotnet DatabaseMigrator.dll "connection_string" "scripts_directory_path"
```

The entire solution is described in detail in the following articles:

1. [Database change management](https://www.kamilgrzybek.com/database/database-change-management/) (theory)
2. [Using database project and DbUp for database management](https://www.kamilgrzybek.com/database/using-database-project-and-dbup-for-database-management/) (implementation)

### 3.17 Continuous Integration

#### Definition

As defined on [Martin Fowler's website](https://martinfowler.com/articles/continuousIntegration.html):
> *Continuous Integration is a software development practice where members of a team integrate their work frequently, usually each person integrates at least daily - leading to multiple integrations per day. Each integration is verified by an automated build (including test) to detect integration errors as quickly as possible.*

#### YAML Implementation [OBSOLETE]

*Originally the build was implemented using yaml and GitHub Actions functionality. Currently, the build is implemented with NUKE (see next section). See [buildPipeline.yml](.github/workflows/buildPipeline.yml)* file history.

##### Pipeline description

CI was implemented using [GitHub Actions](https://docs.github.com/en/actions/getting-started-with-github-actions/about-github-actions). For this purpose, one workflow, which triggers on Pull Request to *master* branch or Push to *master* branch was created. It contains 2 jobs:

- build test, execute Unit Tests and Architecture Tests
- execute Integration Tests


**Steps description**<br/>
a) Checkout repository - clean checkout of git repository <br/>
b) Setup .NET - install .NET 8.0 SDK<br/>
c) Install dependencies - resolve NuGet packages<br/>
d) Build - build solution<br/>
e) Run Unit Tests - run automated Unit Tests (see section 3.10)<br/>
f) Run Architecture Tests - run automated Architecture Tests (see section 3.12)<br/>
g) Initialize containers - setup Docker container for MS SQL Server<br/>
h) Wait for SQL Server initialization - after container initialization MS SQL Server is not ready, initialization of server itself takes some time so 30 seconds timeout before execution of next step is needed<br/>
i) Create Database - create and initialize database<br/>
j) Migrate Database - execute database upgrade using *DatabaseMigrator* application (see 3.16 section)<br/>
k) Run Integration Tests - perform Integration and System Integration Testing (see section 3.13 and 3.14)<br/>

##### Workflow definition

Workflow definition: [buildPipeline.yml](.github/workflows/buildPipeline.yml)

##### Example workflow execution

Example workflow output:

#### NUKE

[Nuke](https://nuke.build/) is *the cross-platform build automation solution for .NET with C# DSL.*

The 2 main advantages of its use over pure yaml defined in GitHub actions are as follows:

- You run the same code on local machine and in the build server. See [buildPipeline.yml](.github/workflows/buildPipeline.yml)
- You use C# with all the goodness (debugging, compilation, packages, refactoring and so on)

This is how one of the stage definition looks like (execute Build, Unit Tests, Architecture Tests) [Build.cs](build/Build.cs):

If you want to see more complex scenario when integration tests are executed (with SQL Server database creation using docker) see [BuildIntegrationTests.cs](build/BuildIntegrationTests.cs) file.

#### SQL Server database project build

Currently, compilation of database projects is not supported by the .NET Core and dotnet tool. For this reason, the [MSBuild.Sdk.SqlProj](https://github.com/rr-wfm/MSBuild.Sdk.SqlProj/) library was used. In order to do that, you need to create .NET standard library, change SDK and create links to scripts folders. Final [database project](src/Database/CompanyName.MyMeetings.Database.Build/CompanyName.MyMeetings.Database.Build.csproj) looks as follows:


### 3.18 Static code analysis

In order to standardize the appearance of the code and increase its readability, the [StyleCopAnalyzers](https://github.com/DotNetAnalyzers/StyleCopAnalyzers) library was used. This library implements StyleCop rules using the .NET Compiler Platform and is responsible for the static code analysis.<br/>

Using this library is trivial - it is just added as a NuGet package to all projects. There are many ways to configure rules, but currently the best way to do this is to edit the [.editorconfig](src/.editorconfig) file. More information can be found at the link above.<br/>

**Note! Static code analysis works best when the following points are met:**<br/>

1. Each developer has an IDE that respects the rules and helps to follow them
2. The rules are checked during the project build process as part of Continuous Integration
3. The rules are set to *help your system grow*. **Static analysis is not a value in itself.** Some rules may not make complete sense and should be turned off. Other rules may have higher priority. It all depends on the project, company standards and people involved in the project. Be pragmatic.

### 3.19 System Under Test SUT

There is always a need to prepare the entire system in a specific state, e.g. for manual, exploratory, UX / UI tests. The fact that the tests are performed manually does not mean that we cannot automate the preparation phase (Given / Arrange). Thanks to the automation of system state preparation ([System Under Test](https://en.wikipedia.org/wiki/System_under_test)), we are able to recreate exactly the same state in any environment. In addition, such automation can be used later to automate the entire test (e.g. through an [3.13 Integration Tests](#313-integration-tests)).<br/>

The implementation of such automation based on the use of NUKE and the test framework is presented below. As in the case of integration testing, we use the public API of modules.

![](docs/Images/sut-preparation.jpg)

Below is a SUT whose task is to go through the whole process - from setting up a *Meeting Group*, through its *Payment*, adding a new *Meeting* and signing up for it by another user.


### 3.20 Mutation Testing

#### Description

Mutation testing is an approach to test and evaluate our existing tests. During mutation testing a special framework modifies pieces of our code and runs our tests. These modifications are called *mutations* or *mutants*. If a given *mutation* does not cause a failure of at least once test, it means that the mutant has *survived* so our tests are probably not sufficient.


## 4. Technology

List of technologies, frameworks and libraries used for implementation:

- [.NET 8.0](https://dotnet.microsoft.com/download) (platform). Note for Visual Studio users: **VS 2019** is required.
- [MS SQL Server Express](https://www.microsoft.com/en-us/sql-server/sql-server-editions-express) (database)
- [Entity Framework Core 8.0](https://docs.microsoft.com/en-us/ef/core/) (ORM Write Model implementation for DDD)
- [Autofac](https://autofac.org/) (Inversion of Control Container)
- [IdentityServer4](http://docs.identityserver.io) (Authentication and Authorization)
- [Serilog](https://serilog.net/) (structured logging)
- [Hellang.Middleware.ProblemDetails](https://github.com/khellang/Middleware/tree/master/src/ProblemDetails) (API Problem Details support)
- [Swashbuckle](https://github.com/domaindrivendev/Swashbuckle) (Swagger automated documentation)
- [Dapper](https://github.com/StackExchange/Dapper) (micro ORM for Read Model)
- [Newtonsoft.Json](https://www.newtonsoft.com/json) (serialization/deserialization to/from JSON)
- [Quartz.NET](https://www.quartz-scheduler.net/) (background processing)
- [FluentValidation](https://fluentvalidation.net/) (data validation)
- [MediatR](https://github.com/jbogard/MediatR) (mediator implementation)
- [Postman](https://www.getpostman.com/) (API tests)
- [NUnit](https://nunit.org/) (Testing framework)
- [NSubstitute](https://nsubstitute.github.io/) (Testing isolation framework)
- [Visual Paradigm Community Edition](https://www.visual-paradigm.com/download/community.jsp) (CASE tool for modeling and documentation)
- [NetArchTest](https://github.com/BenMorris/NetArchTest) (Architecture Unit Tests library)
- [Polly](https://github.com/App-vNext/Polly) (Resilience and transient-fault-handling library)
- [SQL Stream Store](https://github.com/SQLStreamStore) (Library to assist with Event Sourcing)
- [DbUp](https://dbup.readthedocs.io/en/latest/) (Database migrations deployment)
- [SSDT Database Project](https://docs.microsoft.com/en-us/sql/ssdt/how-to-create-a-new-database-project) (Database structure versioning)
- [GitHub Actions](https://docs.github.com/en/actions) (Continuous Integration workflows implementation)
- [StyleCopAnalyzers](https://github.com/DotNetAnalyzers/StyleCopAnalyzers) (Static code analysis library)
- [PlantUML](https://plantuml.com) (UML diagrams from textual description, diagrams as text)
- [C4 Model](https://c4model.com/) (Model for visualising software architecture)
- [C4-PlantUML](https://github.com/plantuml-stdlib/C4-PlantUML) (C4 Model for PlantUML plugin)
- [NUKE](https://nuke.build/) (Build automation system)
- [MSBuild.Sdk.SqlProj](https://github.com/rr-wfm/MSBuild.Sdk.SqlProj/) (Database project compilation)
- [Stryker.NET](https://stryker-mutator.io/docs/stryker-net/Introduction) (Mutation Testing framework)

## 5. How to Run

### Install .NET 8.0 SDK

- [Download](https://dotnet.microsoft.com/en-us/download/dotnet/8.0) and install .NET 8.0 SDK

### Create database

- Download and install MS SQL Server Express or other
- Create an empty database using [CreateDatabase_Windows.sql](src/Database/CompanyName.MyMeetings.Database/Scripts/CreateDatabase_Windows.sql) or [CreateDatabase_Linux.sql](src/Database/CompanyName.MyMeetings.Database/Scripts/CreateDatabase_Linux.sql). Script adds **app** schema which is needed for migrations journal table. Change database file path if needed.
- Run database migrations using **MigrateDatabase** NUKE target by executing the build.sh script present in the root folder:

```shell
.\build MigrateDatabase --DatabaseConnectionString "connection_string"
```

*"connection_string"* - connection string to your database

### Seed database

- Execute [SeedDatabase.sql](src/Database/CompanyName.MyMeetings.Database/Scripts/SeedDatabase.sql) script
- 2 test users will be created - check the script for usernames and passwords

### Configure connection string

Set a database connection string called `MeetingsConnectionString` in the root of the API project's appsettings.json or use [Secrets](https://blogs.msdn.microsoft.com/mihansen/2017/09/10/managing-secrets-in-net-core-2-0-apps/)

### Configure startup in IDE

- Set the Startup Item in your IDE to the API Project, not IIS Express

### Authenticate

- Once it is running you'll need a token to make API calls. This is done via OAuth2 [Resource Owner Password Grant Type](https://www.oauth.com/oauth2-servers/access-tokens/password-grant/). By default IdentityServer is configured with the following:
- `client_id = ro.client`
- `client_secret = secret` **(this is literally the value - not a statement that this value is secret!)**
- `scope = myMeetingsAPI openid profile`
- `grant_type = password`

Include the credentials of a test user created in the [SeedDatabase.sql](src/Database/CompanyName.MyMeetings.Database/Scripts/SeedDatabase.sql) script - for example:

- `username = testMember@mail.com`
- `password = testMemberPass`

**Example HTTP Request for an Access Token:**

```http
POST /connect/token HTTP/1.1
Host: localhost:5000
 
grant_type=password
&username=testMember@mail.com
&password=testMemberPass
&client_id=ro.client
&client_secret=secret
```

This will fetch an access token for this user to make authorized API requests using the HTTP request header `Authorization: Bearer <access_token>`

If you use a tool such as Postman to test your API, the token can be fetched and stored within the tool itself and appended to all API calls. Check your tool documentation for instructions.

### Run using Docker Compose

You can run whole application using [docker compose](https://docs.docker.com/compose/) from root folder:

```shell
docker-compose up
```

It will create following services: <br/>

- MS SQL Server Database
- Database Migrator
- Application

### Run Integration Tests in Docker

You can run all Integration Tests in Docker (exactly the same process is executed on CI) using **RunAllIntegrationTests** NUKE target:

```shell
.\build RunAllIntegrationTests
```

## 6. Contribution

This project is still under analysis and development. I assume its maintenance for a long time and I would appreciate your contribution to it. Please let me know by creating an Issue or Pull Request.

## 7. Roadmap

List of features/tasks/approaches to add:

| Name                               | Status | Release date |
|------------------------------------| -------- |--------------|
| Domain Model Unit Tests            |Completed | 2019-09-10   |
| Architecture Decision Log update   |  Completed | 2019-11-09   |
| Integration automated tests        | Completed | 2020-02-24   |
| Migration to .NET Core 3.1         |Completed  | 2020-03-04   |
| System Integration Testing         | Completed  | 2020-03-28   |
| More advanced Payments module      | Completed  | 2020-07-11   |
| Event Sourcing implementation      | Completed  | 2020-07-11   |
| Database Change Management         | Completed  | 2020-08-23   |
| Continuous Integration             | Completed  | 2020-09-01   |
| StyleCop Static Code Analysis      | Completed  | 2020-09-05   |
| FrontEnd SPA application           | Completed | 2020-11-08   |
| Docker support                     | Completed | 2020-11-26   |
| PlantUML Conceptual Model          | Completed | 2021-03-22   |
| C4 Model                           | Completed | 2021-03-29   |
| Meeting comments feature           | Completed | 2021-03-30   |
| NUKE build automation              | Completed | 2021-06-15   |
| Database project compilation on CI | Completed | 2021-06-15   |
| System Under Test implementation   | Completed | 2022-07-17   |
| Mutation Testing                   | Completed | 2022-08-23   |
| Migration to .NET 8.0              | Completed | 2023-12-09   |

NOTE: Please don't hesitate to suggest something else or a change to the existing code. All proposals will be considered.

## 8. Authors

Kamil Grzybek

Blog: [https://kamilgrzybek.com](https://kamilgrzybek.com)

Twitter: [https://twitter.com/kamgrzybek](https://twitter.com/kamgrzybek)

LinkedIn: [https://www.linkedin.com/in/kamilgrzybek/](https://www.linkedin.com/in/kamilgrzybek/)

GitHub: [https://github.com/kgrzybek](https://github.com/kgrzybek)

### 8.1 Main contributors

- [Andrei Ganichev](https://github.com/AndreiGanichev)
- [Bela Istok](https://github.com/bistok)
- [Almar Aubel](https://github.com/AlmarAubel)

## 9. License

The project is under [MIT license](https://opensource.org/licenses/MIT).

## 10. Inspirations and Recommendations
