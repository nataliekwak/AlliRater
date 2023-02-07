import { Component } from '@angular/core';
import { Router } from '@angular/router';
import {FormControl, Validators} from '@angular/forms';

@Component({
  selector: 'app-log-in',
  templateUrl: './log-in.component.html',
  styleUrls: ['./log-in.component.css']
})
export class LogInComponent {

  user = new FormControl('', [Validators.required]);
  pass = new FormControl('', [Validators.required]);

  constructor(private router: Router){
    
  }

  onSubmit(){
    this.router.navigate(['homepage'])
  }

  getErrorMessage()
  {
    if (this.user.hasError('required') || this.pass.hasError('required'))
    {
      return 'You must enter a value'
    }

    return;
    // ADD : Check is user exists in database
  }
}
