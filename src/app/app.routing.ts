import { ProjectsComponent } from "./projects/projects.component";
import { CreateComponent } from "./create/create.component";

export const AvailableRoutes: any = [
    { path: "", component: ProjectsComponent},
    { path: "create", component: CreateComponent }
];