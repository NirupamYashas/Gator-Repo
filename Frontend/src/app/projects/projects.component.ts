import { ProjectsService } from './projects.service';

import {  Component , OnInit} from '@angular/core';

@Component({
    selector: 'projects',
    templateUrl: './projects.component.html',
    styleUrls: ['./projects.component.css']
})
export class ProjectsComponent {
    title = "List Of Projects";
    
    constructor(public service: ProjectsService){}
    
    projects: any;
    ngOnInit(): void{

        this.service.getProjects().subscribe(data => {
            this.projects = data;
        })  
    }
}