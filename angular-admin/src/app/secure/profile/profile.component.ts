import { Component, OnInit } from '@angular/core';
import {FormGroup, FormBuilder} from '@angular/forms';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {
  infoForm: FormGroup;
  passwordForm: FormGroup;

  constructor(
    private formsBuilder: FormBuilder,
    private authService: AuthService
    ){}
  
  ngOnInit(): void{
    this.infoForm = this.formsBuilder.group({
      first_name: '',
      last_name: '',
      email: '',
    });

    this.passwordForm = this.formsBuilder.group({
      password: '',
      password_confirm:''
    });

  }

  infoSubmit(): void{
    this.authService.updateInfo(this.infoForm.getRawValue()).subscribe(
      res => console.log(res)
    );

    }

    passwordSubmit(): void{
      this.authService.updatePassword(this.passwordForm.getRawValue()).subscribe(
        res => console.log(res)
      );
    }

}
