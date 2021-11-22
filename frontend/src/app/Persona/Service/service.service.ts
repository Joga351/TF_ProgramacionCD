import { Injectable } from '@angular/core';
import { HttpClient} from '@angular/common/http';
import { Persona } from '../../models/Persona';

@Injectable({
  providedIn: 'root'
})
export class ServiceService {

  constructor(private http:HttpClient) { }

  url='http://localhost:3505/listarData';

  getPersonas(){
    return this.http.get<Persona[]>(this.url)
  }
}
