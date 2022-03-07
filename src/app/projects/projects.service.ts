import { Injectable } from "@angular/core";
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http";


@Injectable({
    providedIn: 'root'
})
export class ProjectsService{

    constructor(private httpclient: HttpClient){}

    getProjects(){

        //Headers
        const httpHeaders = new HttpHeaders();
        httpHeaders.append('content-type','application/json')

        // if(department === null){
            //Get the HTTP Method working for you
            return  this.httpclient.get('http://localhost:8080/api/projects', {headers: httpHeaders});
        // }else{
        //     return  this.httpclient.get('http://localhost:8080/api/projects/departments/'+department, {headers: httpHeaders});
        // }  

    }

    createProject(createResource: any){
        //Headers
        const httpHeaders = new HttpHeaders();
        httpHeaders.append('content-type','application/json')

        return this.httpclient.post("http://localhost:8080/api/projects", createResource, {headers: httpHeaders});
    }
}