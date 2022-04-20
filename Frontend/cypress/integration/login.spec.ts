describe('Login component', () => {
  it('should Not Login if the form details are invalid', () => {
    cy.visit('/login');
    cy.url().should('includes','login');
    cy.get('[formControlName="email"]').type('wrongemail@gmail.com');
    cy.get('#login-button').should('be.disabled');
    cy.url().should('include','/login');
    cy.screenshot();
  });

  it('should Login if the form details are valid because entries are validated correctly', () => {
    cy.login('nirupamyashas@gmail.com','password');
    cy.url().should('include','/');
    cy.screenshot();
  });
})
