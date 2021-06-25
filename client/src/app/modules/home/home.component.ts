import { Component, OnInit } from '@angular/core';
import { PrimeResponse } from 'app/models/http/prime-response';
import { PrimeService } from 'app/services/prime.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  constructor(private _primeService: PrimeService) { }

  ngOnInit(): void {
    this._primeService.getLargestPrime(-2).subscribe( (primeResponse: PrimeResponse) => {
      console.log(primeResponse)
    })
  }

}
