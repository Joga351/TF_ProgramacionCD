import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
 
import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { ListarComponent } from './Persona/listar/listar.component';
import { ServiceService } from './Persona/Service/service.service';


@NgModule({
  declarations: [
    AppComponent,
    ListarComponent,

  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
