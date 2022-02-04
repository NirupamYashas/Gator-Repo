import { ProjectsService } from './projects.service';

import {  Component } from '@angular/core';

@Component({
    selector: 'projects',
    templateUrl: './projects.component.html',
    styleUrls: ['./projects.component.css']
})
export class ProjectsComponent {
    title = "List Of Projects";
    projects;
    

    constructor(service: ProjectsService){
        this.projects = service.getProjects();
    }
}