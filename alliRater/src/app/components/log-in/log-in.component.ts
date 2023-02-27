import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import {FormControl, Validators} from '@angular/forms';
import { AuthService } from '../../_services/auth.service';
import { StorageService } from '../../_services/storage.service';

@Component({
  selector: 'app-log-in',
  templateUrl: './log-in.component.html',
  styleUrls: ['./log-in.component.css']
})

export class LogInComponent {

  form: any = {
    username: null,
    password: null
  };
  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';
  roles: string[] = [];

  //user = new FormControl('', [Validators.required]);
  //pass = new FormControl('', [Validators.required]);

  constructor(private authService: AuthService, private storageService: StorageService, private router: Router){}

  ngOnInit(): void {
    if (this.storageService.isLoggedIn()) 
    {
      this.isLoggedIn = true;
      this.roles = this.storageService.getUser().roles;
    }
  }

  onSubmit(){
    const { username, password } = this.form;

    this.router.navigate(['homepage'])

    this.authService.login(username, password).subscribe({
      next: data => {
        this.storageService.saveUser(data);

        this.isLoginFailed = false;
        this.isLoggedIn = true;
        this.roles = this.storageService.getUser().roles;
        this.reloadPage();
      },
      error: err => {
        this.errorMessage = err.error.message;
        this.isLoginFailed = true;
      }
    });
  }

  reloadPage(): void {
    window.location.reload();
  }

  // getErrorMessage()
  // {
  //   if (this.user.hasError('required') || this.pass.hasError('required'))
  //   {
  //     return 'You must enter a value'
  //   }

  //   return;
  // }
}
