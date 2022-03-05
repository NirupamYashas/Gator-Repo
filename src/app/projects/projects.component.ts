import { ProjectsService } from './projects.service';
import { NavbarComponent } from '../navbar/navbar.component';

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
        //this.projects = this.service.getProjects();
        

        this.service.getProjects().subscribe(data => {
            this.projects = data;
        })
    }

    // refreshProjects(Value: String){
    //     this.service.getProjects(Value).subscribe(data => {
    //         this.projects = data;
    //     })
    // }
   
}