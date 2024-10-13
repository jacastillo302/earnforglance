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

