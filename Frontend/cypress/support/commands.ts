// ***********************************************
// This example namespace declaration will help
// with Intellisense and code completion in your
// IDE or Text Editor.
// ***********************************************
declare namespace Cypress {
  interface Chainable<Subject = any> {
    login(email: string,password: string): typeof login;
    register(firstname: string,lastname:string,email: string,password: string):typeof register;
  }
}

function login(email: string,password: string): void {
    cy.visit('/login');
    cy.url().should('includes','login');
    cy.get('[formControlName="email"]').type(email);
    cy.get('[formControlName="password"]').type(password);
    cy.get('#login-button').click();
}

function register(firstname: string,lastname: string,email: string,password: string): void {
  cy.visit('/register');
  cy.url().should('includes','register');
  cy.get('[formControlName="firstname"]').type(firstname);
  cy.get('[formControlName="lastname"]').type(lastname);
  cy.get('[formControlName="email"]').type(email);
  cy.get('[formControlName="password"]').type(password);
  cy.get('#register-button').click();
}
//
// NOTE: You can use it like so:
Cypress.Commands.add('login', login);
Cypress.Commands.add('register', register);
//
// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add("login", (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add("drag", { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add("dismiss", { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite("visit", (originalFn, url, options) => { ... })
