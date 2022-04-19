import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { variable } from '../variables/variable';
import { Router } from "@angular/router";

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  departmentForm! : FormGroup;
  searchForm! : FormGroup;

  departments: any = ['Computer Science Engineering', 'Electrical Engineering', 'Mechanical Engineering', 'Law','Arts and Sciences']

  projects: any;

  constructor(public fb: FormBuilder, private router: Router) { }

  ngOnInit(): void {

    this.departmentForm = this.fb.group({
      department_name: ['']
    });

    this.searchForm = this.fb.group({
      search_input: ['']
    });

    this.departmentForm.get('department_name')?.valueChanges.subscribe(data => {
      variable.departmentName = data;
      // console.log(variable.departmentName);
      this.reloadComponent();
    });

  }

  onSubmit(){
    alert(JSON.stringify(this.departmentForm.value))
  }

  reloadComponent() {
    let currentUrl = this.router.url;
        this.router.routeReuseStrategy.shouldReuseRoute = () => false;
        this.router.onSameUrlNavigation = 'reload';
        this.router.navigate([currentUrl]);
  }

  displayProjects(){
    variable.departmentName = "All-Projects";
    this.reloadComponent();
  }

  searchInput(inputValue: any){
    console.log(inputValue.search_input);
    variable.searchInput = inputValue.search_input;
    variable.departmentName = "Search";
    this.reloadComponent();
  }

}
