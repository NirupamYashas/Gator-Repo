import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup ,Validators } from '@angular/forms';
import { ProjectsService } from '../projects/projects.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {

  addProjectForm! : FormGroup; 

  constructor(public projectsService: ProjectsService) { }

  msgTrue = false;
  ngOnInit(): void {

    this.addProjectForm = new FormGroup({
      project_name: new FormControl('', [Validators.required]),
      department_name: new FormControl('', [Validators.required]),
      uf_mail: new FormControl('', [Validators.required, Validators.email]),
      github_link: new FormControl('', [Validators.required])
    })
  }

  addProjectSubmit(){
     var projectName = this.addProjectForm.getRawValue().project_name;
     var departmentName = this.addProjectForm.getRawValue().department_name;
     var ufMail = this.addProjectForm.getRawValue().uf_mail;
     var githubLink = this.addProjectForm.getRawValue().github_link;
     const newFormData = { name: projectName, department: departmentName, email: ufMail , link: githubLink };
     console.log(newFormData);
     this.projectsService.createProject(newFormData).subscribe(data => {})
  }

}
