export class PrimeResponse {
    public constructor(init? : Partial<PrimeResponse>){
        Object.assign(this,init)
    }

    primeNumber: number = 0
}