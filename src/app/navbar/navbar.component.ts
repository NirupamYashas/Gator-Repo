import { Component, OnInit } from '@angular/core';
import { ProjectsService } from '../projects/projects.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  selectedValue: string = '';
  projects: any;

  //event handler for the select element's change event
  selectChangeHandler (event: any) {
    //update the ui
    this.selectedValue = event.target.value;
  }

  constructor(public service: ProjectsService) { }

  ngOnInit(): void {

    this.service.getProjectsByDepartment(this.selectedValue).subscribe(data => {
      this.projects = data;
  })
  }

}
