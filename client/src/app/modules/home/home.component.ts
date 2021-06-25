import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { PrimeResponse } from 'app/models/http/prime-response';
import { PrimeService } from 'app/services/prime.service';
import { Observable, of, Subject, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  primeResponse$: Observable<PrimeResponse> | undefined
  loadingError$ = new Subject<boolean>()
  number: number | undefined

  constructor(private _primeService: PrimeService) { }

  ngOnInit(): void {
    
  }

  submitValue(number: number): void{
    this.number = number
    this.primeResponse$ = this._primeService.getLargestPrime(number).pipe(
      
      catchError((error: HttpErrorResponse) => {
          this.loadingError$.next(true)
          return throwError(error)
      })
    )
  }

}
