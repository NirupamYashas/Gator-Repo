describe('Checking Edit Project page' , ()=> {

    it('Should login and navigate to edit project page for any project in the list', ()=>{
        cy.login('ravulapraveen31@gmail.com','password');
        cy.get('app-toolbar li').eq(4).click();
        cy.get('#user-items a').eq(1).click();
        cy.get('#project-row').eq(0).click();
        cy.url().should('includes','projects/project-page');
        cy.screenshot();
    })
    
})