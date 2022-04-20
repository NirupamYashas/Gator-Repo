describe('Registration component', () => {
    it('should Not Register if the form details are invalid', () => {
      cy.visit('/register');
      cy.url().should('includes','register');
      cy.get('[formControlName="firstname"]').type('wrongemail@gmail.com');
      cy.get('#register-button').should('be.disabled');
      cy.url().should('include','/register');
      cy.screenshot();
    });
  
    it('should Register User Data in database if the form details are valid and login directly to homepage', () => {
      cy.register('Alin','Dobra','newmail@gmail.com','password');
      cy.url().should('include','/');
      cy.screenshot();
    });
  })