describe('Logout Component' , ()=> {

    it('Should login and navigate to logout', ()=>{
        cy.login('ravulapraveen31@gmail.com','password');
        cy.get('app-toolbar li').eq(4).click();
        cy.get('#user-items a').eq(3).click().should(() => {
            expect(localStorage.getItem('token')).to.be.null;
        }) ;
        cy.url().should('include','/');
        cy.screenshot();
    })
    
});