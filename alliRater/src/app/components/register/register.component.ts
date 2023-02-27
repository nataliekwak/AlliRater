import { Component, OnInit } from '@angular/core';
//import {FormControl, Validators} from '@angular/forms';
import { AuthService } from '../../_services/auth.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})

export class RegisterComponent implements OnInit {
  form: any = {
    username: null,
    email: null,
    password: null
  };
  isSuccessful = false;
  isSignUpFailed = false;
  errorMessage = '';

  Roles: any = ['Student', 'Faculty', 'Guest'];
  //email = new FormControl('', [Validators.required, Validators.email]);
  //user = new FormControl('', [Validators.required]);
  //pass = new FormControl('', [Validators.required]);

  constructor(private authService: AuthService) { }
  ngOnInit(): void {
  }

  onSubmit(): void {
    const { username, email, password } = this.form;

    this.authService.register(username, email, password).subscribe({
      next: (data: any) => {
        console.log(data);
        this.isSuccessful = true;
        this.isSignUpFailed = false;
      },
      error: (err: { error: { message: string; }; }) => {
        this.errorMessage = err.error.message;
        this.isSignUpFailed = true;
      }
    });
  }

  // getErrorMessage()
  // {
  //   if (this.email.hasError('required') || this.user.hasError('required') || this.pass.hasError('required'))
  //   {
  //     return 'You must enter a value'
  //   }

  //   return this.email.hasError('email') ? 'Not a valid email' : '';
  // }
}