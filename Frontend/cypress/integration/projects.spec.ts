describe('Projects Page' , ()=> {

    it('Should login and navigate to projects page', ()=>{
        cy.login('ravulapraveen31@gmail.com','password');
        cy.get('app-toolbar li').eq(4).click();
        cy.get('#user-items a').eq(1).click();
        cy.url().should('includes','projects');
        cy.screenshot();
    })
    
});