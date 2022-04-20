describe('Profile Page' , ()=> {

    it('Should login and navigate to profile page', ()=>{
        cy.login('nirupamyashas@gmail.com','password');
        cy.get('app-toolbar li').eq(4).click();
        cy.get('#user-items a').eq(0).click();
        cy.url().should('includes','user-profile');
        cy.screenshot();
    })
    
});