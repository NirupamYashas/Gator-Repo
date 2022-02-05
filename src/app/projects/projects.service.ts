
export class ProjectsService{
    getProjects(){
        return  [
            { projectName: 'abc',department:'Computer Science', email: 'frank.murphy@ufl.edu',githubLink: 'User' },
            { projectName: 'def',department:'Electrical Engineering', email: 'vic.reynolds@ufl.edu', githubLink: 'User' },
            { projectName: 'ghi',department:'Mechanical Engineering', email: 'gina.jabowski@ufl.edu', githubLink: 'User' },
            { projectName: 'jkl',department:'Law', email: 'jessi.glaser@ufl.edu',githubLink: 'User' },
            { projectName: 'mno',department:'Arts and Sciences', email: 'jay.bilzerian@ufl.edu',githubLink: 'User' }
        ];
    }
}