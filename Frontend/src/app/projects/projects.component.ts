import { ProjectsService } from './projects.service';
import { Router,ActivatedRoute } from "@angular/router";

import {  Component , OnInit} from '@angular/core';
import { Project } from '../_models/project';

@Component({
    selector: 'projects',
    templateUrl: './projects.component.html',
    styleUrls: ['./projects.component.css']
})
export class ProjectsComponent {
    title = "List Of Projects";
    
    constructor(public service: ProjectsService,private router: Router,
        private route: ActivatedRoute,){}
    
    projects: any;
    ngOnInit(): void{

        this.service.getProjects().subscribe(data => {
            this.projects = data;
        })  
    }

    editProject(project: Project) {
        let route = '/projects/project-page';
        this.router.navigate([route], { queryParams: { id: project.id } });
    }
}