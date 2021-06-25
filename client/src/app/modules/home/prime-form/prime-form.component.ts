import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { AbstractControl, FormBuilder, FormGroup, ValidationErrors, ValidatorFn } from '@angular/forms';
import { Validators } from '@angular/forms';

@Component({
  selector: 'prime-form',
  templateUrl: './prime-form.component.html',
  styleUrls: ['./prime-form.component.scss']
})
export class PrimeFormComponent implements OnInit {

  @Output() submitPrime: EventEmitter<number> = new EventEmitter<number>();


  primeForm =  new FormGroup({})

  constructor(
    private _formBuilder: FormBuilder
  ) { }

  ngOnInit(): void {
    this.primeForm = this._formBuilder.group({
      number: ['', [Validators.required, positiveIntegerValidator()]]
    })
  }

  get number() { return this.primeForm.get('number'); }

  submit(): void{
    //emit the event back to the parent component.
  }

}

export function positiveIntegerValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    return control.value < 0 ? {invalidNumber: true}: null;
  };
}
