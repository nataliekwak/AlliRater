import { Component, OnInit } from '@angular/core';
import {FormGroup, FormBuilder} from '@angular/forms';
import { AuthService } from 'src/app/services/auth.service';
import { Auth } from 'src/app/classes/auth';

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

    Auth.userEmitter.subscribe(
      user => {
        this.infoForm.patchValue(user);
      });

    this.passwordForm = this.formsBuilder.group({
      password: '',
      password_confirm:''
    });

  }

  infoSubmit(): void{
    this.authService.updateInfo(this.infoForm.getRawValue()).subscribe(
      user => Auth.userEmitter.emit(user)
    );

    }

    passwordSubmit(): void{
      this.authService.updatePassword(this.passwordForm.getRawValue()).subscribe(
        res => console.log(res)
      );
    }

}
