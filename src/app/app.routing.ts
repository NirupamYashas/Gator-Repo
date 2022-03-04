import { ProjectsComponent } from "./projects/projects.component";
import { CreateComponent } from "./create/create.component";
import { HomepageComponent } from "./homepage/homepage.component";
import { RegisterComponent } from "./register/register.component";
import { LoginComponent } from "./login/login.component";

export const AvailableRoutes: any = [
    { path: "", component: ProjectsComponent},
    { path: "create", component: CreateComponent },
    { path: "homepage", component: HomepageComponent},
    { path: "register", component: RegisterComponent},
    { path: "login", component: LoginComponent}
];