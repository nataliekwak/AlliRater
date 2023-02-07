import { Component, OnInit } from '@angular/core';
import {FormControl, Validators} from '@angular/forms';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})

export class RegisterComponent implements OnInit {
  Roles: any = ['Student', 'Faculty', 'Guest'];
  email = new FormControl('', [Validators.required, Validators.email]);
  user = new FormControl('', [Validators.required]);
  pass = new FormControl('', [Validators.required]);

  constructor() { }
  ngOnInit() {
  }

  getErrorMessage()
  {
    if (this.email.hasError('required') || this.user.hasError('required') || this.pass.hasError('required'))
    {
      return 'You must enter a value'
    }

    return this.email.hasError('email') ? 'Not a valid email' : '';
  }
}