import { Injectable } from "@angular/core";
import { HttpClient, HttpErrorResponse, HttpHeaders } from "@angular/common/http";
import { Observable, throwError } from "rxjs";
import {catchError, retry} from 'rxjs/operators';
import { Info } from "../models/info";

@Injectable({
    providedIn: 'root'
})
export class HttpDataService {
    basePath = 'http://localhost:3505';
    constructor(private http:HttpClient){

    }

    httpOptions = {
        headers: new HttpHeaders({

        })
    };

    handleError(error: HttpErrorResponse): Observable<never> {
        if(error.error instanceof ErrorEvent){
            console.log('An error ocurred: ', error.error.message);
        } else {
            console.log(`API returned code ${error.status}, body was: ${error.error}`);
        }
        return throwError('Something happened with request, please try again later');

    }
    // Get All
    getAll(): Observable<Info> {
        return this.http.get<Info>(`${this.basePath}/listarData`, this.httpOptions)
            .pipe(retry(2), catchError(this.handleError));
    }

    // Get by id
    getById(index: any): Observable<Info> {
        return this.http.get<Info>(`${this.basePath}/buscarDato?Index=${index}`, this.httpOptions)
            .pipe(retry(2), catchError(this.handleError));
    }

}