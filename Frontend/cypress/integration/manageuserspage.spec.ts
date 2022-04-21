describe('Manage users by admin authentication' , ()=> {

    it('Should login and navigate to manage users', ()=>{
        cy.login('nirupamyashas@gmail.com','password');
        cy.get('app-toolbar li').eq(4).click();
        cy.get('#user-items a').eq(3).click();
        cy.url().should('includes','admin/users');
        cy.screenshot();
    })
    
});