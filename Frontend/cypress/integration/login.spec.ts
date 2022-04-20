describe('Login component', () => {
  it('should Not Login if the form details are invalid beacuse wrong email is entered', () => {
    cy.visit('/login');
    cy.url().should('includes','login');
    cy.get('[formControlName="email"]').type('wrongemail@gmail.com');
    cy.get('[formControlName="password"]').type('password');
    cy.get('#login-button').click();
    cy.url().should('include','/login');
    cy.screenshot();
  });

  it('should Login if the form details are valid because entries are validated correctly', () => {
    cy.visit('/login');
    cy.url().should('includes','login');
    cy.get('[formControlName="email"]').type('nirupamyashas@gmail.com');
    cy.get('[formControlName="password"]').type('password');
    cy.get('#login-button').click();
    cy.url().should('include','/');
    cy.screenshot();
  });
})
