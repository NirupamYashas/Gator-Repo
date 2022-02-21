import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup ,Validators } from '@angular/forms';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {

  addProjectForm! : FormGroup; 

  constructor(private http: HttpClient) { }

  ngOnInit(): void {

    this.addProjectForm = new FormGroup({
      project_name: new FormControl('', [Validators.required]),
      department_name: new FormControl('', [Validators.required]),
      uf_mail: new FormControl('', [Validators.required, Validators.email]),
      github_link: new FormControl('', [Validators.required])
    })
  }

  addProjectSubmit(): void{
     var projectName = this.addProjectForm.getRawValue().project_name;
     var departmentName = this.addProjectForm.getRawValue().department_name;
     var ufMail = this.addProjectForm.getRawValue().uf_mail;
     var githubLink = this.addProjectForm.getRawValue().github_link;
     console.log(this.addProjectForm.getRawValue())
     this.http.post<any>('http://localhost:8080/api/projects/', { name: projectName, department: departmentName,
                          email: ufMail , link: githubLink }).subscribe(data => { })
  }

}
