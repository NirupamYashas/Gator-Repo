# Gator-Repo

### Project Idea:

Gator Repo is a web application that can be used to maintain all the projects done by students at UF. It offers additional features for professors and other students to verify projects and also allows the contributors to add their project updates.

### Features:

- List out all your projects
- View Projects
   + Filter By Department
   + Search By project name
   + Search By email
- check out your UF peers' projects 
- Rate your peers' projects 

### Team Members

- Jagan Mohan Reddy Dwarampudi - jdwarampudi@ufl.edu - Backend Developer
- Satya Venkata Sai Nirupam Yashas Kuchimanchi - kuchimanchi.s@ufl.edu - Frontend Developer
- Praveen Ravula - praveenravula@ufl.edu - Software Tester and Full Stack

### Technology Stack:
- Front-end: Angular, TypeScript
- Database: SQLite
- Back-end: GOlang

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 13.2.2.

# Running the project
Step 1: [Introduction and Environment Setup for GoLang (Windows & Mac)](https://www.youtube.com/watch?v=dgIh-VYcWYw "Introduction and Environment Setup for GoLang (Windows & Mac)")

Step 2: [Angular Project Setup in Visual Studio Code](https://www.youtube.com/watch?v=ZJejjL1Iev0 "Angular Project Setup in Visual Studio Code")

Step 3: Arrange the files according to the file paths given below 
- Gator-Repo
  - Frontend
    > go to client and run "npm i" t install all npm libraries
    - api
    > run the cleint using command "ng serve" to launch cleint side webite
    - src
    > Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

  - Server
    > install go
    - bin
    - controllers
    - models
    - utils
    - main.go
    > run main.go file using "go run main.go" in termianal in the directory of server
    - Users.db

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `ng build` to build the project. The build artifacts will be stored in the `dist/` directory.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.
