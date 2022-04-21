describe('Add New Projects Page' , ()=> {

    it('Should login and navigate to new projects page', ()=>{
        cy.login('ravulapraveen31@gmail.com','password');
        cy.get('app-toolbar li').eq(4).click();
        cy.get('#user-items a').eq(2).click();
        cy.url().should('includes','create');
        cy.get('[formControlName="project_name"]').type('New Project Name');
        cy.get('[formControlName="department_name"]').select('Computer Science Engineering');
        cy.get('[formControlName="uf_mail"]').type('student@ufl.edu.in');
        cy.get('[formControlName="github_link"]').type('https://github.com/NirupamYashas/Gator-Repo/');
        cy.get('#create-project').click();
        cy.url().should('include','/projects');
        cy.screenshot();
    })
    
});