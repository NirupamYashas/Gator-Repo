import { Component, Input, OnInit } from '@angular/core';
import { Project } from '../_models/project';

@Component({
  selector: 'app-project-page',
  templateUrl: './project-page.component.html',
  styleUrls: ['./project-page.component.css']
})
export class ProjectPageComponent implements OnInit {

  projectName: string = 'Project Name';

  @Input() project: Project;

  constructor() { }

  ngOnInit(): void {
    if (!this.project) {
      this.project = {} as Project;
    } else {
      this.setProjectName();
    }
  }

  private setProjectName() {
    this.projectName = `${this.project.name}`;
  }

}
