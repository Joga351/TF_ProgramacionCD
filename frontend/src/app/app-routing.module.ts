import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { ListarComponent } from "./Persona/listar/listar.component";


const routes: Routes = [
    {path:'listar', component:ListarComponent}
];


@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule{}