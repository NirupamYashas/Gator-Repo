import { ProjectsService } from './projects.service';

import {  Component } from '@angular/core';

@Component({
    selector: 'projects',
    template: `
            <h2>{{ title }}</h2>
            <ul>
                <li *ngFor="let project of projects">
                    {{ project }}
                </li>
            </ul>
    `
})
export class ProjectsComponent {
    title = "List Of Projects";
    projects;

    constructor(service: ProjectsService){
        this.projects = service.getProjects();
    }
}