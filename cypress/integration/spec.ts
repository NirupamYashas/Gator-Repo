describe('Login component', () => {
  it('should Not Login if the form is invalid beacuse password is not entered so the button will be disabled', () => {
    cy.visit('/login')
    cy.url().should('includes','login')
    cy.get('[formControlName="useremail"]').type('Robinhood@gmail.com')
    cy.get('button').should('be.disabled');
    cy.url().should('not.include','projects');
  });

  it('should Login if the form is valid because entries are validated correctly', () => {
    cy.visit('/login')
    cy.url().should('includes','login')
    cy.get('[formControlName="useremail"]').type('Robinhood@gmail.com')
    cy.get('[formControlName="password"]').type('sr3sr3sr3')
    cy.get('button').click();
    cy.url().should('include','projects');
  });
})
