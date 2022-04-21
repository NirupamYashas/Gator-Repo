describe('Searching Projects' , ()=> {

    it('Should login and navigate to Searching Projects', ()=>{
        cy.login('ravulapraveen31@gmail.com','password');
        cy.get('app-toolbar li').eq(4).click();
        cy.get('#user-items a').eq(1).click();
        cy.get('[formControlName="department_name"]').select('Computer Science Engineering');
        cy.get('[formControlName="search_input"]').type('test 123');
        cy.url().should('includes','projects');
        cy.screenshot();
    })
    
})