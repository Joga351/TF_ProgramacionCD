import { Component } from '@angular/core';
import { AfterViewInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Info } from './models/info';
import { HttpDataService } from './services/http-data.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'CrimesGO';
  info: Info[];
  index = 0;
  cai = 'no';
  edad = ' ';
  trabajo = ' ';
  vinculo = ' ';
  tipovio=' ';
  cAlcohol=' ';
  fuma = '';
  cDroga=' ';
  adiccion= '';
  riesgo= '';
  mes='';

  constructor(private httpDataService: HttpDataService ) {
  }

  async ngOnInit(){
    this.info = []
  }



  



}




