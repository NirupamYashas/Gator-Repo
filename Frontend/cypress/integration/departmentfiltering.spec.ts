describe('filtering of departments' , ()=> {

    it('Should login and navigate to filtering of departments', ()=>{
        cy.login('nirupamyashas@gmail.com','password');
        cy.get('app-toolbar li').eq(4).click();
        cy.get('#user-items a').eq(1).click();
        cy.get('[formControlName="department_name"]').select('Computer Science Engineering');
        cy.url().should('includes','projects');
        cy.screenshot();
    })
    
});