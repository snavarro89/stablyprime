import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { retry, catchError } from 'rxjs/operators';

import { PrimeResponse } from 'app/models/http/prime-response';
import { environment } from 'environments/environment';

@Injectable({
  providedIn: 'root'
})
export class PrimeService {

  constructor(
    private _http: HttpClient
  ) { }

  getLargestPrime(number: number): Observable<PrimeResponse>{
    return this._http.get<PrimeResponse>(`${environment.api_url}/prime/${number}`)
    .pipe(
      retry(1),
      catchError(this.handleError)
    )
  }

  private handleError(error: any) {
    let errorMessage = '';
    if(error.error instanceof ErrorEvent) {
      // Get client-side error
      errorMessage = error.error.message;
    } else {
      // Get server-side error
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.error.desc}`;
    }
    console.log(errorMessage)
    return throwError(errorMessage);
 }
}
